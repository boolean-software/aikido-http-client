package aikidohttp

import "encoding/json"

type AikidoUser struct {
	ID                 int    `json:"id"`
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	Active             int    `json:"active"`
	LastLoginTimestamp int    `json:"last_login_timestamp"`
	Role               string `json:"role"`
	AuthType           string `json:"auth_type"`
}

func (c *AikidoHttpClient) ListUsers() ([]AikidoUser, error) {
	req, err := c.makeRequest("GET", "api/public/v1/users", nil)
	if err != nil {
		return []AikidoUser{}, err
	}

	token, err := c.getToken()
	if err != nil {
		return []AikidoUser{}, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 200, []int{})
	if err != nil {
		return []AikidoUser{}, err
	}

	var users []AikidoUser

	err = json.Unmarshal(responseBody, &users)
	if err != nil {
		return []AikidoUser{}, err
	}

	return users, nil
}
