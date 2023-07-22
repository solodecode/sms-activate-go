package sms_activate_go

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

type (
	Info struct {
		Cost  float64 `json:"cost"`  // cost per number of service
		Count int     `json:"count"` // count numbers available for this service
	}
)

const pricesAction = "getPrices"

// GetPrices returns actual prices by country and service. Service may be empty if you need to get all services. Country may be -1 if you need to get all countries.
// You can get tables of services and countries at https://sms-activate.org/api2
//
// Example
//
//	data, err := client.GetPrices("", -1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for k, v := range data {
//		fmt.Println("Country:", k)
//		for k, v := range v {
//			fmt.Printf("Service: %s. Cost: %f, Count: %d\n", k, v.Cost, v.Count)
//		}
//	}
func (act *SMSActivate) GetPrices(service string, country int) (map[string]map[string]Info, error) {
	if country < -1 || country > maxAvailableCountries {
		return nil, BadCountryNum
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	pricesReq := baseRequest{
		APIKey:  act.APIKey,
		Action:  pricesAction,
		Service: service,
		Country: country,
	}

	val, err := query.Values(pricesReq)
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
	switch string(body) {
	case BadKey:
		return nil, ErrBadKey
	case ErrorSQL:
		return nil, ErrSQL
	}
	var data map[string]map[string]Info
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
