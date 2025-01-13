package aikido

import "github.com/google/go-querystring/query"

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
	Page    int32 `url:"page"`
	PerPage int32 `url:"per_page"`
}

var DefaultListTeamsFilters = ListTeamsFilters{
	Page:    0,
	PerPage: 20,
}

func (c *Client) ListTeams(filters ListTeamsFilters) ([]Team, error) {
	params, err := query.Values(filters)
	if err != nil {
		return []Team{}, err
	}

	return makeBearerRequestAndDecode[[]Team](
		c,
		"GET",
		"api/public/v1/teams?"+params.Encode(),
		nil,
		200,
		[]int{},
	)
}
