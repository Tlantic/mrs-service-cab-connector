package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/Tlantic/go-util/v4/grpcutil"
	"github.com/Tlantic/go-util/v4/mrs"
	"github.com/Tlantic/mrs-service-cab-connector/internal/configuration/catalog"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/auth"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/stores"
	apierr "github.com/Tlantic/mrs-service-cab-connector/pkg/errors"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/formatter"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	erylogrus "github.com/Tlantic/mrs-service-cab-connector/utils/logger/logrus"
	"github.com/andrepinto/erygo"

	log "github.com/sirupsen/logrus"
)

const (
	innerRequest = "cabInnerRequest"
	emptyString  = ""

	// Price types
	erp = "erp"
	pos = "pos"
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
	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))
	var stockResps []*proto.StockResponse
	switch {
	case req.Store != emptyString:

		// Convert Store Attrib to EIN
		storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
		if err != nil {
			return nil, err
		}

		// Call PriceStock service
		respBytes, err := srv.CallExternalService(req.Sku, catalog.StockPriceResourceKey, storeInfo.ExternalCode)
		if err != nil {
			return nil, err
		}

		// Deserialize response
		var cliResp *ClientStockPriceResponse
		err = json.Unmarshal(respBytes, &cliResp)
		if err != nil {
			return nil, err
		}

		if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
			return nil, apierr.ErrSkuNotFound()
		}

		stockResp := &proto.StockResponse{
			Sku:         req.Sku,
			StoreId:     req.Store,
			StockOnHand: cliResp.ServiceResponse.ProductDetails.Estoque1,
		}

		stockResps = append(stockResps, stockResp)

	case req.Group != emptyString:
		return nil, apierr.ErrNotImplemented()
	case req.Chain != emptyString:
		return nil, apierr.ErrNotImplemented()

	default:
		storesStock, err := srv.GetStoresStock(ctx, &proto.GetStoresStockRequest{Sku: req.Sku})
		if err != nil {
			return nil, err
		}

		stockResps = storesStock.Result
	}

	result := proto.GetStockResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  stockResps,
	}

	return &result, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetStocks(ctx context.Context, req *proto.GetStocksRequest) (*proto.GetStocksResult, error) {
	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to EIN
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Skus[0])
	if err != nil {
		return nil, err
	}

	// Call PriceStock service
	respBytes, err := srv.CallExternalService(req.Skus[0], catalog.StockPriceResourceKey, storeInfo.ExternalCode)
	if err != nil {
		return nil, err
	}

	var stockResps []*proto.GetStocksResponse

	// Deserialize response
	var cliResp *ClientStockPriceResponse
	err = json.Unmarshal(respBytes, &cliResp)
	if err != nil {
		return nil, err
	}

	if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
		return nil, apierr.ErrSkuNotFound()
	}

	stockResp := &proto.GetStocksResponse{
		Sku:   req.Skus[0],
		Stock: cliResp.ServiceResponse.ProductDetails.Estoque1,
	}

	stockResps = append(stockResps, stockResp)

	result := proto.GetStocksResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  stockResps,
	}

	return &result, nil

}

func (srv *MrsCatalogCabConnectorSRV) GetPrice(ctx context.Context, req *proto.GetPriceRequest) (*proto.GetPriceResult, error) {

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to EIN
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.StockPriceResourceKey, storeInfo.ExternalCode)
	if err != nil {
		return nil, err
	}

	// Deserialize response
	var cliResp *ClientStockPriceResponse
	err = json.Unmarshal(respBytes, &cliResp)
	if err != nil {
		return nil, err
	}

	if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
		return nil, apierr.ErrSkuNotFound()
	}

	result := proto.GetPriceResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result: &proto.PriceResponse{
			Type:  erp,
			Sku:   req.Sku,
			Store: req.Store,
			Value: cliResp.ServiceResponse.ProductDetails.Preco,
		},
	}

	return &result, nil
}

func (srv *MrsCatalogCabConnectorSRV) CallExternalService(product string, resourceId string, externalCode string) ([]byte, error) {
	extCfg := catalog.Get(nil).External
	uri := extCfg.BaseURI + "/" + product + "/" + externalCode

	extReq, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, err := http.DefaultClient.Do(extReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read response's body
	resContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http request returned status: %s / message: %s", resp.Status, string(resContent)))
	}

	return resContent, nil

}

func (srv *MrsCatalogCabConnectorSRV) GetPricePOS(ctx context.Context, req *proto.GetPricePOSRequest) (*proto.GetPricePOSResult, error) {

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to EIN
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.StockPriceResourceKey, storeInfo.ExternalCode)
	if err != nil {
		return nil, err
	}

	// Deserialize response
	var cliResp *ClientStockPriceResponse
	err = json.Unmarshal(respBytes, &cliResp)
	if err != nil {
		return nil, err
	}

	if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
		return nil, apierr.ErrSkuNotFound()
	}

	result := proto.GetPricePOSResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result: &proto.PriceResponse{
			Type:  pos,
			Sku:   req.Sku,
			Store: req.Store,
			Value: cliResp.ServiceResponse.ProductDetails.Preco,
		},
	}

	return &result, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetPrices(ctx context.Context, req *proto.GetPricesRequest) (*proto.GetPricesResult, error) {

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to EIN
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.StockPriceResourceKey, storeInfo.ExternalCode)
	if err != nil {
		return nil, err
	}

	// Deserialize response
	var cliResp *ClientStockPriceResponse
	err = json.Unmarshal(respBytes, &cliResp)
	if err != nil {
		return nil, err
	}

	if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
		return nil, apierr.ErrSkuNotFound()
	}

	// Initialize Proto Prices Response
	var PricesResp []*proto.PricesResponse

	ERPPrice := cliResp.ServiceResponse.ProductDetails.Preco

	POSPrice := cliResp.ServiceResponse.ProductDetails.PrecoPDV

	PricesResp = []*proto.PricesResponse{
		{
			Type:  "erp",
			Group: "erp",
			Value: ERPPrice,
		},
		{
			Type:  "pos",
			Group: "pos",
			Value: POSPrice,
		},
	}

	result := proto.GetPricesResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  PricesResp,
	}

	return &result, nil

}

func (srv *MrsCatalogCabConnectorSRV) GetLabel(ctx context.Context, req *proto.GetLabelRequest) (*proto.GetLabelResult, error) {

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to EIN
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.StockPriceResourceKey, storeInfo.ExternalCode)
	if err != nil {
		return nil, err
	}

	// Deserialize response
	var cliResp *ClientLabelResponseBytes
	err = json.Unmarshal(respBytes, &cliResp)
	if err != nil {
		return nil, err
	}

	if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
		return nil, apierr.ErrSkuNotFound()
	}

	// Map bytes to map <string><string>
	fields, err := formatter.Mapify(cliResp.ServiceResponse.ProductDetails)
	if err != nil {
		return nil, err
	}

	// Create Proto Response
	result := proto.GetLabelResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result: &proto.LabelResponse{
			Code: req.Code,
			Data: fields,
		},
	}
	return &result, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetStoresStock(ctx context.Context, req *proto.GetStoresStockRequest) (*proto.GetStoresStockResult, error) {
	log.Info("GRPC - GetStoresStock")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Printf("mrsCtx: %+v", mrsCtx)

	// Fetch all store ids
	allStores, err := srv.storesClient.ListAllStores(ctx)
	if err != nil {
		return nil, err
	}

	// For every store, get stock
	storesStock := struct {
		data []*StoresStockInfo
		mx   *sync.Mutex
	}{[]*StoresStockInfo{}, &sync.Mutex{}}

	var wg sync.WaitGroup
	wg.Add(len(allStores))

	ctx = context.WithValue(ctx, innerRequest, true)

	var erygoErrs []*erygo.Err
	for i := range allStores {
		go func(store *stores.Store) {
			defer wg.Done()

			// Get Stock
			if store.ID != emptyString {
				stockResp, err := srv.GetStock(ctx, &proto.GetStockRequest{
					Store: store.ID,
					Sku:   req.Sku,
				})

				if err != nil {
					switch err := err.(type) {
					case *erygo.Err:
						erygoErrs = append(erygoErrs, err)
					}

					log.Warn("GetStock error: " + err.Error())
					return
				}

				storeStock := StoresStockInfo{
					StoreID:         store.ID,
					StoreName:       store.Name,
					StockAfterSales: stockResp.Result[0].StockOnHand,
				}

				storesStock.mx.Lock()
				storesStock.data = append(storesStock.data, &storeStock)
				storesStock.mx.Unlock()
			}

		}(&allStores[i])
	}

	wg.Wait()

	for _, erygoErr := range erygoErrs {
		if erygoErr.StatusHTTP == http.StatusUnauthorized {
			return nil, erygoErr
		}
	}

	if len(storesStock.data) < 1 {
		return nil, apierr.ErrValidation().AddDetails("couldn't retrieve stock for given sku")
	}

	result := proto.GetStoresStockResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  stockOtherStoresToProto(storesStock.data),
	}

	return &result, nil
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
