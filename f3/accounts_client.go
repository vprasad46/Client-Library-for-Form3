package f3

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	logger "github.com/sirupsen/logrus"

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
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
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
	if resp.StatusCode > http.StatusCreated {
		logger.Debug("Account not created. Got Response code: %v", resp.StatusCode)
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
