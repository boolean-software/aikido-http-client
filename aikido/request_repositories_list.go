package aikido

import (
	"github.com/boolean-software/aikido-http-client/internal/util"
)

type ListRepositoriesFilters struct {
	Page            int32 `url:"page"`
	PerPage         int32 `url:"per_page"`
	IncludeInactive bool  `url:"include_inactive"`
}

var DefaultListRepositoriesFilters = ListRepositoriesFilters{
	Page:            0,
	PerPage:         20,
	IncludeInactive: false,
}

func (c *Client) ListRepositories(filters ListRepositoriesFilters) ([]Repository, error) {
	params, err := util.BuildURLParams(filters)
	if err != nil {
		return []Repository{}, err
	}

	return makeBearerRequestAndDecode[[]Repository](
		c,
		"GET",
		"api/public/v1/repositories/code?"+params,
		nil,
		200,
		[]int{},
	)
}
