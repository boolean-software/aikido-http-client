package aikido

type removeCloudResponse struct {
	Success bool `json:"success"`
}

func (c *Client) RemoveCloud(id string) (bool, error) {
	res, err := makeBearerRequestAndDecode[removeCloudResponse](
		c,
		"DELETE",
		"api/public/v1/clouds/"+id,
		nil,
		200,
		[]int{400, 404},
	)

	if err != nil {
		return false, err
	}

	return res.Success, nil
}
