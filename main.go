package main

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/vprasad46/interview-accountapi/f3"
	"os"
)

func init() {
	lvl := os.Getenv("LOG_LEVEL")
	formatter := &logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger.SetFormatter(formatter)
	if lvl == "DEBUG" {
		logger.SetLevel(logger.DebugLevel)
	}
}

func main() {
	client, err := f3.NewClient()
	if err != nil {
		logger.Error("Form3 client creation got:", err)
		os.Exit(1)
	}
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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fetchAccountResponse, err := client.FetchAccount("9ea1dcb1-eac9-43cc-a58c-02e07e7b752a")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.DeleteAccount("9ea1dcb1-eac9-43cc-a58c-02e07e7b752a", 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fetchAccountResponse.Data.ID)
	fmt.Println(createAccountResponse.Data.ID)
}
