# sms-activate-go
A non-official library for working with the API sms-activate.org
## Usage
```go
func main() {
	client, err := SMSActivate.New("yourapikey")
	if err != nil {
		log.Fatal(err)
	}
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %f", balance)
}
```
### WARNING
Unfortunately, the server of the sms-activate service is very kind and always responds with 200 OK, so the returned responses may be unhandled errors.
