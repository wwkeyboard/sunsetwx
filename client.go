package sunsetwx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// BaseURL of the API we target
const BaseURL = "https://sunburst.sunsetwx.com/v1"

// Client for requesting data from the API
type Client struct {
	accessToken string
}

// AccessTokenResponse we get back from the API
type AccessTokenResponse struct {
	Message     string `json:"message"`
	Notice      string `json:"notice"`
	AccessToken string `json:"token"`
	ExpiresIn   int    `json:"token_exp_sec"`
	Scope       string `json:"scope"`
}

// NewClient for the API
func NewClient() Client {
	return Client{}
}

type amendRequest func(*http.Request) error

func (c *Client) get(path string, reqSetup amendRequest) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", BaseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("building GET request %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.accessToken))

	// Let the caller modify the quest before it's sent, if they want.
	if reqSetup != nil {
		err = reqSetup(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) post(path string, data string, reqSetup amendRequest) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", BaseURL+path, strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("building POST request %v", err)
	}

	// Add the authorization to every request, for logins we can strip it out since the Bearer will be empty
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.accessToken))

	// reqSetup is intended for modifying the request before it's sent, e.g. setting headers
	if reqSetup != nil {
		err = reqSetup(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// if it's a non 200 response there is probably part of the body we're interested in, so we extract and return it.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// This may not be an error for go, but it's an error for us
	if resp.StatusCode > 300 {
		return body, fmt.Errorf("POST to %v returned %v", path, resp.StatusCode)
	}

	return body, nil
}

// Login to the SunsetWX API
func (c *Client) Login(username, password string) error {
	data := url.Values{}

	data.Set("grant_type", "password")
	data.Set("type", "remember_me")
	body, err := c.post("/login", data.Encode(), func(req *http.Request) error {
		req.Header.Del("Authorization")
		req.SetBasicAuth(username, password)
		return nil
	})
	if err != nil {
		return nil
	}
	return c.setAuthToken(body)
}

// Login to the sunsetwx API
func (c *Client) setAuthToken(data []byte) error {
	var atr AccessTokenResponse
	err := json.Unmarshal(data, &atr)
	if err != nil {
		return err
	}

	c.accessToken = atr.AccessToken
	return nil
}

// GetQuality prediction from the API
func (c *Client) GetQuality(lat, lon float64) (*FeatureCollection, error) {
	data := url.Values{}
	data.Set("geo", fmt.Sprintf("%f,%f", lat, lon))

	query := fmt.Sprintf("/quality?geo=%f,%f", lat, lon)

	body, err := c.get(query, nil)
	if err != nil {
		return nil, fmt.Errorf("getting quality: %v", string(body))
	}
	return FromJSON(body)
}
