package aikido

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type createTeamResponse struct {
	ID int32 `json:"id"`
}

func (c *Client) CreateTeam(request CreateTeamRequest) (int32, error) {

	res, err := makeBearerRequestAndDecode[createTeamResponse](
		c,
		"POST",
		"api/public/v1/teams",
		request,
		201,
		[]int{400},
	)

	if err != nil {
		return -1, err
	}

	return res.ID, nil
}
