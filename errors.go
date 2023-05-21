package sms_activate_go

import "errors"

const (
	BadKey        = "BAD_KEY"
	ErrorSQL      = "ERROR_SQL"
	NoNumbers     = "NO_NUMBERS"
	NoBalance     = "NO_BALANCE"
	NoOperators   = `{"status":"error","error":"OPERATORS_NOT_FOUND"}`
	NoActivations = `{"status":"error","error":"NO_ACTIVATIONS"}`
)

var (
	ErrBadLengthKey  = errors.New("bad length key")
	ErrBadKey        = errors.New("invalid API access key")
	BadCountryNum    = errors.New("the country number must be at least -1 and no more than 196 where -1 is all countries")
	ErrSQL           = errors.New("one of the params has an invalid value")
	ErrNoNumbers     = errors.New("there are no free numbers for receiving SMS from this service")
	ErrNoBalance     = errors.New("not enough funds on current api key")
	ErrNoOperators   = errors.New("have no operators in this country")
	ErrNoActivations = errors.New("have no active activations")
)
