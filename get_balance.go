package sms_activate_go

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
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
		return 0, RequestError{
			RequestName: balanceAction,
			Err:         fmt.Errorf("%w: %w", ErrEncoding, err),
		}
	}
	req.URL.RawQuery = val.Encode()

	resp, err := act.httpClient.Do(req)
	if err != nil {
		return 0, RequestError{
			RequestName: balanceAction,
			Err:         fmt.Errorf("%w: %w", ErrWithReq, err),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, RequestError{
			RequestName: balanceAction,
			Err:         fmt.Errorf("%w: %w", ErrBodyRead, err),
		}
	}

	data := string(body)

	switch {
	case strings.HasPrefix(data, BalancePref):
		balance := strings.TrimPrefix(data, BalancePref)
		return strconv.ParseFloat(balance, 64)
	case data == BadKey:
		return 0, RequestError{
			RequestName: balanceAction,
			Err:         ErrBadKey,
		}
	case data == ErrorSQL:
		return 0, RequestError{
			RequestName: balanceAction,
			Err:         ErrSQL,
		}
	}

	return 0, RequestError{
		RequestName: balanceAction,
		Err:         fmt.Errorf("unknown response: %s", data),
	}
}
