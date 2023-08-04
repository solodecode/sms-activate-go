package sms_activate_go

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const getStatusAction = "getStatus"

func (act *SMSActivate) GetStatus(actID string) (string, error) {
	if len(actID) == 0 {
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         ErrBadLength,
		}
	}
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	statusReq := baseRequest{
		APIKey:       act.APIKey,
		Action:       getStatusAction,
		ActivationID: actID,
	}
	val, err := query.Values(statusReq)
	if err != nil {
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         fmt.Errorf("%w:%w", ErrEncoding, err),
		}
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         fmt.Errorf("%w:%w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         fmt.Errorf("%w:%w", ErrBodyRead, err),
		}
	}
	switch string(body) {
	case BadKey:
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         ErrBadKey,
		}
	case ErrorSQL:
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         ErrSQL,
		}
	case WrongActivationID:
		return "", RequestError{
			RequestName: getStatusAction,
			Err:         ErrWrongActivationID,
		}
	}
	return string(body), nil
}
