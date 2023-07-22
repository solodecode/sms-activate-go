package sms_activate_go

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

const getNumberAction = "getNumberV2"

type (
	GetNumberRequest struct {
		Service        string   `url:"service"`
		Country        string   `url:"country"`
		Forward        bool     `url:"forward,omitempty"`
		FreePrice      bool     `url:"freePrice,omitempty"`
		MaxPrice       float64  `url:"maxPrice,omitempty"`
		PhoneException []string `url:"phoneException,omitempty" del:","`
		Operator       []string `url:"operator,omitempty" del:","`
		Verification   bool     `url:"verification,omitempty"`
	}
	Number struct {
		ActivationID     string `json:"activationId"`
		PhoneNumber      string `json:"phoneNumber"`
		ActivationCost   string `json:"activationCost"`
		CountryCode      string `json:"countryCode"`
		CanGetAnotherSMS bool   `json:"canGetAnotherSms"`
		Time             string `json:"activationTime"`
		Operator         string `json:"activationOperator"`
	}
)

func (act *SMSActivate) GetNumber(request GetNumberRequest) (Number, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	val, err := query.Values(request)
	if err != nil {
		return Number{}, err
	}
	val.Add(apiKeyQuery, act.APIKey)
	val.Add(actionQuery, getNumberAction)
	req.URL.RawQuery = val.Encode()
	resp, err := act.httpClient.Do(req)
	if err != nil {
		return Number{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	switch string(body) {
	case NoBalance:
		return Number{}, ErrNoBalance
	case BadKey:
		return Number{}, ErrBadKey
	case ErrorSQL:
		return Number{}, ErrSQL
	case NoNumbers:
		return Number{}, ErrNoNumbers
	}
	var num Number
	err = json.Unmarshal(body, &num)
	if err != nil {
		return Number{}, err
	}
	return num, nil
}
