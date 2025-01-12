package aikido

type ConnectAzureRequest struct {
	Name           string `json:"name"`
	ApplicationID  string `json:"application_id"`
	DirectoryID    string `json:"directory_id"`
	KeyValue       string `json:"key_value"`
	SubscriptionID string `json:"subscription_id"`
	Environment    string `json:"environment"`
}

type connectAzureResponse struct {
	ID string `json:"id"`
}

func (c *Client) ConnectAzureCloud(request ConnectAzureRequest) (string, error) {
	res, err := makeBearerRequestAndDecode[connectAzureResponse](
		c,
		"POST",
		"api/public/v1/clouds/azure",
		request,
		201,
		[]int{422},
	)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}
