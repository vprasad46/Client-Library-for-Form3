package main

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vprasad46/interview-accountapi/f3"
	"os"
	"testing"
)

func NewRequestObject() f3.CreateAccountRequest {
	country, accountClassification, status := "GB", "Personal", "confirmed"
	accountMatchingOptOut, jointAccount, switched := false, false, false
	accountAttributes := f3.AccountAttributes{
		AccountNumber:           "41426815",
		BankID:                  "40020",
		BankIDCode:              "GBDSC",
		BaseCurrency:            "GBP",
		Bic:                     "NWBKGB22",
		Country:                 &country,
		Name:                    []string{"Vishwa"},
		AlternativeNames:        []string{"Vishwa1", "Prasad1"},
		Iban:                    "GB11NWBK40030041426819",
		JointAccount:            &jointAccount,
		AccountClassification:   &accountClassification,
		SecondaryIdentification: "ASD",
		AccountMatchingOptOut:   &accountMatchingOptOut,
		Status:                  &status,
		Switched:                &switched,
	}
	accountData := f3.AccountData{
		Attributes:     &accountAttributes,
		ID:             uuid.New().String(),
		OrganisationID: "eb0bd6f5-c3f5-44b3-c677-acd23cdde73c",
		Type:           "accounts",
	}

	return f3.CreateAccountRequest{
		Data: &accountData,
	}
}

func TestF3ClientCreation(t *testing.T) {
	if os.Getenv("F3_BASE_URL") == "" {
		t.Skip("WARN: F3_BASE_URL not set, skipping test")
	}
	_, err := f3.NewClient()
	require.NoError(t, err)
}

func TestAccountCreation(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	createAccountRequest := NewRequestObject()
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)
	assert.Equal(t, createAccountRequest.Data.ID, createAccountResponse.Data.ID)
	assert.Equal(t, createAccountRequest.Data.OrganisationID, createAccountResponse.Data.OrganisationID)
	assert.Equal(t, createAccountRequest.Data.Type, createAccountResponse.Data.Type)
	assert.Equal(t, createAccountRequest.Data.Attributes.AccountNumber, createAccountResponse.Data.Attributes.AccountNumber)
	assert.Equal(t, createAccountRequest.Data.Attributes.BankID, createAccountResponse.Data.Attributes.BankID)
	assert.Equal(t, createAccountRequest.Data.Attributes.BankIDCode, createAccountResponse.Data.Attributes.BankIDCode)
	assert.Equal(t, createAccountRequest.Data.Attributes.BaseCurrency, createAccountResponse.Data.Attributes.BaseCurrency)
	assert.Equal(t, createAccountRequest.Data.Attributes.Bic, createAccountResponse.Data.Attributes.Bic)
	assert.Equal(t, *createAccountRequest.Data.Attributes.Country, *createAccountResponse.Data.Attributes.Country)
	assert.Equal(t, createAccountRequest.Data.Attributes.Name, createAccountResponse.Data.Attributes.Name)
	assert.Equal(t, *createAccountRequest.Data.Attributes.AccountClassification, *createAccountResponse.Data.Attributes.AccountClassification)
	assert.Equal(t, *createAccountRequest.Data.Attributes.AccountMatchingOptOut, *createAccountResponse.Data.Attributes.AccountMatchingOptOut)
	assert.Equal(t, createAccountRequest.Data.Attributes.AlternativeNames, createAccountResponse.Data.Attributes.AlternativeNames)
	assert.Equal(t, createAccountRequest.Data.Attributes.Iban, createAccountResponse.Data.Attributes.Iban)
	assert.Equal(t, *createAccountRequest.Data.Attributes.JointAccount, *createAccountResponse.Data.Attributes.JointAccount)
	assert.Equal(t, createAccountRequest.Data.Attributes.SecondaryIdentification, createAccountResponse.Data.Attributes.SecondaryIdentification)
	assert.Equal(t, *createAccountRequest.Data.Attributes.Status, *createAccountResponse.Data.Attributes.Status)
	assert.Equal(t, *createAccountRequest.Data.Attributes.Switched, *createAccountResponse.Data.Attributes.Switched)

	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)

}

func TestAccountCreationWithOnlyRequiredAttributes(t *testing.T) {

	client, err := f3.NewClient()
	require.Nil(t, err)

	var country string = "GB"
	accountAttributes := f3.AccountAttributes{
		Country: &country,
		Name:    []string{"Vishwa", "Prasad"},
	}
	accountData := f3.AccountData{
		Attributes:     &accountAttributes,
		ID:             "8aa1dcb1-eac9-43cc-a58c-02e07e7b752c",
		OrganisationID: "eb0bd6f5-c3f5-44b3-c677-acd23cdde73c",
		Type:           "accounts",
	}

	createAccountRequest := f3.CreateAccountRequest{
		Data: &accountData,
	}
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)

	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)
}

func TestAccountCreationWithoutRequiredAttributes(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)
	/*
	   Creating Account with Empty Account Data
	*/
	accountData := f3.AccountData{}
	createAccountRequest := f3.CreateAccountRequest{
		Data: &accountData,
	}
	_, err = client.CreateAccount(&createAccountRequest)
	require.Error(t, err)
	/*
	   AccountData with Empty AccountAttributes
	*/
	accountAttributes := f3.AccountAttributes{}
	accountData = f3.AccountData{
		Attributes: &accountAttributes,
	}
	createAccountRequest = f3.CreateAccountRequest{
		Data: &accountData,
	}
	_, err = client.CreateAccount(&createAccountRequest)
	require.Error(t, err)
}

func TestDuplicateAccountCreation(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	createAccountRequest := NewRequestObject()
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)
	_, err = client.CreateAccount(&createAccountRequest)
	require.Error(t, err)

	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)
}

func TestAccountFetch(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	createAccountRequest := NewRequestObject()
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)

	fetchAccountResponse, err := client.FetchAccount(createAccountRequest.Data.ID)
	require.Nil(t, err)

	assert.Equal(t, fetchAccountResponse.Data.ID, createAccountRequest.Data.ID)
	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)
}

func TestFailedAccountFetch(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	createAccountRequest := NewRequestObject()
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)

	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)

	_, err = client.FetchAccount(createAccountResponse.Data.ID)
	require.Error(t, err)
}

func TestAccountDelete(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	createAccountRequest := NewRequestObject()
	createAccountResponse, err := client.CreateAccount(&createAccountRequest)
	require.Nil(t, err)

	err = client.DeleteAccount(createAccountResponse.Data.ID, 0)
	require.Nil(t, err)

	_, err = client.FetchAccount(createAccountResponse.Data.ID)
	require.Error(t, err)
}

func TestFailedAccountDelete(t *testing.T) {
	client, err := f3.NewClient()
	require.Nil(t, err)

	err = client.DeleteAccount(uuid.New().String(), 0)
	require.Error(t, err)
}
