package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	numsStatusAction = "getNumbersStatus"
	RussiaID         = 0
	UkraineID        = 1
	KazakhstanID     = 2
)

func (act *SMSActivate) GetAvailableNumbers(country int, operator []string) (map[string]string, error) {
	if country < 0 || country > maxAvailableCountries {
		return nil, BadCountryNum
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	numsReq := baseRequest{
		APIKey:  apiKeyQuery,
		Action:  numsStatusAction,
		Country: country,
	}
	val, err := Values(numsReq)
	if err != nil {
		return nil, err
	}
	if country == RussiaID || country == UkraineID || country == KazakhstanID {
		var operators string
		operators = strings.Join(operator, ",")
		val.Add(operatorQuery, operators)
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
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
