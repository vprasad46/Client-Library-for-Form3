package f3

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
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

func (c *RequestClient) drainBody(body io.ReadCloser) {
	defer body.Close()
	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, 10))
	if err != nil {
		logger.Error("Error draining requestBody", err)
	}
}

func (c *RequestClient) Do(req *Request) (*http.Response, error) {
	logger.Debug(req.Method, req.URL)
	client, _ := NewClient()
	i := 1
	for {
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				logger.Error("Error occured while setting the seek back to 0 in request body")
				return nil, errors.New("request body seek failed")
			}
		}
		resp, err := c.HTTPClient.Do(req.Request)
		retry, err := IsRetryNeeded(resp, err)
		if !retry {
			return resp, err
		}
		c.drainBody(resp.Body)
		remain := c.MaxTries - i
		if remain <= 0 {
			break
		}
		wait := TimeToSleep(c.WaitMin, c.WaitMax, i)
		logger.Debug("Retry in ", wait, "seconds")
		time.Sleep(wait)
		i = i + 1
	}
	return nil, errors.New("request  not sent error after maximux retries")
}

func IsRetryNeeded(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}
	if resp.StatusCode >= 0 {
		return true, nil
	}
	return false, nil
}

func TimeToSleep(min, max time.Duration, attemptNum int) time.Duration {
	timeToSleep := math.Pow(1.414, float64(attemptNum)) * float64(min)
	floatTimeToSleep := time.Duration(timeToSleep)

	if floatTimeToSleep > max {
		floatTimeToSleep = max
	}
	return floatTimeToSleep
}

type RequestClient struct {
	HTTPClient *http.Client
	WaitMin    time.Duration
	WaitMax    time.Duration
	MaxTries   int
}

func NewRequestClient() *RequestClient {
	return &RequestClient{
		HTTPClient: &http.Client{},
		WaitMin:    1 * time.Second,
		WaitMax:    5 * time.Second,
		MaxTries:   3,
	}
}

type Request struct {
	body io.ReadSeeker
	*http.Request
}

func NewRequest(method, url string, body io.ReadSeeker) (*Request, error) {
	var reqBody io.ReadCloser
	if body != nil {
		reqBody = ioutil.NopCloser(body)
	}

	httpReq, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	return &Request{body, httpReq}, nil
}
