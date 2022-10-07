package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Tlantic/go-util/v4/cache"
	pool "github.com/Tlantic/go-util/v4/grpcutil/connection-pool"
	"github.com/Tlantic/mrs-service-cab-connector/internal/configuration/catalog"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/retryer"
	"google.golang.org/grpc"
)

type Client interface {
	// Auth retrieves a valid access token for the given context
	Auth(ctx context.Context, credentials *Credentials, opts ...grpc.CallOption) (string, error)
}

type transport interface {
	Client
}

const (
	defaultInitialConnections = 1
	defaultMaxConnections     = 5
	defaultDialTimeOut        = 10 * time.Second
	defaultIdleTimeout        = 0
	defaultConnTimeOut        = 30 * time.Second
)

var (
	_                         Client = (*client)(nil)
	defaultRetryer                   = retryer.NewExponentialBackoff(50*time.Millisecond, 1.75, 1.5, 5, 0)
	defaultConnectionPoolOpts        = &catalog.ConnectionPoolOpts{
		InitialConnections: defaultInitialConnections,
		MaxConnections:     defaultMaxConnections,
		DialTimeout:        defaultDialTimeOut,
		IdleTimeout:        defaultIdleTimeout,
		ConnectionTimeout:  defaultConnTimeOut,
	}
)

type ClientOption func(client *client)

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
		var retryerValue retryer.Retryer
		if retryOpts == nil {
			return

		}

		retryerValue = retryer.NewExponentialBackoff(
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
		retryer:  defaultRetryer,
		poolOpts: defaultConnectionPoolOpts,
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

	return cli, nil
}
