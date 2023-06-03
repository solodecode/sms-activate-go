package sms_activate_go

import (
	"encoding/json"
	"io"
	"net/http"
)

const activationsAction = "getActiveActivations"

type (
	ActivationList struct {
		Activations []Activation `json:"activeActivations"`
	}
	Activation struct {
		ActivationID     string `json:"activationId"`
		ServiceCode      string `json:"serviceCode"`
		PhoneNumber      string `json:"phoneNumber"`
		ActivationCost   string `json:"activationCost"`
		ActivationStatus string `json:"activationStatus"`
		SMSCode          string `json:"smsCode"`
		SMSText          string `json:"smsText"`
		ActivationTime   string `json:"activationTime"`
		Discount         string `json:"discount"`
		Repeated         string `json:"repeated"`
		CountryCode      string `json:"countryCode"`
		CountryName      string `json:"countryName"`
		CanGetAnotherSMS string `json:"canGetAnotherSMS"`
	}
)

func (act *SMSActivate) GetActiveActivations() (ActivationList, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	activationsReq := baseRequest{
		APIKey: act.APIKey,
		Action: activationsAction,
	}
	val, err := Values(activationsReq)
	if err != nil {
		return ActivationList{}, err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return ActivationList{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	switch string(body) {
	case BadKey:
		return ActivationList{}, ErrBadKey
	case ErrorSQL:
		return ActivationList{}, ErrSQL
	case NoActivations:
		return ActivationList{}, ErrNoActivations
	}

	var data ActivationList
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ActivationList{}, err
	}
	return data, nil
}
