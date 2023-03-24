package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestRequestOptions(t *testing.T) {
	apiKey := "test_api_key"
	headerKey := "Test-Header"
	headerValue := "Test-Value"
	queryKey := "test_query_key"
	queryValue := "test_query_value"

	opts := &models.RequestOptions{}
	models.APIKey(apiKey)(opts)
	models.Header(headerKey, headerValue)(opts)
	models.QueryParam(queryKey, queryValue)(opts)

	assert.NotNil(t, opts.APIKey)
	assert.Equal(t, apiKey, *opts.APIKey)
	assert.NotNil(t, opts.Headers)
	assert.Equal(t, headerValue, opts.Headers.Get(headerKey))
	assert.NotNil(t, opts.QueryParams)
	assert.Equal(t, queryValue, opts.QueryParams.Get(queryKey))
}

func TestRequiredEdgeHeaders(t *testing.T) {
	edgeID := "test_edge_id"
	edgeIPAddress := "127.0.0.1"

	opts := &models.RequestOptions{}
	models.RequiredEdgeHeaders(edgeID, edgeIPAddress)(opts)

	assert.NotNil(t, opts.Headers)
	assert.Equal(t, edgeID, opts.Headers.Get(models.HeaderEdgeID))
	assert.Equal(t, edgeIPAddress, opts.Headers.Get(models.HeaderEdgeIPAddress))
}

func TestEdgeUserAgent(t *testing.T) {
	userAgent := "test_user_agent"

	opts := &models.RequestOptions{}
	models.EdgeUserAgent(userAgent)(opts)

	assert.NotNil(t, opts.Headers)
	assert.Equal(t, userAgent, opts.Headers.Get(models.HeaderEdgeUserAgent))
}
