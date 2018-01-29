package bfapi

import (
	"io/ioutil"
	"net/http"

	"github.com/tarb/util/www"
)

//
func ListMarketCatalogue(lmc ListMarketCatalogueArg) ([]MarketCatalogue, error) {
	var err error
	var result []MarketCatalogue

	err = www.Build(http.MethodPost, scheme, exchangeHost, listMarketCatalogue).
		WithJSONBody(lmc).
		CollectJSON(&result)

	return result, err
}

//
func ListMarketBook(lmb ListMarketBookArg) ([]MarketBook, error) {
	var err error
	var result []MarketBook

	err = www.Build(http.MethodPost, scheme, exchangeHost, listMarketBook).
		WithJSONBody(lmb).
		CollectJSON(&result)

	return result, err
}

//
func PlaceOrders(arg PlaceOrderArg) (PlaceExecutionReport, error) {
	var err error
	var result PlaceExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, placeOrders).
		WithJSONBody(arg).
		CollectJSON(&result)

	return result, err
}

//
func CancelOrders(req CancelOrderRequest) (CancelExecutionReport, error) {
	var err error
	var result CancelExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, cancelOrders).
		WithJSONBody(req).
		CollectJSON(&result)

	return result, err
}

//
func ReplaceOrders(req ReplaceOrderRequest) (ReplaceExecutionReport, error) {
	var err error
	var result ReplaceExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, replaceOrders).
		WithJSONBody(req).
		CollectJSON(&result)

	return result, err
}

//
func UpdateOrders(ins []UpdateInstruction) (UpdateExecutionReport, error) {
	var err error
	var result UpdateExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, updateOrders).
		WithJSONBody(ins).
		CollectJSON(&result)

	return result, err
}

//
func ListClearedOrders(arg ListClearedOrdersArgs) (ClearedOrderSummaryReport, error) {
	var err error
	var result ClearedOrderSummaryReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, listClearedOrders).
		WithJSONBody(arg).
		CollectJSON(&result)

	return result, err
}

// GetMarketMenuJSON   Collects JSON Market Menu from Betfair API
func GetMarketMenuJSON() ([]byte, error) {
	var err error
	var result []byte
	var resp *http.Response

	resp, err = www.Build(http.MethodGet, scheme, exchangeHost, getMarketMenuJSON).
		CollectResponse()

	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return result, err
}
