package bfapi

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tarb/util/www"
)

//
func CertLogin(username, password string) (string, error) {
	var err error
	var result CertLoginResult

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
func Login(username, password string) (string, error) {
	var err error
	var result LoginResult

	err = www.Build(http.MethodPost, scheme, accountHost, login).
		WithFormBody(func(form url.Values) {
			form.Set("username", username)
			form.Set("password", password)
		}).
		CollectJSON(&result)

	// Connection worked but login was not successful
	// Probably caused by username/password error
	if err != nil {
		return "", err
	} else if result.Status != Success {
		return "", LoginError{Status: result.Status, Err: result.Error}
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
	const maxFails int = 5
	var numFails int

	for {
		var err error
		var result LoginResult

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
