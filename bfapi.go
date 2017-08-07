package bfapi

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/tarb/www"
)

const (
	// HTTPS only scheme
	scheme string = "https"

	// Hosts
	testHost         string = "localhost"
	prodAccountHost  string = "identitysso.betfair.com"
	prodExchangeHost string = "api.betfair.com"

	// Supported login method paths
	certLogin string = "api/certlogin"
	keepAlive string = "api/keepAlive"

	// Supported account method paths
	getAccountFunds   string = "exchange/account/rest/v1.0/getAccountFunds"
	getAccountDetails string = "exchange/account/rest/v1.0/getAccountDetails"

	// Supported exchange method paths
	listMarketCatalogue string = "exchange/betting/rest/v1.0/listMarketCatalogue"
	cancelOrders        string = "exchange/betting/rest/v1.0/cancelOrders"
	placeOrders         string = "exchange/betting/rest/v1.0/placeOrders"
	replaceOrders       string = "exchange/betting/rest/v1.0/replaceOrders"
	updateOrders        string = "exchange/betting/rest/v1.0/updateOrders"
)

var exchangeHost string
var accountHost string

// Init creates a new http.Client, sets up default headers and configures
// which host to use
// certBytes, keyBytes - sed to make the certificate used in the http.Client
// appkey - default header value for X-Application <appkey>
// testing - sets which host is used in urls - production or testing
func Init(certBytes, keyBytes []byte, appkey string, testing bool) {
	var err error
	var cert tls.Certificate

	if cert, err = tls.X509KeyPair(certBytes, keyBytes); err != nil {
		log.Fatal("Error initilizing bfapi: ", err)
	}

	var transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateFreelyAsClient,
			Rand:               rand.Reader,
		},
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Duration(time.Second*3))
		},
		MaxIdleConnsPerHost: 50,
	}

	www.SetClient(&http.Client{Transport: transport})

	// X-Authentication SessionToken is set in defaults after successful certLogin
	www.SetDefaultHeaders(func(h http.Header) {
		h.Set("X-Application", appkey)
		h.Set("Accept", "application/json")
		h.Set("Connection", "keep-alive")
	})

	if testing {
		exchangeHost, accountHost = testHost, testHost
	} else {
		exchangeHost, accountHost = prodExchangeHost, prodAccountHost
	}
}
