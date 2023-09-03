package sms_activate_go

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	setStatusAction    = "setStatus"
	ReadyStatus        = 1
	NewSMSStatus       = 3
	FinishActStatus    = 6
	BadNumStatus       = 8
	NumReadyStatusMsg  = "ACCESS_READY"
	NewSMSStatusMsg    = "ACCESS_RETRY_GET"
	SuccessStatusMsg   = "ACCESS_ACTIVATION"
	CancelledStatusMsg = "ACCESS_CANCEL"
)

// SetStatus sets the phone number activation status and returns true if the status is set.
//
// Example
//
//	num, err := client.GetNumber(SMSActivate.GetNumberRequest{
//	    Service: "ig",
//	    Country: "6",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	sent, err := client.SetStatus(num.ActivationID, SMSActivate.BadNumStatus)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if sent {
//	    // ...
//	}
func (act *SMSActivate) SetStatus(id string, status int) (bool, error) {
	if len(id) == 0 {
		return false, ErrBadLength
	}
	if status != ReadyStatus && status != NewSMSStatus && status != FinishActStatus && status != BadNumStatus {
		return false, ErrWrongStatus
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
		return false, err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := string(body)

	switch data {
	case badKeyMsg:
		return false, ErrBadKey
	case errorSQLMsg:
		return false, ErrSQL
	case wrongActivationIDMsg:
		return false, ErrWrongActivationID
	case earlyCancelMsg:
		return false, ErrEarlyCancel
	case NumReadyStatusMsg, NewSMSStatusMsg, SuccessStatusMsg, CancelledStatusMsg:
		return true, nil

	}
	return false, fmt.Errorf("%w: %s", ErrUnknownResp, data)
}
