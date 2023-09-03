package sms_activate_go

import (
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const getStatusAction = "getStatus"

// GetStatus returns status of your activation.
//
// Example
//
//	data, err := client.GetActiveActivations()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	status, err := client.GetStatus(data.Activations[0].ActivationID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(status)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	switch string(body) {
	case badKeyMsg:
		return "", ErrBadKey
	case errorSQLMsg:
		return "", ErrSQL
	case wrongActivationIDMsg:
		return "", ErrWrongActivationID
	}
	return string(body), nil
}
