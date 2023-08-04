package sms_activate_go

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
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
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         ErrBadLength,
		}
	}
	if status != ReadyStatus || status != NewSMSStatus || status != FinishActStatus || status != BadNumStatus {
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         ErrWrongStatus,
		}
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
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         fmt.Errorf("%w: %w", ErrBodyRead, err),
		}
	}

	switch string(body) {
	case BadKey:
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         ErrBadKey,
		}
	case ErrorSQL:
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         ErrSQL,
		}
	case WrongActivationID:
		return "", RequestError{
			RequestName: setStatusAction,
			Err:         ErrWrongActivationID,
		}
	}
	return string(body), nil
}
