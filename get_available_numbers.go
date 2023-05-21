package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	numsStatusAction = "getNumbersStatus"
	RussiaID         = 0
	UkraineID        = 1
	KazakhstanID     = 2
)

func (act *SMSActivate) GetAvailableNumbers(country int, operator []string) (map[string]string, error) {
	if country < -1 || country > maxAvailableCountries {
		return nil, BadCountryNum
	}
	query := map[string]string{
		apiKeyQuery: act.APIKey,
		actionQuery: numsStatusAction,
	}
	if country > -1 {
		query[countryQuery] = strconv.Itoa(country)
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	q := req.URL.Query()
	if country == RussiaID || country == UkraineID || country == KazakhstanID {
		var operators string
		operators = strings.Join(operator, ",")
		q.Add(operatorQuery, operators)
	}
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

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
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
