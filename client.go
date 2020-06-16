package sunsetwx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
		fmt.Println("could make request", err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.accessToken))

	// Let's the called fixup the request before we send it.
	err = reqSetup(req)
	if err != nil {
		return nil, err
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

func (c *Client) post(path string, data url.Values, reqSetup amendRequest) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", BaseURL+path, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("could make request", err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.accessToken))

	// Let's the called fixup the request before we send it.
	err = reqSetup(req)
	if err != nil {
		return nil, err
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

// Login to the SunsetWX API
func (c *Client) Login(username, password string) error {
	data := url.Values{}

	data.Set("grant_type", "password")
	data.Set("type", "remember_me")

	body, err := c.post("/login", data, func(req *http.Request) error {
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
	body, err := c.get("/quality", func(req *http.Request) error {
		location := fmt.Sprintf("%f,%f", lat, lon)
		req.URL.Query().Add("geo", location)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return FromJSON(body)
}
