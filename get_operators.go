package sms_activate_go

import (
	"encoding/json"
	"fmt"
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
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         BadCountryNum,
		}
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	operatorsReq := baseRequest{
		APIKey: act.APIKey,
		Action: operatorsAction,
	}
	val, err := query.Values(operatorsReq)
	if err != nil {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
	}
	if country > allCountries {
		val.Add(countryQuery, strconv.Itoa(country))
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         fmt.Errorf("%w: %w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         fmt.Errorf("%w: %w", ErrBodyRead, err),
		}
	}

	switch string(body) {
	case BadKey:
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         ErrBadKey,
		}
	case ErrorSQL:
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         ErrSQL,
		}
	case NoOperators:
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         ErrNoOperators,
		}
	}

	var data CountryOperators
	err = json.Unmarshal(body, &data)
	if err != nil {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         fmt.Errorf("%w: %w", ErrUnmarshalling, err),
		}
	}
	return data, nil
}
