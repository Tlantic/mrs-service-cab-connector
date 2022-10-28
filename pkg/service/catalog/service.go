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
	"github.com/Tlantic/mrs-service-cab-connector/pkg/clients/stores"
	apierr "github.com/Tlantic/mrs-service-cab-connector/pkg/errors"
	"github.com/Tlantic/mrs-service-cab-connector/pkg/formatter"
	"github.com/Tlantic/mrs-service-cab-connector/proto"
	erylogrus "github.com/Tlantic/mrs-service-cab-connector/utils/logger/logrus"
	"github.com/andrepinto/erygo"

	log "github.com/sirupsen/logrus"
)

const (
	emptyString = ""

	// Price types
	erp = "erp"
	pos = "pos"
)

type MrsCatalogCabConnectorSRV struct {
	logger       *erylogrus.Logger
	storesClient stores.Client
}

type MrsCatalogCabConnectorOptions struct {
	StoresClient stores.Client
}

// NewMrsCatalogCabConnectorSRV ...
func NewMrsCatalogCabConnectorSRV(opts *MrsCatalogCabConnectorOptions) *MrsCatalogCabConnectorSRV {

	lg := log.WithField("component", "NewMrsCatalogCabConnectorSRV")
	lg.Info("initializing")

	return &MrsCatalogCabConnectorSRV{
		logger:       erylogrus.NewLogger(lg),
		storesClient: opts.StoresClient,
	}
}

// Connector Methods

func (srv *MrsCatalogCabConnectorSRV) GetStock(ctx context.Context, req *proto.GetStockRequest) (*proto.GetStockResult, error) {

	log.Info("GetStock")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	var stockResps []*proto.StockResponse
	switch {
	case req.Store != emptyString:

		// Append MRS context to Retail Stores Call
		ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

		// Convert Store Attrib to CNPJ
		storeInfo, err := srv.storesClient.GetStore(ctx, req.GetStore())
		if err != nil {
			return nil, err
		}

		// Call Stock service
		respBytes, err := srv.CallExternalService(req.Sku, catalog.StockResourceKey, storeInfo.ExternalCode)
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
			StoreName:   storeInfo.Name,
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

	return &proto.GetStockResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  stockResps,
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetStocks(ctx context.Context, req *proto.GetStocksRequest) (*proto.GetStocksResult, error) {

	log.Info("GetStocks")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to CNPJ
	storeInfo, err := srv.storesClient.GetStore(ctx, req.GetStoreId())
	if err != nil {
		return nil, err
	}

	skus := req.GetSkus()

	skuStocks := struct {
		data []*proto.GetStocksResponse
		mx   *sync.Mutex
	}{
		make([]*proto.GetStocksResponse, 0, len(skus)),
		&sync.Mutex{},
	}

	var wg sync.WaitGroup
	wg.Add(len(skus))

	var erygoErrs []*erygo.Err
	for i := range skus {

		go func(sku string) {

			defer wg.Done()

			respBytes, err := srv.CallExternalService(sku, catalog.StockResourceKey, storeInfo.ExternalCode)
			if err != nil {

				switch err := err.(type) {
				case *erygo.Err:
					erygoErrs = append(erygoErrs, err)
				}

				log.Warn("Error calling external service: " + err.Error())
				return
			}

			// Deserialize response
			var cliResp *ClientStockPriceResponse
			err = json.Unmarshal(respBytes, &cliResp)
			if err != nil {
				log.Warn("Error unmarshalling external response: " + err.Error())
				return
			}

			if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
				log.Warn("Empty response: " + err.Error())
				return
			}

			stockResp := &proto.GetStocksResponse{
				Sku:   sku,
				Stock: cliResp.ServiceResponse.ProductDetails.Estoque1,
			}

			skuStocks.mx.Lock()
			skuStocks.data = append(skuStocks.data, stockResp)
			skuStocks.mx.Unlock()

		}(skus[i])
	}

	wg.Wait()

	return &proto.GetStocksResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  skuStocks.data,
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetStoresStock(ctx context.Context, req *proto.GetStoresStockRequest) (*proto.GetStoresStockResult, error) {

	log.Info("GetStoresStock")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Fetch all store ids
	allStores, err := srv.storesClient.ListAllStores(ctx)
	if err != nil {
		return nil, err
	}

	// For every store, get stock
	storesStock := struct {
		data []*proto.StockResponse
		mx   *sync.Mutex
	}{
		make([]*proto.StockResponse, 0, len(allStores)),
		&sync.Mutex{},
	}

	var wg sync.WaitGroup
	wg.Add(len(allStores))

	var erygoErrs []*erygo.Err
	for i := range allStores {

		go func(store *stores.Store) {

			defer wg.Done()

			respBytes, err := srv.CallExternalService(req.Sku, catalog.StockResourceKey, store.ExternalCode)
			if err != nil {

				switch err := err.(type) {
				case *erygo.Err:
					erygoErrs = append(erygoErrs, err)
				}

				log.Warn("Error calling external service: " + err.Error())
				return
			}

			// Deserialize response
			var cliResp *ClientStockPriceResponse
			err = json.Unmarshal(respBytes, &cliResp)
			if err != nil {
				log.Warn("Error unmarshalling external response: " + err.Error())
				return
			}

			if cliResp == nil || cliResp.ServiceResponse.ProductDetails == nil {
				log.Warn("Empty response: " + err.Error())
				return
			}

			stockResp := &proto.StockResponse{
				Sku:         req.Sku,
				StoreId:     store.ID,
				StoreName:   store.Name,
				StockOnHand: cliResp.ServiceResponse.ProductDetails.Estoque1,
			}

			storesStock.mx.Lock()
			storesStock.data = append(storesStock.data, stockResp)
			storesStock.mx.Unlock()

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

	return &proto.GetStoresStockResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  storesStock.data,
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetPrice(ctx context.Context, req *proto.GetPriceRequest) (*proto.GetPriceResult, error) {

	log.Info("GetPrice")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to CNPJ
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.PriceResourceKey, storeInfo.ExternalCode)
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

	return &proto.GetPriceResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result: &proto.PriceResponse{
			Type:  erp,
			Sku:   req.Sku,
			Store: req.Store,
			Value: cliResp.ServiceResponse.ProductDetails.Preco,
		},
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetPricePOS(ctx context.Context, req *proto.GetPricePOSRequest) (*proto.GetPricePOSResult, error) {

	log.Info("GetPricePOS")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to CNPJ
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.PriceResourceKey, storeInfo.ExternalCode)
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

	return &proto.GetPricePOSResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result: &proto.PriceResponse{
			Type:  pos,
			Sku:   req.Sku,
			Store: req.Store,
			Value: cliResp.ServiceResponse.ProductDetails.Preco,
		},
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetPrices(ctx context.Context, req *proto.GetPricesRequest) (*proto.GetPricesResult, error) {

	log.Info("GetPrices")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to CNPJ
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.PriceResourceKey, storeInfo.ExternalCode)
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

	// Proto Prices Response
	ERPPrice := cliResp.ServiceResponse.ProductDetails.Preco
	POSPrice := cliResp.ServiceResponse.ProductDetails.PrecoPDV
	PricesResp := []*proto.PricesResponse{
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

	return &proto.GetPricesResult{
		Id:      mrsCtx.RequestID,
		Success: true,
		Message: http.StatusText(http.StatusOK),
		Result:  PricesResp,
	}, nil
}

func (srv *MrsCatalogCabConnectorSRV) GetLabel(ctx context.Context, req *proto.GetLabelRequest) (*proto.GetLabelResult, error) {

	log.Info("GetLabel")

	mrsCtx := mrs.GetMRSContext(ctx)
	log.Println(mrsCtx)

	// Append MRS context to Retail Stores Call
	ctx = grpcutil.AppendToOutgoingContext(ctx, mrs.ExtractMetadata(mrs.GetMRSContext(ctx)))

	// Convert Store Attrib to CNPJ
	storeInfo, err := srv.storesClient.GetStore(ctx, req.Store)
	if err != nil {
		return nil, err
	}

	// Call Price service
	respBytes, err := srv.CallExternalService(req.Sku, catalog.LabelResourceKey, storeInfo.ExternalCode)
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

func (srv *MrsCatalogCabConnectorSRV) GetSkuStoreSales(_ context.Context, _ *proto.GetSkuStoreSalesRequest) (*proto.GetSkuStoreSalesResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetDamagesDestinations(_ context.Context, _ *proto.GetDamagesDestinationsRequest) (*proto.GetDamagesDestinationsResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) GetDamagesReasons(_ context.Context, _ *proto.GetDamageReasonsRequest) (*proto.GetDamageReasonsResult, error) {
	return nil, apierr.ErrNotImplemented()
}

func (srv *MrsCatalogCabConnectorSRV) CallExternalService(product string, resourceId string, externalCode string) ([]byte, error) {

	extCfg := catalog.Get(nil).External
	uri := extCfg.BaseURI + "/" +
		extCfg.Resources[resourceId].Path + "/" +
		product + "/" +
		externalCode

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
