package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestPaginationHooks_NextPage(t *testing.T) {
	nextURL := "https://example.com/api/next?page=2"

	hooks := models.PaginationHooks{
		NextURL: nextURL,
	}

	assert.Equal(t, nextURL, hooks.NextPage())
}

func TestErrorResponse_Error(t *testing.T) {
	statusCode := 400
	status := "ERROR"
	requestID := "test_request_id"
	errorMessage := "An error occurred"

	errorResp := &models.ErrorResponse{
		BaseResponse: models.BaseResponse{
			Status:       status,
			RequestID:    requestID,
			ErrorMessage: errorMessage,
		},
		StatusCode: statusCode,
	}

	expectedError := "bad status with code '400': message 'An error occurred': request ID 'test_request_id': internal status: 'ERROR'"
	assert.Equal(t, expectedError, errorResp.Error())
}
