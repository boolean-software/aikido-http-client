package aikido

import "fmt"

type errorBody struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Example          string `json:"example,omitempty"`
	ReasonPhrase     string `json:"reason_phrase,omitempty"`
}

func (err errorBody) error() error {
	if err.Error != "" {
		return fmt.Errorf("%s (%s) \n%s", err.ErrorDescription, err.Error, err.Example)
	}

	if err.ReasonPhrase != "" {
		return fmt.Errorf("%s", err.ReasonPhrase)
	}

	return fmt.Errorf("failed")
}
