package sunsetwx

import (
	"encoding/json"
)

// Client for requesting data from the API
type Client struct {
	username    string
	password    string
	accessToken string
}

type AccessTokenResponse struct {
	Message     string `json:"message"`
	Notice      string `json:"notice"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// NewClient for the API
func NewClient(username, password string) Client {
	return Client{
		username: username,
		password: password,
	}
}

func (c *Client) setAuthToken(data []byte) error {
	var atr AccessTokenResponse
	err := json.Unmarshal(data, &atr)
	if err != nil {
		return err
	}

	c.accessToken = atr.AccessToken
	return nil
}
