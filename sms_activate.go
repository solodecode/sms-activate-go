package sms_activate_go

import (
	"net/http"
	"net/url"
)

const (
	apiKeyLength          = 32
	apiBaseURL            = "https://api.sms-activate.org/stubs/handler_api.php"
	maxAvailableCountries = 196
	serviceCodeMinLength  = 2
)

const (
	apiKeyQuery         = "api_key"
	actionQuery         = "action"
	countryQuery        = "country"
	serviceQuery        = "service"
	operatorQuery       = "operator"
	freePriceQuery      = "freePrice"
	forwardQuery        = "forward"
	maxPriceQuery       = "maxPrice"
	phoneExceptionQuery = "phoneException"
	verificationQuery   = "verification"
	refCode             = "194015"
)

type (
	SMSActivate struct {
		APIKey     string
		BaseURL    *url.URL
		RefCode    string
		httpClient http.Client
	}
	baseRequest struct {
		APIKey       string `url:"api_key"`
		Action       string `url:"action"`
		Service      string `url:"service,omitempty"`
		Country      int    `url:"country,omitempty"`
		ActivationID string `url:"id,omitempty"`
		Status       int    `url:"status,omitempty"`
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
		RefCode:    refCode,
	}
	return act, nil
}
