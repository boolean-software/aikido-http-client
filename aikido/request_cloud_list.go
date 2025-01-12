package aikido

import (
	"github.com/boolean-software/aikido-http-client/internal/util"
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
	params, err := util.BuildURLParams(filters)
	if err != nil {
		return []CloudEnvironment{}, err
	}

	return makeBearerRequestAndDecode[[]CloudEnvironment](
		c,
		"GET",
		"api/public/v1/clouds?"+params,
		nil,
		200,
		[]int{},
	)
}
