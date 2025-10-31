package iter_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"github.com/massive-com/client-go/v2/rest/client"
	"github.com/massive-com/client-go/v2/rest/iter"
	"github.com/massive-com/client-go/v2/rest/models"
)

const (
	listResourcePath = "/resource/{ticker}"
)

type Client struct {
	client.Client
}

type ListResourceResponse struct {
	models.BaseResponse
	Results []Resource `json:"results,omitempty"`
}

type Resource struct {
	Price string `json:"price"`
}

type ListResourceParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Timestamp *string `query:"timestamp"`
}

func (c *Client) ListResource(ctx context.Context, params *ListResourceParams, options ...models.RequestOption) *iter.Iter[Resource] {
	return iter.NewIter(ctx, listResourcePath, params, func(uri string) (iter.ListResponse, []Resource, error) {
		res := &ListResourceResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

func TestListResource(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	resource1 := Resource{Price: "price1"}
	registerResponder(200, "https://api.massive.com/resource/ticker1", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.massive.com/resource/ticker1?cursor=NEXT",
			},
		},
		Results: []Resource{resource1},
	})

	resource2 := Resource{Price: "price2"}
	resource3 := Resource{Price: "price3"}
	registerResponder(200, "https://api.massive.com/resource/ticker1?cursor=NEXT", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     2,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.massive.com/resource/ticker1?cursor=NEXTER",
			},
		},
		Results: []Resource{resource2, resource3},
	})

	registerResponder(200, "https://api.massive.com/resource/ticker1?cursor=NEXTER", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req3",
			Count:     0,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.massive.com/resource/ticker1?cursor=NEXTER",
			},
		},
	})

	it := c.ListResource(context.Background(), &ListResourceParams{
		Ticker: "ticker1",
	})

	// verify the first page
	assert.Nil(t, it.Err())
	assert.NotNil(t, it.Item())
	// verify the first and second quotes
	assert.True(t, it.Next())
	assert.Nil(t, it.Err())
	assert.Equal(t, resource1, it.Item())

	// verify the second page
	assert.True(t, it.Next())
	assert.Nil(t, it.Err())
	assert.Equal(t, resource2, it.Item())
	assert.True(t, it.Next())
	assert.Nil(t, it.Err())
	assert.Equal(t, resource3, it.Item())

	// verify the third page (end of list)
	assert.False(t, it.Next()) // this should be false since the third page has no results
	assert.Nil(t, it.Err())
}

func TestListResourceError(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	baseRes := models.BaseResponse{
		Status:       "NOT FOUND",
		RequestID:    "req1",
		ErrorMessage: "resource not found",
	}
	expectedErr := models.ErrorResponse{
		StatusCode:   404,
		BaseResponse: baseRes,
	}
	registerResponder(404, "https://api.massive.com/resource/ticker1", ListResourceResponse{
		BaseResponse: baseRes,
	})

	it := c.ListResource(context.Background(), &ListResourceParams{
		Ticker: "ticker1",
	})

	// it.Next() should return false and the error should not be nil
	assert.NotNil(t, it.Err())
	assert.False(t, it.Next())
	assert.NotNil(t, it.Err())
	assert.Equal(t, expectedErr.Error(), it.Err().Error())

	// subsequent calls to it.Next() should be false, item should be not nil, page should be an empty response
	assert.False(t, it.Next())
	assert.NotNil(t, it.Item())
}

func TestListResourceEncodeError(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	it := c.ListResource(context.Background(), nil)

	// it.Next() should return false and the error should not be nil
	assert.NotNil(t, it.Err())
	assert.False(t, it.Next())
	assert.NotNil(t, it.Err())

	// subsequent calls to it.Next() should be false, item should be not nil, page should be an empty response
	assert.False(t, it.Next())
	assert.NotNil(t, it.Item())
}

func registerResponder(status int, url string, res ListResourceResponse) {
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(res)
			if err != nil {
				return nil, err
			}
			resp := httpmock.NewStringResponse(status, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)
}
