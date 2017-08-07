package bfapi

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/tarb/www"
)

//
func CertLogin(username, password string) (string, error) {
	type loginResult struct {
		SessionToken string `json:"sessionToken"`
		LoginStatus  string `json:"loginStatus"`
	}

	var err error
	var result loginResult

	err = www.Build(http.MethodPost, scheme, accountHost, certLogin).
		WithFormBody(func(form url.Values) {
			form.Set("username", username)
			form.Set("password", password)
		}).
		CollectJSON(&result)

	// Connection worked but login was not successful
	// Probably caused by username/password error
	if result.LoginStatus != "SUCCESS" {
		return "", errors.New("CertLogin Failed with returned status: " + result.LoginStatus)
	}

	// store the session token as a default header
	www.SetDefaultHeaders(func(h http.Header) {
		h.Set("X-Authentication", result.SessionToken)
	})

	// Not necessarily login success
	// result.SessionToken can be empty here with non nil err value
	return result.SessionToken, err
}

//
func KeepAlive() error {
	type keepAliveResult struct {
		Token   string `json:"token"`
		Product string `json:"product"`
		Status  string `json:"status"`
		Error   string `json:"error"`
	}

	var err error
	var result keepAliveResult

	err = www.Build(http.MethodGet, scheme, accountHost, keepAlive).
		CollectJSON(&result)

	// Connection worked but keepAlive was not successful
	if result.Status != "SUCCESS" {
		return errors.New("KeepAlive Failed with returned status: " + result.Status)
	}

	// Not necessarily KeepAlive success
	return err
}
