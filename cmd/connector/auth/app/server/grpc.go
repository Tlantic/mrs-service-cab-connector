package server

import (
	"context"
	"errors"
	"net"
	"runtime/debug"
	"time"

	"github.com/Tlantic/go-util/v4/mrs"
	"github.com/Tlantic/go-util/v4/mrs/collection"
	apierr "github.com/Tlantic/mrs-service-cab-connector/pkg/errors"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	erygogrpc "github.com/Tlantic/mrs-service-cab-connector/utils/middleware/erygo/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

type grpcServer struct {
	listenAddr string
	server     *grpc.Server
}

func panicHandler(p interface{}) (err error) {
	logrus.Errorf("panic: %v", p)
	debug.PrintStack()
	return apierr.ErrInternal()
}

//NewGRPCServer ...

func NewGRPCServer(listenAddr string, service proto.MRSAuthConnectorExternalServer) Server {
	erygogrpc.JSONMarshal = jsoniter.ConfigFastest.Marshal
	erygogrpc.JSONUnmarshal = jsoniter.ConfigFastest.Unmarshal
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			erygogrpc.UnaryServerInterceptor(apierr.ErrInternal),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
			grpc_logrus.UnaryServerInterceptor(logrus.WithField("component", "grpc_server")),
			includeContextMetadata,
		),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	proto.RegisterMRSAuthConnectorExternalServer(server, service)

	return &grpcServer{
		listenAddr: listenAddr,
		server:     server,
	}
}

func (s *grpcServer) Run() error {
	logrus.WithField("listenAddr", s.listenAddr).Infof("Starting GRPC Server")
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	return s.server.Serve(listener)
}

func (s *grpcServer) Stop() error {
	logrus.Infof("Stopping GRPC server")
	s.server.GracefulStop()
	return nil
}

func includeContextMetadata(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get context values")
	}

	// Update context
	mrsCtx := mrs.NewMRSContext(mrs.FromMetadata(collection.Metadata(md)))
	newCtx := mrs.SetMRSContext(ctx, mrsCtx)

	return handler(newCtx, req)
}
