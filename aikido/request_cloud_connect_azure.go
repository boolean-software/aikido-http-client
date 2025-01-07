package aikido

import (
	"encoding/json"
)

type connectAzureRequest struct {
	Name           string `json:"name"`
	ApplicationID  string `json:"application_id"`
	DirectoryID    string `json:"directory_id"`
	KeyValue       string `json:"key_value"`
	SubscriptionID string `json:"subscription_id"`
	Environment    string
}

type connectAzureResponse struct {
	ID string `json:"id"`
}

func (c *Client) ConnectAzureCloud(name string, applicationID string, directoryID string, keyValue string, subscriptionID string, environment string) (string, error) {

	req, err := c.makeRequest("POST", "api/public/v1/clouds/azure", connectAzureRequest{
		Name:           name,
		ApplicationID:  applicationID,
		DirectoryID:    directoryID,
		KeyValue:       keyValue,
		SubscriptionID: subscriptionID,
		Environment:    environment,
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

	var cloud connectAzureResponse

	err = json.Unmarshal(responseBody, &cloud)
	if err != nil {
		return "", err
	}

	return cloud.ID, nil
}
