package bfapi

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/tarb/util/www"
)

const (
	// HTTPS only scheme
	scheme string = "https"

	// Hosts
	accountHost     string = "identitysso.betfair.com"
	accountHostCert string = "identitysso-cert.betfair.com"
	exchangeHost    string = "api.betfair.com"
	streamHost      string = "stream-api.betfair.com:443"

	// Supported login method paths
	certLogin string = "/api/certlogin/"
	login     string = "/api/login"
	logout    string = "/api/logout"
	keepAlive string = "/api/keepAlive/"

	// Supported account method paths
	getToken                          string = "/exchange/account/rest/v1.0/token/"
	createDeveloperAppKeys            string = "/exchange/account/rest/v1.0/createDeveloperAppKeys/"
	getDeveloperAppKeys               string = "/exchange/account/rest/v1.0/getDeveloperAppKeys/"
	getAccountFunds                   string = "/exchange/account/rest/v1.0/getAccountFunds/"
	getAccountDetails                 string = "/exchange/account/rest/v1.0/getAccountDetails/"
	listCurrencyRates                 string = "/exchange/account/rest/v1.0/listCurrencyRates/"
	getApplicationSubscriptionHistory string = "/exchange/account/rest/v1.0/getApplicationSubscriptionHistory/"
	listApplicationSubscriptionTokens string = "/exchange/account/rest/v1.0/listApplicationSubscriptionTokens/"
	getVendorClientID                 string = "/exchange/account/rest/v1.0/getVendorClientId/"
	updateApplicationSubscription     string = "/exchange/account/rest/v1.0/updateApplicationSubscription/"
	getApplicationSubscriptionToken   string = "/exchange/account/rest/v1.0/getApplicationSubscriptionToken/"
	activateApplicationSubscription   string = "/exchange/account/rest/v1.0/activateApplicationSubscription/"

	// Market menu
	getMarketMenuJSON string = "/exchange/betting/rest/v1/en/navigation/menu.json"

	// Supported exchange method paths
	listMarketCatalogue     string = "/exchange/betting/rest/v1.0/listMarketCatalogue/"
	listMarketBook          string = "/exchange/betting/rest/v1.0/listMarketBook/"
	listClearedOrders       string = "/exchange/betting/rest/v1.0/listClearedOrders/"
	listCurrentOrders       string = "/exchange/betting/rest/v1.0/listCurrentOrders/"
	listMarketTypes         string = "/exchange/betting/rest/v1.0/listMarketTypes/"
	listEventTypes          string = "/exchange/betting/rest/v1.0/listEventTypes/"
	listCompetitions        string = "/exchange/betting/rest/v1.0/listCompetitions/"
	listEvents              string = "/exchange/betting/rest/v1.0/listEvents/"
	listCountries           string = "/exchange/betting/rest/v1.0/listCountries/"
	listVenues              string = "/exchange/betting/rest/v1.0/listVenues/"
	cancelOrders            string = "/exchange/betting/rest/v1.0/cancelOrders/"
	placeOrders             string = "/exchange/betting/rest/v1.0/placeOrders/"
	replaceOrders           string = "/exchange/betting/rest/v1.0/replaceOrders/"
	updateOrders            string = "/exchange/betting/rest/v1.0/updateOrders/"
	listMarketProfitAndLoss string = "/exchange/betting/rest/v1.0/listMarketProfitAndLoss/"
)

//
const (
	DefaultDialTimeOut   = 8 * time.Second
	DefaultClientTimeOut = 24 * time.Second
)

//
type AuthType uint8

//
const (
	OAuthToken   AuthType = 1
	SessionToken AuthType = 2
)

func (at AuthType) String() string {
	switch at {
	case OAuthToken:
		return "Authorization"
	case SessionToken:
		return "X-Authentication"
	default:
		panic("no auth type on bfapi.Client")
	}
}

//
type Client struct {
	client      *www.Client
	certificate *tls.Certificate
	sem         *semaphore.Weighted
	appKey      string
	token       atomic.Value

	authType  AuthType
	authValue string
}

//
type Token struct {
	Token  string
	Logged Time
	Update Time
	VcID   string
	Sub    SubscriptionHistoryItem
}

//
func (t Token) ActiveSub() bool {
	return t.Sub.SubscriptionStatus == SubStatusActivated
}

//
func NewClient(appKey string, c *http.Client, cert *tls.Certificate) *Client {
	if c == nil {
		var tlsConf *tls.Config
		if cert != nil {
			tlsConf = &tls.Config{
				Certificates:       []tls.Certificate{*cert},
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateFreelyAsClient,
				Rand:               rand.Reader,
			}
		}

		c = &http.Client{
			Timeout: DefaultClientTimeOut,
			Transport: &http.Transport{
				TLSClientConfig: tlsConf,
				Dial: func(network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, time.Duration(DefaultDialTimeOut))
				},
				TLSHandshakeTimeout: DefaultDialTimeOut,
				MaxIdleConnsPerHost: 50,
			},
		}
	}

	client := www.New(c)

	// h.Set("Accept-Encoding", "gzip, deflate")
	nc := &Client{
		client:      client,
		certificate: cert,
		appKey:      appKey,
		sem:         semaphore.NewWeighted(1),
	}
	nc.token.Store(Token{})

	return nc
}

// Client returns the www/http client for making manual requests
func (c *Client) Client() *www.Client { return c.client }

//
func (c *Client) GetAuth() (string, string) { return c.authType.String(), c.authValue }

//
func (c *Client) SetAuth(t AuthType, v string) {
	c.authType = t
	if t == OAuthToken {
		c.authValue = "BEARER " + v
	} else {
		c.authValue = v
	}
}

//
func (c *Client) ClearAuth() {
	c.authType, c.authValue = 0, ""
}

//
type APIError struct {
	Faultcode   string `json:"faultcode"`
	Faultstring string `json:"faultstring"`
	Detail      struct {
		Exceptionname         string    `json:"exceptionname"`
		AccountAPINGException Exception `json:"AccountAPINGException"`
		APINGException        Exception `json:"APINGException"`
	} `json:"detail"`
}

//
type Exception struct {
	RequestUUID  string `json:"requestUUID"`
	ErrorCode    string `json:"errorCode"`
	ErrorDetails string `json:"errorDetails"`
}

func (e APIError) Error() string {
	if e.Detail.AccountAPINGException.ErrorCode != "" {
		return e.Detail.AccountAPINGException.ErrorCode
	} else if e.Detail.APINGException.ErrorCode != "" {
		return e.Detail.APINGException.ErrorCode
	} else {
		return e.Faultstring
	}
}

func statusToAPINGException(err error) error {
	if err == nil {
		return nil
	}

	serr, ok := err.(www.StatusError)
	if !ok {
		return err
	}

	if len(serr.Body) > 0 {
		var apiErr APIError

		jsonErr := json.Unmarshal(serr.Body, &apiErr)
		if jsonErr != nil {
			// return the original err, not the json err
			return err
		}

		return apiErr
	}

	return err
}
