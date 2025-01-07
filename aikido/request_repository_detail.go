package aikido

import "encoding/json"

type Repository struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Provider       string `json:"provider"`
	ExternalRepoID string `json:"external_repo_id"`
}

func (c *Client) RepositoryDetail(id string) (Repository, error) {
	req, err := c.makeRequest("GET", "api/public/v1/repositories/code/"+id, nil)
	if err != nil {
		return Repository{}, err
	}

	token, err := c.getToken()
	if err != nil {
		return Repository{}, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 200, []int{})
	if err != nil {
		return Repository{}, err
	}

	var repo Repository

	err = json.Unmarshal(responseBody, &repo)
	if err != nil {
		return Repository{}, err
	}

	return repo, nil
}
