package aikido

import (
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

	return makeBearerRequestAndDecode[[]Repository](
		c,
		"GET",
		"api/public/v1/repositories/code?"+params.Encode(),
		nil,
		200,
		[]int{},
	)
}
