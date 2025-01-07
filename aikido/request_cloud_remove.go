package aikido

import (
	"encoding/json"
	"fmt"
)

type removeCloudResponse struct {
	Success bool `json:"success"`
}

func (c *Client) RemoveCloud(id string) (bool, error) {
	path := fmt.Sprintf("api/public/v1/clouds/%s", id)

	req, err := c.makeRequest("DELETE", path, nil)
	if err != nil {
		return false, err
	}

	token, err := c.getToken()
	if err != nil {
		return false, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, 200, []int{400, 404})
	if err != nil {
		return false, err
	}

	var removalSuccess removeCloudResponse

	err = json.Unmarshal(responseBody, &removalSuccess)
	if err != nil {
		return false, err
	}

	return removalSuccess.Success, nil
}
