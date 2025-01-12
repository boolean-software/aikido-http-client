package aikido

type ConnectAWSRequest struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	RoleArn     string `json:"role_arn"`
}

type connectAwsResponse struct {
	ID string `json:"id"`
}

func (c *Client) ConnectAWSCloud(request ConnectAWSRequest) (string, error) {
	res, err := makeBearerRequestAndDecode[connectAwsResponse](
		c,
		"POST",
		"api/public/v1/clouds/aws",
		request,
		201,
		[]int{422},
	)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}
