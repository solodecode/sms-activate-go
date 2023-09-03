package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

const topCountriesAction = "getTopCountriesByService"

type (
	TopCountriesList map[string]TopCountriesInfo

	FreePriceMap map[float64]int

	TopCountriesInfo struct {
		FreePrice   FreePriceMap `json:"freePriceMap,omitempty"`
		CountryID   int          `json:"country"`
		Count       int          `json:"count"`
		Price       float64      `json:"price"`
		RetailPrice float64      `json:"retail_price"`
	}
)

func (m *FreePriceMap) UnmarshalJSON(data []byte) error {
	var stringMap map[string]int
	if err := json.Unmarshal(data, &stringMap); err != nil {
		return err
	}

	// string to float64 conversion
	fp := make(map[float64]int)
	for key, value := range stringMap {
		floatKey, err := strconv.ParseFloat(key, 64)
		if err != nil {
			return err
		}
		fp[floatKey] = value
	}

	*m = fp
	return nil
}

// GetTopCountries returns top countries by service.
//
// Example
//
//	topList, err := client.GetTopCountries("bz", true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for topPlace, info := range topList {
//	    fmt.Printf("Place in the top: %s, Country ID: %d, Count: %d, Price: %f\n", topPlace, info.CountryID, info.Count, info.Price)
//		for price, count := range info.FreePrice {
//			fmt.Printf("FreePrice: %f, FreePriceCount: %d\n", price, count)
//		}
//	}
func (act *SMSActivate) GetTopCountries(service string, freePrice bool) (TopCountriesList, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	baseReq := baseRequest{
		APIKey:    act.APIKey,
		Action:    topCountriesAction,
		Service:   service,
		FreePrice: freePrice,
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
