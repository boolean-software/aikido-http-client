package aikido

type DeleteTeamRequest struct {
	ID string
}

func (c *Client) DeleteTeam(request DeleteTeamRequest) (bool, error) {

	err := makeBearerRequest(
		c,
		"DELETE",
		"api/public/v1/teams/"+request.ID,
		nil,
		204,
		[]int{400, 404},
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
