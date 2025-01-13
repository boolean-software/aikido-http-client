package aikido

import (
	"github.com/google/go-querystring/query"
)

type CloudEnvironment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Environment string `json:"environment"`
	ExternalID  string `json:"external_id"`
}

type ListCloudsFilters struct {
	Page    int32 `url:"page"`
	PerPage int32 `url:"per_page"`
}

var DefaultListCloudsFilters = ListCloudsFilters{
	Page:    0,
	PerPage: 20,
}

func (c *Client) ListClouds(filters ListCloudsFilters) ([]CloudEnvironment, error) {
	params, err := query.Values(filters)
	if err != nil {
		return []CloudEnvironment{}, err
	}

	return makeBearerRequestAndDecode[[]CloudEnvironment](
		c,
		"GET",
		"api/public/v1/clouds?"+params.Encode(),
		nil,
		200,
		[]int{},
	)
}
