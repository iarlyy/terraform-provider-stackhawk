package stackhawk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const ApiURL string = "https://api.stackhawk.com/api/v1"

// Client -
type Client struct {
	HTTPClient *http.Client
	UserAgent  string
	ApiURL     string
	ApiKey     string
	Token      string
	OrgId      string
}

func NewClient(apiUrl, orgId, apiKey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		ApiURL:     ApiURL,
		ApiKey:     *apiKey,
		OrgId:      *orgId,
	}

	if apiUrl != nil {
		c.ApiURL = *apiUrl
	}

	ar, err := c.LogIn()
	if err != nil {
		return nil, fmt.Errorf("Login error: %s", err)
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	if c.UserAgent == "" {
		req.Header.Set("User-Agent", "stackhawk-go-client")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request failed, bad status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
