package sms_activate_go

import (
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

const getStatusAction = "getStatus"

func (act *SMSActivate) GetStatus(actID string) (string, error) {
	if len(actID) == 0 {
		return "", ErrBadLength
	}
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	statusReq := baseRequest{
		APIKey:       act.APIKey,
		Action:       getStatusAction,
		ActivationID: actID,
	}
	val, err := query.Values(statusReq)
	if err != nil {
		return "", err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	switch string(body) {
	case BadKey:
		return "", ErrBadKey
	case ErrorSQL:
		return "", ErrSQL
	case WrongActivationID:
		return "", ErrWrongActivationID
	}
	return string(body), nil
}
