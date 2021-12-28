package stackhawk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateApplication(application Application) (*Application, error) {
	rb, err := json.Marshal(application)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/org/%s/app", c.ApiURL, c.OrgId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("%s,%s", rb, err)
	}

	newApplication := Application{}
	err = json.Unmarshal(body, &newApplication)
	if err != nil {
		return nil, err
	}

	return &newApplication, nil
}

func (c *Client) GetApplication(applicationId string) (*Application, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/app/%s", c.ApiURL, applicationId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	application := Application{}
	err = json.Unmarshal(body, &application)
	if err != nil {
		return nil, err
	}

	return &application, nil
}

func (c *Client) UpdateApplication(application Application) (*Application, error) {
	rb, err := json.Marshal(application)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/app/%s", c.ApiURL, application.ApplicationId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedApplication := Application{}
	err = json.Unmarshal(body, &updatedApplication)
	if err != nil {
		return nil, err
	}

	return &updatedApplication, nil
}

func (c *Client) DeleteApplication(applicationId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/app/%s", c.ApiURL, applicationId), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
