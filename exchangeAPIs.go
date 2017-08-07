package bfapi

import (
	"net/http"

	"github.com/tarb/www"
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
func PlaceOrders(ins []PlaceInstruction) (PlaceExecutionReport, error) {
	var err error
	var result PlaceExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, placeOrders).
		WithJSONBody(ins).
		CollectJSON(&result)

	return result, err
}

//
func CancelOrders(ins []CancelInstruction) (CancelExecutionReport, error) {
	var err error
	var result CancelExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, cancelOrders).
		WithJSONBody(ins).
		CollectJSON(&result)

	return result, err
}

//
func ReplaceOrders(ins []ReplaceInstruction) (ReplaceExecutionReport, error) {
	var err error
	var result ReplaceExecutionReport

	err = www.Build(http.MethodPost, scheme, exchangeHost, replaceOrders).
		WithJSONBody(ins).
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
