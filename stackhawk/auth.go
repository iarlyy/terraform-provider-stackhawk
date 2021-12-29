package stackhawk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Client) LogIn() (*LoginResponse, error) {
	if c.ApiKey == "" {
		return nil, fmt.Errorf("apiKey is not defined")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/auth/login", c.ApiURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-ApiKey", c.ApiKey)
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := LoginResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
