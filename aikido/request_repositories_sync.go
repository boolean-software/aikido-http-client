package aikido

func (c *Client) RepositoriesSync() error {
	return makeBearerRequest(
		c,
		"POST",
		"api/public/v1/repositories/import",
		nil,
		200,
		[]int{},
	)
}
