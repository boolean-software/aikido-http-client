package aikido

type Repository struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Provider       string `json:"provider"`
	ExternalRepoID string `json:"external_repo_id"`
}

func (c *Client) RepositoryDetail(id string) (Repository, error) {
	return makeBearerRequestAndDecode[Repository](
		c,
		"GET",
		"api/public/v1/repositories/code/"+id,
		nil,
		200,
		[]int{},
	)
}
