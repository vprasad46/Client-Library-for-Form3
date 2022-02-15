# Form3 Go client Library for Accounts Resource

A simple go client library for the Form3 REST APIs (Accounts CREATE, FETCH and DELETE) actions.

### Environment variables

The following environment variables need to be set when constructing the client
using `form3.NewFromEnv`.

| Environment variable   | Description                                   |
|:-----------------------|:----------------------------------------------|
| `F3_BASE_URL`          | Form3 host URL, e.g. https://api.form3.tech   |
| `LOG_LEVEL`(optional)  | set to DEBUG to debug else logs above INFO    |


### Usage
```bash
export F3_BASE_URL="https://api.form3.tech"
```

```go
import "github.com/vprasad46/interview-accountapi/f3"

client, err := f3.NewClient()
if err != nil {
	logger.Error("Form3 client creation got:", err)
	os.Exit(1)
}
```

###Supported Methods

`func (apiClient *Client) CreateAccount(createAccountRequest *CreateAccountRequest) (*CreateAccountResponse, error)` - This function is used to create `Account` using `client` object

Example:
```go
var country string = "GB"
accountAttributes := f3.AccountAttributes{
	AccountNumber: "41426815",
	BankID:        "40020",
	BankIDCode:    "GBDSC",
	BaseCurrency:  "GBP",
	Bic:           "NWBKGB22",
	Country:       &country,
	Name:          []string{"Vishwa", "Prasad"},
}
accountData := f3.AccountData{
	Attributes:     &accountAttributes,
	ID:             "9aa1dcb1-eac9-43cc-a58c-02e07e7b752a",
	OrganisationID: "eb0bd6f5-c3f5-44b3-c677-acd23cdde73c",
	Type:           "accounts",
}
                                                                          
createAccountRequest := f3.CreateAccountRequest{
	Data: &accountData,
}
createAccountResponse, err := client.CreateAccount(&createAccountRequest)
```

`func (apiClient *Client) FetchAccount(accountID string) (*FetchAccountResponse, error)` - This function is used to fetch `Account Details` using `client` object

Example:
```go
fetchAccountResponse, err := client.FetchAccount("9ea1dcb1-eac9-43cc-a58c-02e07e7b752a")
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
```
`func (apiClient *Client) DeleteAccount(accountID string, version int) error` - This function is used to delete `Account` using `client` object

Example
```go
err = client.DeleteAccount("9ea1dcb1-eac9-43cc-a58c-02e07e7b752a", 0)
if err != nil {
	fmt.Println(err)
}
```

### REST API Documentation

For further details on the API behind client please visit: https://api-docs.form3.tech/api.html




