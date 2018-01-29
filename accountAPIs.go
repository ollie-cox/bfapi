package bfapi

import (
	"net/http"

	"github.com/tarb/www"
)

//
func GetAccountFunds() (AccountFundsResponse, error) {
	var err error
	var result AccountFundsResponse

	err = www.Build(http.MethodPost, scheme, exchangeHost, getAccountFunds).
		WithHeaders(func(h http.Header) {
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		CollectJSON(&result)

	return result, err
}

//
func GetAccountDetails() (AccountDetailsResponse, error) {
	var err error
	var result AccountDetailsResponse

	err = www.Build(http.MethodPost, scheme, exchangeHost, getAccountDetails).
		WithHeaders(func(h http.Header) {
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		CollectJSON(&result)

	return result, err
}

//
func GetApplicationSubscriptionToken(arg GetSubTokenArg) (string, error) {
	var err error
	var result string

	result, err = www.Build(http.MethodPost, scheme, exchangeHost, getApplicationSubscriptionToken).
		WithJSONBody(arg).
		CollectString()

	return result, err
}

//
func GetApplicationSubscriptionHistory() (SubscriptionHistory, error) {
	var err error
	var result SubscriptionHistory

	err = www.Build(http.MethodPost, scheme, exchangeHost, getApplicationSubscriptionHistory).
		WithHeaders(func(h http.Header) {
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		CollectJSON(&result)

	return result, err
}
