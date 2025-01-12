package aikido

type UpdateTeamRequest struct {
	ID               string           `json:"id"`
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
		"api/public/v1/teams"+request.ID,
		nil,
		200,
		[]int{400, 404},
	)

	if err != nil {
		return false, err
	}

	return res.Status == "ok", nil
}
