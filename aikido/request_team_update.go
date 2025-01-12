package aikido

import "encoding/json"

type UpdateTeamRequest struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Responsibilities []Responsibility `json:"responsabilities"`
}

func (c *Client) UpdateTeam(request UpdateTeamRequest) (bool, error) {
	req, err := c.makeRequest("PUT", "api/public/v1/teams"+request.ID, nil)
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
