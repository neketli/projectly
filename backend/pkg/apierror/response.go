package apierror

// Response is the standard API error envelope sent to the client.
type Response struct {
	Error ErrorBody `json:"error"`
}

// ErrorBody contains the error code and human-readable message.
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewResponse builds an error response from a Kind and message.
func NewResponse(kind Kind, message string) *Response {
	return &Response{
		Error: ErrorBody{
			Code:    kind.String(),
			Message: message,
		},
	}
}
