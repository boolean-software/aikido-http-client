package aikidohttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"
)

type AikidoHttpClient struct {
	client       *http.Client
	url          string
	clientId     string
	clientSecret string
	accessToken  string
	refreshTime  time.Time
}

func NewAikidoHttpClient(clientId string, clientSecret string) *AikidoHttpClient {
	return &AikidoHttpClient{
		client: &http.Client{
			Timeout: 1 * time.Minute,
		},
		clientId:     clientId,
		clientSecret: clientSecret,
		url:          "https://app.aikido.dev",
	}
}

func (c *AikidoHttpClient) urlFor(path string) string {
	return fmt.Sprintf("%s/%s", c.url, path)
}

func (c *AikidoHttpClient) getToken() (string, error) {
	if time.Now().After(c.refreshTime) {
		err := c.Auth(c.clientId, c.clientSecret)
		if err != nil {
			return "", err
		}
	}

	return c.accessToken, nil
}

func (c *AikidoHttpClient) makeRequest(method string, path string, body any) (*http.Request, error) {
	url := c.urlFor(path)

	var bodyReader io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return &http.Request{}, err
		}

		bodyReader = strings.NewReader(string(jsonData))
	}

	return http.NewRequest(method, url, bodyReader)
}

func (c *AikidoHttpClient) do(req *http.Request, auth Auth, expectedStatusCode int, failureStatusCodes []int) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", auth.HeaderValue())

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if slices.Contains(failureStatusCodes, resp.StatusCode) {
		var error *ErrorBody

		err = json.Unmarshal(responseBody, &error)
		if err != nil {
			return []byte{}, err
		}

		return []byte{}, error.error()
	}

	if resp.StatusCode != expectedStatusCode {
		return []byte{}, fmt.Errorf("unexpected status code `%d` instead of ``%d`", resp.StatusCode, 200)
	}

	return responseBody, nil
}
