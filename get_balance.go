package sms_activate_go

import (
	"errors"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	BalancePref   = "ACCESS_BALANCE:"
	balanceAction = "getBalance"
)

// GetBalance returns balance on your account.
//
// Example
//
//	balance, err := client.GetBalance()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Balance: %s", balance)
func (act *SMSActivate) GetBalance() (float64, error) {
	req, _ := http.NewRequest(http.MethodGet, act.BaseURL.String(), nil)

	balanceReq := baseRequest{
		APIKey: act.APIKey,
		Action: balanceAction,
	}
	val, err := query.Values(balanceReq)
	if err != nil {
		return 0, err
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	data := string(body)
	switch {
	case strings.HasPrefix(data, BalancePref):
		balance := strings.TrimPrefix(data, BalancePref)
		var fBal float64
		fBal, err = strconv.ParseFloat(balance, 64)
		if err != nil {
			return 0, err
		}
		return fBal, nil
	case data == BadKey:
		return 0, ErrBadKey
	case data == ErrorSQL:
		return 0, ErrSQL
	}
	return 0, errors.New("unknown response: " + data)
}
