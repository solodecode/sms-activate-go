package sms_activate_go

import (
	"errors"
	"fmt"
)

const (
	BadKey            = "BAD_KEY"
	ErrorSQL          = "ERROR_SQL"
	NoNumbers         = "NO_NUMBERS"
	NoBalance         = "NO_BALANCE"
	NoOperators       = `{"status":"error","error":"OPERATORS_NOT_FOUND"}`
	NoActivations     = `{"status":"error","error":"NO_ACTIVATIONS"}`
	WrongActivationID = "WRONG_ACTIVATION_ID"
	EarlyCancel       = "EARLY_CANCEL_DENIED"
)

type RequestError struct {
	RequestName string
	Err         error
}

func (r RequestError) Error() string {
	return fmt.Sprintf("(%s):%v", r.RequestName, r.Err)
}

var (
	ErrEncoding          = errors.New("error while encoding query")
	ErrWithReq           = errors.New("error with doing request")
	ErrBodyRead          = errors.New("error while reading body")
	ErrUnmarshalling     = errors.New("error while unmarshalling body")
	ErrBadLength         = errors.New("one or more params have wrong length")
	ErrBadLengthKey      = errors.New("bad length key")
	ErrBadKey            = errors.New("invalid API access key")
	BadCountryNum        = errors.New("the country number must be at least -1 and no more than 196")
	ErrSQL               = errors.New("one of the params has an invalid value")
	ErrNoNumbers         = errors.New("there are no free numbers for receiving SMS from this service")
	ErrNoBalance         = errors.New("not enough funds on current api key")
	ErrNoOperators       = errors.New("have no operators in this country")
	ErrNoActivations     = errors.New("have no active activations")
	ErrWrongActivationID = errors.New("wrong activation id")
	ErrWrongStatus       = errors.New("wrong status provided")
	ErrEarlyCancel       = errors.New("cannot cancel the number within the first 2 minutes")
)
