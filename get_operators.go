package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

type CountryOperators struct {
	Operators map[string][]string `json:"countryOperators"`
}

const operatorsAction = "getOperators"

// GetOperators returns all operators available for the transferred country, if the country is -1 it will return all operators available for each country.
//
// Example
//
//	operators, err := client.GetOperators(-1)
//	if err != nil {
//	   log.Fatal(err)
//	}
//	for country, operators := range operators.Operators {
//	    fmt.Printf("Country: %s. Operators: %s\n", country, operators)
//	}
func (act *SMSActivate) GetOperators(country int) (CountryOperators, error) {
	if country < allCountries || country > maxAvailableCountries {
		return CountryOperators{}, BadCountryNum
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	operatorsReq := baseRequest{
		APIKey: act.APIKey,
		Action: operatorsAction,
	}
	val, err := query.Values(operatorsReq)
	if err != nil {
		return CountryOperators{}, err
	}
	if country > allCountries {
		val.Add(countryQuery, strconv.Itoa(country))
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return CountryOperators{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CountryOperators{}, err
	}

	switch string(body) {
	case badKeyMsg:
		return CountryOperators{}, ErrBadKey
	case errorSQLMsg:
		return CountryOperators{}, ErrSQL
	case noOperatorsMsg:
		return CountryOperators{}, ErrNoOperators
	}

	var data CountryOperators

	err = json.Unmarshal(body, &data)
	if err != nil {
		return CountryOperators{}, err
	}
	return data, nil
}
