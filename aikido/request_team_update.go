package aikido

import "fmt"

type UpdateTeamRequest struct {
	ID               int32            `json:"id"`
	Name             string           `json:"name"`
	Responsibilities []Responsibility `json:"responsabilities"`
}

type updateTeamResponse struct {
	Status string `json:"status"`
}

func (c *Client) UpdateTeam(request UpdateTeamRequest) (bool, error) {

	res, err := makeBearerRequestAndDecode[updateTeamResponse](
		c,
		"PUT",
		fmt.Sprintf("api/public/v1/teams/%d", request.ID),
		request,
		200,
		[]int{400, 404},
	)

	if err != nil {
		return false, err
	}

	return res.Status == "ok", nil
}
