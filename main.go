package main

import (
	"fmt"
	"github.com/vprasad46/interview-accountapi/f3"
	logger "github.com/sirupsen/logrus"
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
		logger.Error("Form3 client creation got:%v, want:nil", err)
		os.Exit(1)
	}
	userDefinedData := []f3.Dictionary{
		f3.Dictionary{
			"key":   "UserKey",
			"value": "UserValue",
		},
		f3.Dictionary{
			"key":   "UserKey2",
			"value": 21,
		},
	}
	var validationType, referenceMask, acceptanceQualifier, country string = "card", "############", "same_day", "GB"
	accountAttributes := f3.AccountAttributes{
		AccountNumber:       "41426815",
		BankID:              "40020",
		BankIDCode:          "GBDSC",
		BaseCurrency:        "GBP",
		Bic:                 "NWBKGB22",
		Country:             &country,
		Name:                []string{"Vishwa", "Prasad"},
		ValidationType:      &validationType,
		ReferenceMask:       &referenceMask,
		AcceptanceQualifier: &acceptanceQualifier,
		UserDefinedData:     userDefinedData,
	}
	accountData := f3.AccountData{
		Attributes:     &accountAttributes,
		ID:             "fd27e265-9605-4b4b-a0e5-3003ea8dc4bc",
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
	fmt.Println(createAccountResponse.Data.ID)
}
