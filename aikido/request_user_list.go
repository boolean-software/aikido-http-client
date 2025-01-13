package aikido

import "github.com/google/go-querystring/query"

type User struct {
	ID                 int    `json:"id"`
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	Active             int    `json:"active"`
	LastLoginTimestamp int    `json:"last_login_timestamp"`
	Role               string `json:"role"`
	AuthType           string `json:"auth_type"`
}

type ListUsersFilters struct {
	TeamId          int32 `url:"filter_team_id,omitempty"`
	IncludeInactive int32 `url:"include_inactive"`
}

var DefaultListUsersFilters = ListUsersFilters{
	IncludeInactive: 0,
}

func (c *Client) ListUsers(filters ListUsersFilters) ([]User, error) {
	params, err := query.Values(filters)
	if err != nil {
		return []User{}, err
	}

	return makeBearerRequestAndDecode[[]User](
		c,
		"GET",
		"api/public/v1/users?"+params.Encode(),
		nil,
		200,
		[]int{},
	)
}
