package sms_activate_go

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

type (
	Countries map[string]CountryInfo

	CountryInfo struct {
		ID           int    `json:"id"`           // country id
		RusName      string `json:"rus"`          // country name in Russian
		EngName      string `json:"eng"`          // country name in English
		ChnName      string `json:"chn"`          // country name in Chinese
		Visible      int    `json:"visible"`      // 0 - is not displayed on the site, 1 - is displayed
		Retry        int    `json:"retry"`        // 0 - repeated SMS is not available, 1 - is available
		Rent         int    `json:"rent"`         // 0 - country is not leased, 1 - is leased;
		MultiService int    `json:"multiService"` // 0 - country is not available for multiservice, 1- available
	}
)

const countriesAction = "getCountries"

// GetCountries returns all countries list
//
// Example
//
//	countries, err := client.GetCountries()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for k, v := range countries {
//		fmt.Printf("Country Code: %s. Country Name: %s. Can be rented: %d. Visible: %d\n", k, v.EngName, v.Rent, v.Visible)
//	}
func (act *SMSActivate) GetCountries() (Countries, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	countriesReq := baseRequest{
		APIKey: act.APIKey,
		Action: countriesAction,
	}
	val, err := query.Values(countriesReq)
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

	var countries Countries
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
