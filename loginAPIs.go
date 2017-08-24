package bfapi

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tarb/www"
)

//
func CertLogin() (string, error) {
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
	if err == nil && result.LoginStatus != "SUCCESS" {
		return "", errors.New("CertLogin Failed with returned status: " + result.LoginStatus)
	}

	// store the session token and set as a default header
	token = result.SessionToken
	www.SetDefaultHeaders(func(h http.Header) {
		h.Set("X-Authentication", result.SessionToken)
	})

	// Not necessarily login success
	// result.SessionToken can be empty here with non nil err value
	return result.SessionToken, err
}

//
func KeepAlive(tick time.Duration) {
	type keepAliveResult struct {
		Token   string `json:"token"`
		Product string `json:"product"`
		Status  string `json:"status"`
		Error   string `json:"error"`
	}

	//sleep initially
	// time.Sleep(tick)

	const maxFails int = 5
	var numFails int

	for {
		var err error
		var result keepAliveResult

		err = www.Build(http.MethodGet, scheme, accountHost, keepAlive).
			CollectJSON(&result)

		// Connection worked but keepAlive was not successful
		if result.Status != "SUCCESS" {
			numFails++

			if numFails < maxFails {
				time.Sleep(time.Minute)
				continue
			}
			log.Fatal("Can not keep session alive", err)
		}

		numFails = 0
		time.Sleep(tick)
	}
}
