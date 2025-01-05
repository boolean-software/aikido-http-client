package aikido

import "fmt"

type ErrorBody struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Example          string `json:"example,omitempty"`
	ReasonPhrase     string `json:"reason_phrase,omitempty"`
}

func (err ErrorBody) error() error {
	if err.Error != "" {
		return fmt.Errorf("%s (%s) \n%s", err.ErrorDescription, err.Error, err.Example)
	}

	if err.ReasonPhrase != "" {
		return fmt.Errorf("%s", err.ReasonPhrase)
	}

	return fmt.Errorf("failed")
}
