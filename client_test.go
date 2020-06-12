package sunsetwx

import (
	"testing"
)

func TestLogin(t *testing.T) {
	client := NewClient("testuser", "testpw")

	err := client.setAuthToken([]byte(sample()))
	if err != nil {
		t.Error(err)
		return
	}

	if client.accessToken != "6340beaa-7279-4055-a6ec-fc9422ee0ad8" {
		t.Errorf("didn't parse access token from respones correctly")
	}
}

func sample() string {
	return `{
		"message": "Login successful",
		"notice": "If this token must be used directly by web browsers, which is not recommended, we suggest setting one or more space-delimited origins, using the origins parameter, to discourage token theft.",
		"access_token": "6340beaa-7279-4055-a6ec-fc9422ee0ad8",
		"token_type": "Bearer",
		"expires_in": 604800,
		"scope": "predictions"
	  }`
}
