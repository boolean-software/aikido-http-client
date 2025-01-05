package aikidohttp

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CloudEnvironment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Environment string `json:"environment"`
	ExternalID  string `json:"external_id"`
}

type ListCloudsFilters struct {
	Page    int32
	PerPage int32
}

var DefaultListCloudsFilters = ListCloudsFilters{
	Page:    0,
	PerPage: 20,
}

func (c *AikidoHttpClient) ListClouds(filters ListCloudsFilters) ([]CloudEnvironment, error) {
	params := url.Values{}

	params.Set("page", strconv.FormatInt(int64(filters.Page), 10))

	params.Set("per_page", strconv.FormatInt(int64(filters.PerPage), 10))

	req, err := c.makeRequest("GET", "api/public/v1/clouds?"+params.Encode(), nil)
	if err != nil {
		return []CloudEnvironment{}, err
	}

	token, err := c.getToken()
	if err != nil {
		return []CloudEnvironment{}, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 200, []int{})
	if err != nil {
		return []CloudEnvironment{}, err
	}

	var clouds []CloudEnvironment

	err = json.Unmarshal(responseBody, &clouds)
	if err != nil {
		return []CloudEnvironment{}, err
	}

	return clouds, nil
}
