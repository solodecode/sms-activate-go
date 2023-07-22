package sms_activate_go

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

type CountryOperators struct {
	Operators map[string][]string `json:"countryOperators"`
}

const operatorsAction = "getOperators"

func (act *SMSActivate) GetOperators(country int) (CountryOperators, error) {
	if country < 0 || country > maxAvailableCountries {
		return CountryOperators{}, BadCountryNum
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	operatorsReq := baseRequest{
		APIKey:  act.APIKey,
		Action:  operatorsAction,
		Country: country,
	}
	val, err := query.Values(operatorsReq)
	if err != nil {
		return CountryOperators{}, err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return CountryOperators{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	switch string(body) {
	case BadKey:
		return CountryOperators{}, ErrBadKey
	case ErrorSQL:
		return CountryOperators{}, ErrSQL
	case NoOperators:
		return CountryOperators{}, ErrNoOperators
	}
	var data CountryOperators
	err = json.Unmarshal(body, &data)
	if err != nil {
		return CountryOperators{}, err
	}
	return data, nil
}
