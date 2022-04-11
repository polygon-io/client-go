package models

import "fmt"

// BaseResponse has all possible attributes that any response can use. It's intended to be embedded in a
// domain specific response struct.
type BaseResponse struct {
	Status       string `json:"status,omitempty"`
	RequestID    string `json:"request_id,omitempty"`
	Count        int    `json:"count,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error,omitempty"`

	PaginationHooks
}

// PaginationHooks are links to next and/or previous pages. Embed this struct into an API response if
// the endpoint supports pagination.
type PaginationHooks struct {
	NextURL string `json:"next_url,omitempty"`
}

func (p PaginationHooks) NextPage() string {
	return p.NextURL
}

// ErrorResponse represents an API response with an error status code.
type ErrorResponse struct {
	StatusCode int
	BaseResponse
}

// Error returns the details of an error response.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code '%d': message '%s': request ID '%s': internal status: '%s'", e.StatusCode, e.ErrorMessage, e.RequestID, e.Status)
}
