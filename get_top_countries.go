package sms_activate_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	topCountriesAction = "getTopCountriesByService"
)

type (
	TopCountriesList map[string]TopCountriesInfo
	TopCountriesInfo struct {
		CountryID   int     `json:"country"`
		Count       int     `json:"count"`
		Price       float64 `json:"price"`
		RetailPrice float64 `json:"retail_price"`
	}
)

func (act *SMSActivate) GetTopCountries(service string) (TopCountriesList, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	baseReq := baseRequest{
		APIKey:  act.APIKey,
		Action:  topCountriesAction,
		Service: service,
	}
	val, err := query.Values(baseReq)
	if err != nil {
		return nil, RequestError{
			RequestName: topCountriesAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
	}
	req.URL.RawQuery = val.Encode()
	resp, err := act.httpClient.Do(req)
	if err != nil {
		return nil, RequestError{
			RequestName: topCountriesAction,
			Err:         fmt.Errorf("%w: %w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, RequestError{
			RequestName: topCountriesAction,
			Err:         fmt.Errorf("%w: %w", ErrBodyRead, err),
		}
	}

	data := make(TopCountriesList)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, RequestError{
			RequestName: topCountriesAction,
			Err:         fmt.Errorf("%w: %w", ErrUnmarshalling, err),
		}
	}
	return data, nil
}
