package aikido

import "fmt"

type DeleteTeamRequest struct {
	ID int32
}

func (c *Client) DeleteTeam(request DeleteTeamRequest) (bool, error) {

	err := makeBearerRequest(
		c,
		"DELETE",
		fmt.Sprintf("api/public/v1/teams/%d", request.ID),
		nil,
		204,
		[]int{400, 404},
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
