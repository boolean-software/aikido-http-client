package aikidohttp

import (
	"encoding/json"
	"time"
)

type AuthBody struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (c *AikidoHttpClient) Auth(clientId string, clientSecret string) error {
	req, err := c.makeRequest("POST", "api/oauth/token", AuthBody{
		GrantType: "client_credentials",
	})
	if err != nil {
		return err
	}

	responseBody, err := c.do(req, BasicAuth{clientId, clientSecret}, 200, []int{400})
	if err != nil {
		return err
	}

	var auth *AuthResponse

	err = json.Unmarshal(responseBody, &auth)
	if err != nil {
		return err
	}

	now := time.Now()
	c.accessToken = auth.AccessToken
	c.refreshTime = now.Add(time.Second * time.Duration(auth.ExpiresIn))

	return nil
}
