package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type CountryOperators struct {
	Operators map[string][]string `json:"countryOperators"`
}

const operatorsAction = "getOperators"

func (act *SMSActivate) GetOperators(country int) (CountryOperators, error) {
	if country < -1 || country > maxAvailableCountries {
		return CountryOperators{}, BadCountryNum
	}
	query := map[string]string{
		apiKeyQuery: act.APIKey,
		actionQuery: operatorsAction,
	}
	if country > -1 {
		query[countryQuery] = strconv.Itoa(country)
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

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
