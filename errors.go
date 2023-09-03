package sms_activate_go

import (
	"errors"
)

const (
	badKeyMsg            = "BAD_KEY"
	errorSQLMsg          = "ERROR_SQL"
	noNumbersMsg         = "NO_NUMBERS"
	noBalanceMsg         = "NO_BALANCE"
	noOperatorsMsg       = `{"status":"error","error":"OPERATORS_NOT_FOUND"}`
	noActivationsMsg     = `{"status":"error","error":"NO_ACTIVATIONS"}`
	wrongActivationIDMsg = "WRONG_ACTIVATION_ID"
	earlyCancelMsg       = "EARLY_CANCEL_DENIED"
)

var (
	ErrBadLength         = errors.New("one or more params have wrong length")
	ErrBadLengthKey      = errors.New("bad length key")
	ErrUnknownResp       = errors.New("unknown response")
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
