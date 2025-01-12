package aikido

import (
	"net/url"
	"strconv"
)

type Team struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	ExternalSource   string           `json:"external_source"`
	ExternalSourceID string           `json:"external_source_id"`
	Responsibilities []Responsibility `json:"responsibilities"`
	Active           bool             `json:"active"`
}

type Responsibility struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type ListTeamsFilters struct {
	Page    int32
	PerPage int32
}

var DefaultListTeamsFilters = ListTeamsFilters{
	Page:    0,
	PerPage: 20,
}

func (c *Client) ListTeams(filters ListTeamsFilters) ([]Team, error) {
	params := url.Values{}

	params.Set("page", strconv.FormatInt(int64(filters.Page), 10))

	params.Set("per_page", strconv.FormatInt(int64(filters.PerPage), 10))

	return makeBearerRequestAndDecode[[]Team](
		c,
		"GET",
		"api/public/v1/teams?"+params.Encode(),
		nil,
		200,
		[]int{},
	)
}
