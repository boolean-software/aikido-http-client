package aikido

import (
	"encoding/json"
)

type connectAWSRequest struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	RoleArn     string `json:"role_arn"`
}

type connectAwsResponse struct {
	ID string `json:"id"`
}

func (c *Client) ConnectAWSCloud(roleArn string, name string, environment string) (string, error) {

	req, err := c.makeRequest("POST", "api/public/v1/clouds/aws", connectAWSRequest{
		Name:        name,
		Environment: environment,
		RoleArn:     roleArn,
	})
	if err != nil {
		return "", err
	}

	token, err := c.getToken()
	if err != nil {
		return "", err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 201, []int{422})
	if err != nil {
		return "", err
	}

	var cloud connectAwsResponse

	err = json.Unmarshal(responseBody, &cloud)
	if err != nil {
		return "", err
	}

	return cloud.ID, nil
}
