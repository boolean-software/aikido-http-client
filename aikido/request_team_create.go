package aikido

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type createTeamResponse struct {
	ID string `json:"id"`
}

func (c *Client) CreateTeam(request CreateTeamRequest) (string, error) {

	res, err := makeBearerRequestAndDecode[createTeamResponse](
		c,
		"POST",
		"api/public/v1/teams",
		request,
		201,
		[]int{400},
	)

	if err != nil {
		return "", err
	}

	return res.ID, nil
}
