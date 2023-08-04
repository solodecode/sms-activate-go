package sms_activate_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	numsStatusAction = "getNumbersStatus"
	RussiaID         = 0
	UkraineID        = 1
	KazakhstanID     = 2
)

func (act *SMSActivate) GetAvailableNumbers(country int, operator []string) (map[string]string, error) {
	if country < 0 || country > maxAvailableCountries {
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         BadCountryNum,
		}
	}

	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	numsReq := baseRequest{
		APIKey:  apiKeyQuery,
		Action:  numsStatusAction,
		Country: country,
	}
	val, err := query.Values(numsReq)
	if err != nil {
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
	}

	if country == RussiaID || country == UkraineID || country == KazakhstanID {
		var operators string
		operators = strings.Join(operator, ",")
		val.Add(operatorQuery, operators)
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrBodyRead, err),
		}
	}
	switch string(body) {
	case BadKey:
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         ErrBadKey,
		}
	case ErrorSQL:
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         ErrSQL,
		}
	}
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, RequestError{
			RequestName: numsStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrUnmarshalling, err),
		}
	}
	return data, nil
}
