# ⚙️ sms-activate-go
A non-official library for working with the API sms-activate.org

## 🔧 Setup
• Create a new project and get the library: `go get github.com/solodecode/sms-activate-go`

• Get your API key from https://sms-activate.org/api2

• Use the examples from the Usage category to create an sms-activate object.

## 💡 Usage
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
	fmt.Printf("Balance: %f\n", balance)
}
```
By default, the object is created with the referral code of the author, but you can change it using the options:
```go
func main() {
	client, err := SMSActivate.New("yourapikey", SMSActivate.WithRefCode("yourrefcode"))
	if err != nil {
		log.Fatal(err)
	}
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %f\n", balance)
}
```
You can also create an sms-activate object with a custom http.Client:
```go
func main() {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	client, err := SMSActivate.New("yourapikey", SMSActivate.WithRefCode("yourrefcode"), SMSActivate.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatal(err)
	}
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %f\n", balance)
```
### 🔴 WARNING
Unfortunately, the server of the sms-activate service is very kind and always responds with 200 OK, so the returned responses may be unhandled errors.
