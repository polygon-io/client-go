package models

import "fmt"

// BaseResponse has all possible attributes that any response can use. It's intended to be embedded in a domain specific
// response struct.
type BaseResponse struct {
	PaginationHooks

	// The status of this request's response.
	Status string `json:"status,omitempty"`

	// A request id assigned by the server.
	RequestID string `json:"request_id,omitempty"`

	// The total number of results for this request.
	Count int `json:"count,omitempty"`

	// A response message for successful requests.
	Message string `json:"message,omitempty"`

	// An error message for unsuccessful requests.
	ErrorMessage string `json:"error,omitempty"`
}

// PaginationHooks are links to next and/or previous pages. Embed this struct into an API response if the endpoint
// supports pagination.
type PaginationHooks struct {
	// If present, this value can be used to fetch the next page of data.
	NextURL string `json:"next_url,omitempty"`
}

func (p PaginationHooks) NextPage() string {
	return p.NextURL
}

// ErrorResponse represents an API response with an error status code.
type ErrorResponse struct {
	BaseResponse

	// An HTTP status code for unsuccessful requests.
	StatusCode int
}

// Error returns the details of an error response.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code '%d': message '%s': request ID '%s': internal status: '%s'", e.StatusCode, e.ErrorMessage, e.RequestID, e.Status)
}
