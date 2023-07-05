package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
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
	val, err := Values(baseReq)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = val.Encode()
	resp, err := act.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	data := make(TopCountriesList)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
