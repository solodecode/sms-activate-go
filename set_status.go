package sms_activate_go

import (
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
)

const (
	setStatusAction = "setStatus"
	ReadyStatus     = 1
	NewSMSStatus    = 3
	FinishActStatus = 6
	BadNumStatus    = 8
)

func (act *SMSActivate) SetStatus(id string, status int) (string, error) {
	if len(id) == 0 {
		return "", ErrBadLength
	}
	if status != ReadyStatus || status != NewSMSStatus || status != FinishActStatus || status != BadNumStatus {
		return "", ErrWrongStatus
	}
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	setStatusReq := baseRequest{
		APIKey:       act.APIKey,
		Action:       setStatusAction,
		ActivationID: id,
		Status:       status,
	}
	val, err := query.Values(setStatusReq)
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
