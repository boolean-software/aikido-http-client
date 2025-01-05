package aikido

import (
	"encoding/base64"
	"fmt"
)

type Auth interface {
	HeaderValue() string
}

type BasicAuth struct {
	clientId     string
	clientSecret string
}

func (b BasicAuth) HeaderValue() string {
	concatenatedCreds := fmt.Sprintf("%s:%s", b.clientId, b.clientSecret)
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(concatenatedCreds))
	return fmt.Sprintf("Basic %s", encodedCreds)
}

type BearerAuth struct {
	token string
}

func (b BearerAuth) HeaderValue() string {
	return fmt.Sprintf("Bearer %s", b.token)
}
