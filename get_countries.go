package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type (
	Countries map[string]CountryInfo

	CountryInfo struct {
		// country id
		ID int `json:"id"`

		// country name in Russian
		RusName string `json:"rus"`

		// country name in English
		EngName string `json:"eng"`

		// country name in Chinese
		ChnName string `json:"chn"`

		// 0 - is not displayed on the site, 1 - is displayed
		Visible int `json:"visible"`

		// 0 - repeated SMS is not available, 1 - is available
		Retry int `json:"retry"`

		// 0 - country is not leased, 1 - is leased;
		Rent int `json:"rent"`

		// 0 - country is not available for multiservice, 1- available
		MultiService int `json:"multiService"`
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch string(body) {
	case badKeyMsg:
		return nil, ErrBadKey
	case errorSQLMsg:
		return nil, ErrSQL
	}

	var countries Countries

	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
