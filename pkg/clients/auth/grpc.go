package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Tlantic/go-util/errors"
	"github.com/Tlantic/go-util/grpcutil"
	pool "github.com/Tlantic/go-util/grpcutil/connection-pool"
	"github.com/Tlantic/go-util/mrs"
	"github.com/Tlantic/go-util/mrs/log"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/retryer"
	pb "github.com/Tlantic/mrs-service-cab-connector/proto"
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

func (g *grpcTransport) Auth(ctx context.Context, credentials *Credentials, opts ...grpc.CallOption) (string, error) {
	// Retrieve connection
	ctxConn, cancel := context.WithDeadline(ctx, time.Now().Add(g.connectionTimeout))
	defer cancel()

	conn, err := g.pool.Get(ctxConn)
	if err != nil {
		return "", fmt.Errorf("error retrieving grpc connection: %w", err)
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

		cli      = pb.NewMRSConnectorAPIClient(conn.Unwrap())
		callOpts = append(g.callOpts, opts...)

		req        pb.AuthRequest
		res        *pb.AuthResponse
		errAttempt error
	)

	if credentials != nil {
		req.Username = credentials.Username
		req.Password = credentials.Password
	}

	// Make request
	if err := g.retryer.Do(ctx, func() (bool, error) {
		res, errAttempt = cli.Auth(newCtx, &req, callOpts...)
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
		return "", err
	}

	return res.GetAccessToken(), nil
}
