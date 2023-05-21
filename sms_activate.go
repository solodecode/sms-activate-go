package sms_activate_go

import (
	"net/http"
	"net/url"
)

const (
	apiKeyLength          = 32
	apiBaseURL            = "https://api.sms-activate.org/stubs/handler_api.php"
	maxAvailableCountries = 196
)

const (
	apiKeyQuery   = "api_key"
	actionQuery   = "action"
	countryQuery  = "country"
	serviceQuery  = "service"
	operatorQuery = "operator"
)

type (
	SMSActivate struct {
		APIKey     string
		BaseURL    *url.URL
		httpClient http.Client
	}
)

func New(apikey string) (*SMSActivate, error) {
	if len(apikey) != apiKeyLength {
		return nil, ErrBadLengthKey
	}
	baseURL, _ := url.Parse(apiBaseURL)
	act := &SMSActivate{
		APIKey:     apikey,
		BaseURL:    baseURL,
		httpClient: http.Client{},
	}
	return act, nil
}
