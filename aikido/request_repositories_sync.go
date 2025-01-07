package aikido

func (c *Client) RepositoriesSync() error {
	req, err := c.makeRequest("POST", "api/public/v1/repositories/import", nil)
	if err != nil {
		return err
	}

	token, err := c.getToken()
	if err != nil {
		return err
	}

	_, err = c.do(req, BearerAuth{token}, 200, []int{})
	if err != nil {
		return err
	}

	return nil
}
