package app

import (
	"github.com/Tlantic/mrs-service-cab-connector/cmd/connector/catalog/app/server"
	catalogconfig "github.com/Tlantic/mrs-service-cab-connector/internal/configuration/catalog"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/stores"
	service "github.com/Tlantic/mrs-service-cab-connector/pkg/service/catalog"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/shutdown"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

// CatalogApp ...
type CatalogApp struct {
}

// NewCatalogApp ...
func NewCatalogApp() *CatalogApp {
	return &CatalogApp{}
}

// Run ...
func (*CatalogApp) Run(options *CatalogCmdOptions) error {

	// Configuration
	configs := catalogconfig.Get(&catalogconfig.Options{
		ConfigName: options.Service,
	})

	// GRPC Clients
	// Retail/Stores
	storesCfg := configs.Internal[catalogconfig.RetailInternalKey]
	storesCliOpts := []stores.ClientOption{
		stores.UseCache(storesCfg.Cache),
		stores.WithPoolOptions(storesCfg.ConnectionPool),
		stores.WithRetryOptions(storesCfg.Retry),
		stores.WithDialOptions(grpc.WithInsecure()),
		stores.WithCallOptions(grpc.WaitForReady(true)),
	}

	storesCli, err := stores.NewClient(
		storesCfg.Endpoint.Host,
		storesCliOpts...,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Catalog Service
	var srv proto.MRSCatalogConnectorExternalServer
	srv = service.NewMrsCatalogCabConnectorSRV(&service.MrsCatalogCabConnectorOptions{
		StoresClient: storesCli,
	})

	// Launch grpc Server
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
