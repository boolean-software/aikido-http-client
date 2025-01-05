package aikido

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type AikidoUser struct {
	ID                 int    `json:"id"`
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	Active             int    `json:"active"`
	LastLoginTimestamp int    `json:"last_login_timestamp"`
	Role               string `json:"role"`
	AuthType           string `json:"auth_type"`
}

type ListUsersFilters struct {
	TeamId          int32
	IncludeInactive int32
}

var DefaultListUsersFilters = ListUsersFilters{
	TeamId:          -1,
	IncludeInactive: -1,
}

func (c *Client) ListUsers(filters ListUsersFilters) ([]AikidoUser, error) {
	params := url.Values{}

	if filters.TeamId >= 1 {
		params.Set("filter_team_id", strconv.FormatInt(int64(filters.TeamId), 10))
	}

	if filters.IncludeInactive >= 0 {
		params.Set("include_inactive", strconv.FormatInt(int64(filters.IncludeInactive), 10))
	}

	req, err := c.makeRequest("GET", "api/public/v1/users?"+params.Encode(), nil)
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
