package stores

import (
	"context"
	"fmt"
	"time"

	"github.com/Tlantic/go-util/cache"
	pool "github.com/Tlantic/go-util/grpcutil/connection-pool"
	"github.com/Tlantic/mrs-service-cab-connector/internal/configuration/catalog"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/retryer"
	"google.golang.org/grpc"
)

const (
	defaultRetryDuration = 50 * time.Millisecond
	defaultRetryFactor   = 1.75
	defaultRetryJitter   = 1.5
	defaultRetrySteps    = 5
	defaultRetryCap      = 0

	defaultInitialConnections = 1
	defaultMaxConnections     = 5
	defaultDialTimeOut        = 10 * time.Second
	defaultIdleTimeout        = 0
	defaultConnTimeOut        = 30 * time.Second
)

var (
	_              Client = (*client)(nil)
	defaultRetryer        = retryer.NewExponentialBackoff(
		defaultRetryDuration,
		defaultRetryFactor,
		defaultRetryJitter,
		defaultRetrySteps,
		defaultRetryCap,
	)

	defaultConnectionPoolOpts = &catalog.ConnectionPoolOpts{
		InitialConnections: defaultInitialConnections,
		MaxConnections:     defaultMaxConnections,
		DialTimeout:        defaultDialTimeOut,
		IdleTimeout:        defaultIdleTimeout,
		ConnectionTimeout:  defaultConnTimeOut,
	}
)

type Client interface {
	// GetStore retrieves a store by its id.
	// Example:
	// 		client.GetStore(ctx, "Store123")
	GetStore(ctx context.Context, storeId string, opts ...grpc.CallOption) (*Store, error)

	// ListAllStores retrieves a list of all stores for a given organization (fetched through the MRSContext)
	ListAllStores(ctx context.Context, opts ...grpc.CallOption) (StoreList, error)
}

type ClientOption func(client *client)

// UseCache activates cache functionality using the options given
func UseCache(cacheOpts *catalog.CacheOpts) ClientOption {
	return func(client *client) {
		if cacheOpts == nil {
			return
		}

		client.cache = cache.New(10, cacheOpts.PurgeCycle)
		client.ttl = cacheOpts.TTL
	}
}

// WithPoolOptions specifies the options to pass down to the connection pool creation
func WithPoolOptions(connectionPoolOpts *catalog.ConnectionPoolOpts) ClientOption {
	return func(client *client) {
		if connectionPoolOpts == nil {
			return
		}
		client.poolOpts = connectionPoolOpts
	}
}

// WithRetryOptions specifies the behaviour of faulty requests
func WithRetryOptions(retryOpts *catalog.RetryOpts) ClientOption {
	return func(client *client) {
		if retryOpts == nil {
			return
		}

		retryerValue := retryer.NewExponentialBackoff(
			retryOpts.Duration,
			retryOpts.Factor,
			retryOpts.Jitter,
			retryOpts.Steps,
			retryOpts.Cap,
		)

		client.retryer = retryerValue
	}
}

// WithDialOptions specifies the dial options to pass down when creating the grpc client connection
func WithDialOptions(options ...grpc.DialOption) ClientOption {
	return func(client *client) {
		client.grpcDialOpts = append(client.grpcDialOpts, options...)
	}
}

// WithCallOptions specifies the default call options to pass down to all grpc calls
func WithCallOptions(options ...grpc.CallOption) ClientOption {
	return func(client *client) {
		client.grpcCallOpts = append(client.grpcCallOpts, options...)
	}
}

type transport = Client

type client struct {
	transport
	poolOpts     *catalog.ConnectionPoolOpts
	cache        *cache.Map
	ttl          time.Duration
	retryer      retryer.Retryer
	grpcDialOpts []grpc.DialOption
	grpcCallOpts []grpc.CallOption
}

// NewClient returns a new client instance
func NewClient(address string, opts ...ClientOption) (Client, error) {
	cli := &client{
		poolOpts: defaultConnectionPoolOpts,
		retryer:  defaultRetryer,
	}
	for _, o := range opts {
		o(cli)
	}

	poolConn, err := pool.NewWithContext(context.Background(), func(parent context.Context) (*grpc.ClientConn, error) {
		ctx, cancel := context.WithTimeout(parent, cli.poolOpts.DialTimeout)
		defer cancel()

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return grpc.Dial(address, cli.grpcDialOpts...)
		}
	}, cli.poolOpts.InitialConnections, cli.poolOpts.MaxConnections, cli.poolOpts.IdleTimeout)
	if err != nil {
		return nil, fmt.Errorf("error setting up grpc connection pool: %w", err)
	}

	cli.transport = &grpcTransport{
		addr:              address,
		pool:              poolConn,
		retryer:           cli.retryer,
		connectionTimeout: cli.poolOpts.ConnectionTimeout,
		dialOpts:          cli.grpcDialOpts,
		callOpts:          cli.grpcCallOpts,
	}

	if cli.cache != nil {
		cli.transport = &cachedTransport{
			transport: cli.transport,
			ttl:       cli.ttl,
			cache:     cli.cache,
		}
	}

	return cli, nil
}
