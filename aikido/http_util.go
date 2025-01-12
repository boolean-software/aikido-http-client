package aikido

import "encoding/json"

func makeBearerRequestAndDecode[T any](c *Client, method string, path string, body any, expectedStatusCode int, failureStatusCodes []int) (T, error) {
	var result T

	req, err := c.makeRequest(method, path, body)
	if err != nil {
		return result, err
	}

	token, err := c.getToken()
	if err != nil {
		return result, err
	}

	responseBody, err := c.do(req, BearerAuth{token}, expectedStatusCode, failureStatusCodes)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func makeBearerRequest(c *Client, method string, path string, body any, expectedStatusCode int, failureStatusCodes []int) error {
	req, err := c.makeRequest(method, path, body)
	if err != nil {
		return err
	}

	token, err := c.getToken()
	if err != nil {
		return err
	}

	_, err = c.do(req, BearerAuth{token}, expectedStatusCode, failureStatusCodes)
	if err != nil {
		return err
	}

	return nil
}
