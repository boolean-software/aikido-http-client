package aikido

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type ListRepositoriesFilters struct {
	Page            int32
	PerPage         int32
	IncludeInactive bool
}

var DefaultListRepositoriesFilters = ListRepositoriesFilters{
	Page:            0,
	PerPage:         20,
	IncludeInactive: false,
}

func (c *Client) ListRepositories(filters ListRepositoriesFilters) ([]Repository, error) {
	params := url.Values{}

	params.Set("page", strconv.FormatInt(int64(filters.Page), 10))

	params.Set("per_page", strconv.FormatInt(int64(filters.PerPage), 10))

	params.Set("include_inactive", fmt.Sprintf("%t", filters.IncludeInactive))

	req, err := c.makeRequest("GET", "api/public/v1/repositories/code?"+params.Encode(), nil)
	if err != nil {
		return []Repository{}, err
	}

	token, err := c.getToken()
	if err != nil {
		return []Repository{}, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 200, []int{})
	if err != nil {
		return []Repository{}, err
	}

	var repos []Repository

	err = json.Unmarshal(responseBody, &repos)
	if err != nil {
		return []Repository{}, err
	}

	return repos, nil
}
