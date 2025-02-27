package polygon

import (
	"context"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
	"net/http"
)

const (
	ListFuturesAggsPath             = "/futures/vX/aggs/{ticker}"
	ListFuturesContractsPath        = "/futures/vX/contracts"
	GetFuturesContractPath          = "/futures/vX/contracts/{ticker}"
	ListFuturesMarketStatusesPath   = "/futures/vX/market-status"
	ListFuturesProductsPath         = "/futures/vX/products"
	GetFuturesProductPath           = "/futures/vX/products/{product_code}"
	ListFuturesSchedulesPath        = "/futures/vX/schedules"
	ListFuturesProductSchedulesPath = "/futures/vX/products/{product_code}/schedules"
	ListFuturesTradesPath           = "/futures/vX/trades/{ticker}"
	ListFuturesQuotesPath           = "/futures/vX/quotes/{ticker}"
)

// FuturesClient provides methods for interacting with the futures REST API.
type FuturesClient struct {
	client.Client
}

// ListFuturesAggs retrieves a list of aggregates for a futures contract.
func (fc *FuturesClient) ListFuturesAggs(ctx context.Context, params *models.ListFuturesAggsParams, options ...models.RequestOption) *iter.Iter[models.FuturesAggregate] {
	return iter.NewIter(ctx, ListFuturesAggsPath, params, func(uri string) (iter.ListResponse, []models.FuturesAggregate, error) {
		res := &models.ListFuturesAggsResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListFuturesContracts retrieves a list of futures contracts.
func (fc *FuturesClient) ListFuturesContracts(ctx context.Context, params *models.ListFuturesContractsParams, options ...models.RequestOption) *iter.Iter[models.FuturesContract] {
	return iter.NewIter(ctx, ListFuturesContractsPath, params, func(uri string) (iter.ListResponse, []models.FuturesContract, error) {
		res := &models.ListFuturesContractsResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetFuturesContract retrieves details for a specific futures contract.
func (fc *FuturesClient) GetFuturesContract(ctx context.Context, params *models.GetFuturesContractParams, options ...models.RequestOption) (*models.GetFuturesContractResponse, error) {
	res := &models.GetFuturesContractResponse{}
	err := fc.Call(ctx, http.MethodGet, GetFuturesContractPath, params, res, options...)
	return res, err
}

// ListFuturesMarketStatuses retrieves market statuses for futures products.
func (fc *FuturesClient) ListFuturesMarketStatuses(ctx context.Context, params *models.ListFuturesMarketStatusesParams, options ...models.RequestOption) *iter.Iter[models.FuturesMarketStatus] {
	return iter.NewIter(ctx, ListFuturesMarketStatusesPath, params, func(uri string) (iter.ListResponse, []models.FuturesMarketStatus, error) {
		res := &models.ListFuturesMarketStatusesResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListFuturesProducts retrieves a list of futures products.
func (fc *FuturesClient) ListFuturesProducts(ctx context.Context, params *models.ListFuturesProductsParams, options ...models.RequestOption) *iter.Iter[models.FuturesProduct] {
	return iter.NewIter(ctx, ListFuturesProductsPath, params, func(uri string) (iter.ListResponse, []models.FuturesProduct, error) {
		res := &models.ListFuturesProductsResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetFuturesProduct retrieves details for a specific futures product.
func (fc *FuturesClient) GetFuturesProduct(ctx context.Context, params *models.GetFuturesProductParams, options ...models.RequestOption) (*models.GetFuturesProductResponse, error) {
	res := &models.GetFuturesProductResponse{}
	err := fc.Call(ctx, http.MethodGet, GetFuturesProductPath, params, res, options...)
	return res, err
}

// ListFuturesSchedules retrieves trading schedules for futures.
func (fc *FuturesClient) ListFuturesSchedules(ctx context.Context, params *models.ListFuturesSchedulesParams, options ...models.RequestOption) *iter.Iter[models.FuturesSchedule] {
	return iter.NewIter(ctx, ListFuturesSchedulesPath, params, func(uri string) (iter.ListResponse, []models.FuturesSchedule, error) {
		res := &models.ListFuturesSchedulesResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListFuturesProductSchedules retrieves trading schedules for a specific futures product.
func (fc *FuturesClient) ListFuturesProductSchedules(ctx context.Context, params *models.ListFuturesProductSchedulesParams, options ...models.RequestOption) *iter.Iter[models.FuturesSchedule] {
	return iter.NewIter(ctx, ListFuturesProductSchedulesPath, params, func(uri string) (iter.ListResponse, []models.FuturesSchedule, error) {
		res := &models.ListFuturesProductSchedulesResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListFuturesTrades retrieves a list of trades for a futures contract.
func (fc *FuturesClient) ListFuturesTrades(ctx context.Context, params *models.ListFuturesTradesParams, options ...models.RequestOption) *iter.Iter[models.FuturesTrade] {
	return iter.NewIter(ctx, ListFuturesTradesPath, params, func(uri string) (iter.ListResponse, []models.FuturesTrade, error) {
		res := &models.ListFuturesTradesResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListFuturesQuotes retrieves a list of quotes for a futures contract.
func (fc *FuturesClient) ListFuturesQuotes(ctx context.Context, params *models.ListFuturesQuotesParams, options ...models.RequestOption) *iter.Iter[models.FuturesQuote] {
	return iter.NewIter(ctx, ListFuturesQuotesPath, params, func(uri string) (iter.ListResponse, []models.FuturesQuote, error) {
		res := &models.ListFuturesQuotesResponse{}
		err := fc.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
