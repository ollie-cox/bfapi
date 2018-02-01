package bfapi

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/tarb/util/www"
)

//
type Config struct {
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
	AppKey   string `json:"appKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Testing  bool   `json:"testing"`
}

const (
	// HTTPS only scheme
	scheme string = "https"

	// Hosts
	testHost         string = "localhost"
	testStreamHost   string = ":8080"
	prodAccountHost  string = "identitysso.betfair.com"
	prodExchangeHost string = "api.betfair.com"
	prodStreamHost   string = "stream-api.betfair.com:443"

	// Supported login method paths
	certLogin string = "/api/certlogin/"
	login     string = "/api/login"
	keepAlive string = "/api/keepAlive/"

	// Supported account method paths
	getAccountFunds                   string = "/exchange/account/rest/v1.0/getAccountFunds/"
	getAccountDetails                 string = "/exchange/account/rest/v1.0/getAccountDetails/"
	getApplicationSubscriptionHistory string = "/exchange/account/rest/v1.0/getApplicationSubscriptionHistory/"
	getApplicationSubscriptionToken   string = "/exchange/account/rest/v1.0/getApplicationSubscriptionToken/"

	// Market menu
	getMarketMenuJSON string = "/exchange/betting/rest/v1/en/navigation/menu.json/"

	// Supported exchange method paths
	listMarketCatalogue string = "/exchange/betting/rest/v1.0/listMarketCatalogue/"
	listMarketBook      string = "/exchange/betting/rest/v1.0/listMarketBook/"
	listClearedOrders   string = "/exchange/betting/rest/v1.0/listClearedOrders/"
	cancelOrders        string = "/exchange/betting/rest/v1.0/cancelOrders/"
	placeOrders         string = "/exchange/betting/rest/v1.0/placeOrders/"
	replaceOrders       string = "/exchange/betting/rest/v1.0/replaceOrders/"
	updateOrders        string = "/exchange/betting/rest/v1.0/updateOrders/"
)

var (
	exchangeHost string
	accountHost  string
	streamHost   string

	certificate tls.Certificate
	appKey      string
	token       string
)

// Init creates a new http.Client, sets up default headers and configures
// which host to use
// certBytes, keyBytes - sed to make the certificate used in the http.Client
// appkey - default header value for X-Application <appkey>
// testing - sets which host is used in urls - production or testing
func Init(productAppKey, certFile, keyFile string) {
	var err error

	// store the vars
	appKey = productAppKey
	if certificate, err = tls.LoadX509KeyPair(certFile, keyFile); err != nil {
		log.Fatal("Error initilizing bfapi > ", err)
	}

	exchangeHost, accountHost, streamHost = prodExchangeHost, prodAccountHost, prodStreamHost

	// set up https client with supplied cert
	var transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{certificate},
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
		h.Set("X-Application", appKey)
		h.Set("Accept", "application/json")
		h.Set("Connection", "keep-alive")
	})
}

//
func TestingMode() {
	exchangeHost, accountHost, streamHost = testHost, testHost, testStreamHost
}
