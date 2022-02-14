package f3

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

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

func NewClient() (*Client, error) {
	baseURL := os.Getenv("F3_BASE_URL")
	if baseURL == "" {
		logger.Error("Base URL not set in environment for using F3 Accounts Client")
		return nil, errors.New("base url not set in environment")
	}
	client := Client{baseURL: baseURL, httpClient: &http.Client{}}
	logger.Debug("F3 Accounts Client Creation Successful")
	return &client, nil
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
