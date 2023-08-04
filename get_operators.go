package sms_activate_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type CountryOperators struct {
	Operators map[string][]string `json:"countryOperators"`
}

const operatorsAction = "getOperators"

func (act *SMSActivate) GetOperators(country int) (CountryOperators, error) {
	if country < 0 || country > maxAvailableCountries {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         BadCountryNum,
		}
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	operatorsReq := baseRequest{
		APIKey:  act.APIKey,
		Action:  operatorsAction,
		Country: country,
	}
	val, err := query.Values(operatorsReq)
	if err != nil {
		return CountryOperators{}, RequestError{
			RequestName: operatorsAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
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
