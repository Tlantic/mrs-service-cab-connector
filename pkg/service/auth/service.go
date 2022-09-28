package auth

import (
	"context"

	apierr "github.com/Tlantic/mrs-service-cab-connector/pkg/errors"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	erylogrus "github.com/Tlantic/mrs-service-cab-connector/utils/logger/logrus"
	"github.com/sirupsen/logrus"
)

type MrsAuthCabConnectorSRV struct {
	logger   *erylogrus.Logger
	username string
	password string
	endpoint string
}

type MrsAuthCabConnectorOptions struct {
	Username string
	Password string
	Endpoint string
}

func NewMrsAuthCabConnectorSRV(options *MrsAuthCabConnectorOptions) *MrsAuthCabConnectorSRV {
	lg := logrus.WithField("component", "NewMrsAuthCabConnectorSRV")
	lg.Info("initializing")

	return &MrsAuthCabConnectorSRV{
		logger:   erylogrus.NewLogger(lg),
		username: options.Username,
		password: options.Password,
		endpoint: options.Endpoint,
	}
}

func (srv *MrsAuthCabConnectorSRV) GetTokenWithoutCredentials(_ context.Context, _ *proto.GetTokenWithoutCredentialsRequest) (*proto.AuthResponse, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsAuthCabConnectorSRV) GetTokenWithCredentials(_ context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsAuthCabConnectorSRV) GetLoginByUsername(_ context.Context, _ *proto.AuthRequest) (*proto.AuthResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsAuthCabConnectorSRV) GetUserStores(_ context.Context, _ *proto.GetUserStoresRequest) (*proto.GetUserStoresResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsAuthCabConnectorSRV) GetUserInfo(_ context.Context, _ *proto.GetUserInfoRequest) (*proto.GetUserInfoResult, error) {
	return nil, apierr.ErrNotImplemented()
}
