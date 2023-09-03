package sms_activate_go

import (
	"encoding/json"
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

// GetTopCountries returns top countries by service.
//
// Example
//
//	topList, err := client.GetTopCountries("bz")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for topPlace, info := range topList {
//	    fmt.Printf("Place in the top: %s, Country ID: %d, Count: %d, Price: %f\n", topPlace, info.CountryID, info.Count, info.Price)
//	}
func (act *SMSActivate) GetTopCountries(service string) (TopCountriesList, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	baseReq := baseRequest{
		APIKey:  act.APIKey,
		Action:  topCountriesAction,
		Service: service,
	}
	val, err := query.Values(baseReq)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data TopCountriesList

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
