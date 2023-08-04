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
	operatorQuery = "operator"
	refCode       = "194015"
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
	option func(*SMSActivate)
)

func WithNonDefURL(newURL *url.URL) option {
	return func(act *SMSActivate) {
		act.BaseURL = newURL
	}
}

func WithRefCode(code string) option {
	return func(act *SMSActivate) {
		act.RefCode = code
	}
}

func WithHTTPClient(client http.Client) option {
	return func(act *SMSActivate) {
		act.httpClient = client
	}
}

func New(apikey string, option ...option) (*SMSActivate, error) {
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
	for _, opt := range option {
		opt(act)
	}
	return act, nil
}
