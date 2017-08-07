package bfapi

import (
	"net/http"

	"github.com/tarb/www"
)

//
func GetAccountFunds() (AccountFundsResponse, error) {
	var err error
	var result AccountFundsResponse

	err = www.Build(http.MethodPost, scheme, accountHost, getAccountFunds).
		CollectJSON(&result)

	return result, err
}

//
func GetAccountDetails() (AccountDetailsResponse, error) {
	var err error
	var result AccountDetailsResponse

	err = www.Build(http.MethodPost, scheme, accountHost, getAccountDetails).
		CollectJSON(&result)

	return result, err
}
