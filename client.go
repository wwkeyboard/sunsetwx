package sunsetwx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Client for requesting data from the API
type Client struct {
	username    string
	password    string
	accessToken string
}

// AccessTokenResponse we get back from the API
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

func (c *Client) get(path string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://sunburst.sunsetwx.com/v1/quality", nil)
	if err != nil {
		fmt.Println("could make request", err)
		os.Exit(1)
	}
	token := "foo"
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error querying api", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", body)
	return nil, nil
}
