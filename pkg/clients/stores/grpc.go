package stores

import (
	"context"
	"fmt"
	"time"

	"github.com/Tlantic/go-util/v4/errors"
	"github.com/Tlantic/go-util/v4/grpcutil"
	pool "github.com/Tlantic/go-util/v4/grpcutil/connection-pool"
	"github.com/Tlantic/go-util/v4/mrs"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/retryer"
	pb "github.com/Tlantic/mrs-service-cab-connector/proto/store_pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var _ transport = (*grpcTransport)(nil)

type grpcTransport struct {
	addr              string
	pool              *pool.Pool
	retryer           retryer.Retryer
	connectionTimeout time.Duration
	dialOpts          []grpc.DialOption
	callOpts          []grpc.CallOption
}

func (g *grpcTransport) ListAllStores(ctx context.Context, opts ...grpc.CallOption) (StoreList, error) {
	// Retrieve connection
	ctxConn, cancel := context.WithDeadline(ctx, time.Now().Add(g.connectionTimeout))
	defer cancel()

	conn, err := g.pool.Get(ctxConn)
	if err != nil {
		return nil, fmt.Errorf("error retrieving grpc connection: %w", err)
	}

	// Defer closing the connection
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error returning grpc connection: %s", err)
		}
	}()

	// Prepare variables
	var (
		newCtx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

		cli      = pb.NewStoresClient(conn.Unwrap())
		callOpts = append(g.callOpts, opts...)

		limit, resultCount uint64 = 30, 30
		i                  uint64

		allStores  StoreList
		req        = &pb.ListStoresRequest{Fields: nil, Limit: limit}
		res        *pb.ListStoresResponse
		errAttempt error
	)

	for i = 0; resultCount == limit; i++ {
		req.Offset = i * limit

		// Make request
		if err := g.retryer.Do(newCtx, func() (bool, error) {
			res, errAttempt = cli.ListStores(newCtx, req, callOpts...)
			if errAttempt != nil {
				errAttempt = errorFromGRPC(errAttempt)
				if !errors.IsTemporary(errAttempt) {
					return true, errAttempt
				}
				return false, nil
			}
			return true, nil
		}); err != nil {
			if err == retryer.ErrMaxStepsReached && errAttempt != nil {
				err = errors.Wrap(errAttempt, err.Error())
			}
			return nil, err
		}

		storeValues := res.GetResult()
		if storeValues == nil || len(storeValues) == 0 {
			break
		}

		for _, store := range storeValues {
			allStores = append(allStores, Store{
				ID:           store.Id,
				Name:         store.Name,
				ExternalCode: store.ExternalCode,
			})
		}

		resultCount = uint64(len(storeValues))
	}

	return allStores, nil
}

// GetStore retrieves a store by its id.
// Example:
//
//	client.GetStore(ctx, "Store123")
func (g *grpcTransport) GetStore(ctx context.Context, storeId string, opts ...grpc.CallOption) (*Store, error) {
	// Retrieve connection
	ctxConn, cancel := context.WithDeadline(ctx, time.Now().Add(g.connectionTimeout))
	defer cancel()

	conn, err := g.pool.Get(ctxConn)
	if err != nil {
		return nil, fmt.Errorf("error retrieving grpc connection: %w", err)
	}

	// Defer closing the connection
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error returning grpc connection: %s", err)
		}
	}()

	// Prepare variables
	var (
		newCtx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

		cli      = pb.NewStoresClient(conn.Unwrap())
		callOpts = append(g.callOpts, opts...)

		store      Store
		req        = &pb.GetStoreRequest{StoreId: storeId}
		res        *pb.GetStoreResponse
		errAttempt error
	)

	// Make request
	if err := g.retryer.Do(ctx, func() (bool, error) {
		res, errAttempt = cli.GetStore(newCtx, req, callOpts...)
		if errAttempt != nil {
			errAttempt = errorFromGRPC(errAttempt)
			if !errors.IsTemporary(errAttempt) {
				return true, errAttempt
			}
			return false, nil
		}
		return true, nil
	}); err != nil {
		if err == retryer.ErrMaxStepsReached && errAttempt != nil {
			err = errors.Wrap(errAttempt, err.Error())
		}
		return nil, err
	}

	store = NewStoreFromProto(res.GetResult())
	return &store, nil
}
