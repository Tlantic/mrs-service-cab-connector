package catalog

import (
	"context"

	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/auth"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/stores"
	apierr "github.com/Tlantic/mrs-service-cab-connector/pkg/errors"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	erylogrus "github.com/Tlantic/mrs-service-cab-connector/utils/logger/logrus"
	log "github.com/sirupsen/logrus"
)

type MrsCatalogCabConnectorSRV struct {
	logger       *erylogrus.Logger
	storesClient stores.Client
	authClient   auth.Client
}

type MrsCatalogCabConnectorOptions struct {
	StoresClient stores.Client
	AuthClient   auth.Client
}

//MrsCatalogCabConnectorSRV

func NewMrsCatalogCabConnectorSRV(opts *MrsCatalogCabConnectorOptions) *MrsCatalogCabConnectorSRV {
	lg := log.WithField("component", "NewMrsCatalogCabConnectorSRV")
	lg.Info("initializing")

	return &MrsCatalogCabConnectorSRV{
		logger:       erylogrus.NewLogger(lg),
		storesClient: opts.StoresClient,
		authClient:   opts.AuthClient,
	}
}

// Connector Methods

func (srv *MrsCatalogCabConnectorSRV) GetStock(ctx context.Context, req *proto.GetStockRequest) (*proto.GetStockResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetStocks(ctx context.Context, req *proto.GetStocksRequest) (*proto.GetStocksResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetPrice(ctx context.Context, req *proto.GetPriceRequest) (*proto.GetPriceResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetPricePOS(ctx context.Context, req *proto.GetPricePOSRequest) (*proto.GetPricePOSResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetPrices(ctx context.Context, req *proto.GetPricesRequest) (*proto.GetPricesResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetLabel(ctx context.Context, req *proto.GetLabelRequest) (*proto.GetLabelResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetStoresStock(ctx context.Context, req *proto.GetStoresStockRequest) (*proto.GetStoresStockResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetSkuStoreSales(_ context.Context, _ *proto.GetSkuStoreSalesRequest) (*proto.GetSkuStoreSalesResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetDamagesDestinations(_ context.Context, _ *proto.GetDamagesDestinationsRequest) (*proto.GetDamagesDestinationsResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetDamagesReasons(_ context.Context, _ *proto.GetDamageReasonsRequest) (*proto.GetDamageReasonsResult, error) {
	return nil, apierr.ErrNotImplemented()
}
