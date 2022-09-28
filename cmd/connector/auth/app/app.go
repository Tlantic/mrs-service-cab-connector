package app

import (
	"time"

	"github.com/Tlantic/mrs-service-cab-connector/cmd/connector/auth/app/server"
	authConfig "github.com/Tlantic/mrs-service-cab-connector/internal/configuration/auth"
	service "github.com/Tlantic/mrs-service-cab-connector/pkg/service/auth"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/shutdown"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	log "github.com/sirupsen/logrus"
)

// AuthApp ...
type AuthApp struct {
}

// NewAuthApp ...
func NewAuthApp() *AuthApp {
	return &AuthApp{}
}

// Run ...
func (cli *AuthApp) Run(options *AuthCmdOptions) error {

	// Configuration
	configs := authConfig.Get(&authConfig.Options{
		ConfigName: options.Service,
	})

	// Initialize Authentication Service
	var srv proto.MRSAuthConnectorExternalServer = service.NewMrsAuthCabConnectorSRV(&service.MrsAuthCabConnectorOptions{
		Username: configs.Username,
		Password: configs.Password,
		Endpoint: configs.External.BaseURI + configs.External.Resources[authConfig.ResourceKey].Path,
	})

	servers := []server.Server{
		server.NewGRPCServer(configs.Ports.GRPC, srv),
	}

	server.RunServers(servers...)

	sd := shutdown.NewShutdown(time.Duration(1), log.WithField("process", "shutdown"))

	go func() {
		select {
		case <-sd.InitiateChannel:
			server.StopServers(servers...)
			sd.DoneChannel <- true
		}
	}()

	sd.WaitForSignal()

	return nil
}
