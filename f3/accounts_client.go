package f3

import (
	"bytes"
	"encoding/json"
	"errors"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

func (apiClient *Client) CreateAccount(createAccountRequest *CreateAccountRequest) (*CreateAccountResponse, error) {
	logger.Debug("Create Account Function called...")
	url := apiClient.baseURL + "/v1/organisation/accounts"
	data, err := json.Marshal(createAccountRequest)
	if err != nil {
		logger.Error("Error occured while marshaling CreateAccountRequest object as json")
		return nil, err
	}
	req, err := NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		logger.Error("Error while creating request object for create account")
		return nil, err
	}
	logger.Debug("Request Object for create account created")

	req.Header.Set("Accept", "vnd.api+json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := apiClient.httpClient.Do(req)
	logger.Debug("Response received from create account API")
	if err != nil {
		logger.Error("Error sending request to create account API")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Error while accessing Resource in the given baseURL. Check F3_BASE_URL's value")
	}
	if resp.StatusCode != http.StatusCreated {
		logger.Debug("Account not created. Got Response code: ", resp.StatusCode)
		errBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Error reading error response from create account API")
			return nil, err
		}
		var accountCreationError ErrorMessage
		if err = json.Unmarshal(errBody, &accountCreationError); err != nil {
			logger.Error("Error unmarshling error response from create account API to ErrorMessage")
			return nil, err
		}
		logger.Debug("Sending back ErrorMessage object as error")
		return nil, accountCreationError
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading success response from create account API")
		return nil, err
	}
	var createAccountResponse *CreateAccountResponse
	if err = json.Unmarshal(body, &createAccountResponse); err != nil {
		logger.Error("Error unmarshaling success response from create account API to CreateAccountResponse")
		return nil, err
	}
	return createAccountResponse, nil
}

func (apiClient *Client) FetchAccount(accountID string) (*FetchAccountResponse, error) {
	logger.Debug("Fetch Account Function called...")
	url := apiClient.baseURL + "/v1/organisation/accounts/" + accountID

	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Error("Error while creating request object for fetch account")
		return nil, err
	}
	logger.Debug("Request Object for fetch account created")

	req.Header.Set("Accept", "vnd.api+json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := apiClient.httpClient.Do(req)
	logger.Debug("Response received from fetch account API")
	if err != nil {
		logger.Error("Error sending request to fetch account API")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Error while accessing Resource in the given baseURL. Check F3_BASE_URL's value")
	}
	if resp.StatusCode != http.StatusOK {
		logger.Debug("Account not fetched. Got Response code:", resp.StatusCode)
		errBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Error reading error response from fetch account API")
			return nil, err
		}
		var accountFetchError ErrorMessage
		if err = json.Unmarshal(errBody, &accountFetchError); err != nil {
			logger.Error("Error unmarshling error response from fetch account API to ErrorMessage")
			return nil, err
		}
		logger.Debug("Sending back ErrorMessage object as error")
		return nil, accountFetchError
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading success response from fetch account API")
		return nil, err
	}
	var fetchAccountResponse *FetchAccountResponse
	if err = json.Unmarshal(body, &fetchAccountResponse); err != nil {
		logger.Error("Error unmarshaling success response from fetch account API to FetchAccountResponse")
		return nil, err
	}
	return fetchAccountResponse, nil
}

func (apiClient *Client) DeleteAccount(accountID string, version int) error {
	logger.Debug("Delete Account Function called...")
	url := apiClient.baseURL + "/v1/organisation/accounts/" + accountID

	req, err := NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logger.Error("Error while creating request object for delete account")
		return err
	}
	logger.Debug("Request Object for delete account created")

	q := req.URL.Query()
	q.Add("version", strconv.Itoa(version))

	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accept", "vnd.api+json")
	req.Header.Set("Content-Type", "application/vnd.api+json")

	resp, err := apiClient.httpClient.Do(req)
	logger.Debug("Response received from delete account API")
	if err != nil {
		logger.Error("Error sending request to delete account API")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return errors.New("Error while accessing Resource in the given baseURL. Check F3_BASE_URL's value")
	}
	if resp.StatusCode != http.StatusNoContent {
		logger.Debug("Account not deleted. Got Response code:", resp.StatusCode)
		errBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Error reading error response from fetch account API")
			return err
		}
		var accountDeleteError ErrorMessage
		if err = json.Unmarshal(errBody, &accountDeleteError); err != nil {
			logger.Error("Error unmarshling error response from delete account API to ErrorMessage")
			return err
		}
		logger.Debug("Sending back ErrorMessage object as error")
		return accountDeleteError
	}
	return nil
}
