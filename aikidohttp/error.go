package aikidohttp

type ErrorBody struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Example          string `json:"example,omitempty"`
}
