package models_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	listResourcePath = "/resource/{ticker}"
)

type Client struct {
	client.Client
}

type ListResourceIter struct {
	models.Iter
}

func (it *ListResourceIter) Resource() Resource {
	if it.Item() != nil {
		return it.Item().(Resource)
	}
	return Resource{}
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

func (c *Client) ListResource(ctx context.Context, params *ListResourceParams, options ...models.RequestOption) *ListResourceIter {
	uri, err := c.EncodeParams(listResourcePath, params)
	if err != nil {
		return &ListResourceIter{Iter: models.NewIterErr(ctx, err)}
	}

	return &ListResourceIter{
		Iter: models.NewIter(ctx, uri, func(uri string) (models.ListResponse, []interface{}, error) {
			res := &ListResourceResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

func TestListResource(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	resource1 := Resource{Price: "price1"}
	registerResponder(200, "https://api.polygon.io/resource/ticker1", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXT",
			},
		},
		Results: []Resource{resource1},
	})

	resource2 := Resource{Price: "price2"}
	resource3 := Resource{Price: "price3"}
	registerResponder(200, "https://api.polygon.io/resource/ticker1?cursor=NEXT", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     2,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXTER",
			},
		},
		Results: []Resource{resource2, resource3},
	})

	registerResponder(200, "https://api.polygon.io/resource/ticker1?cursor=NEXTER", ListResourceResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req3",
			Count:     0,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXTER",
			},
		},
	})

	iter := c.ListResource(context.Background(), &ListResourceParams{
		Ticker: "ticker1",
	})

	// verify the first page
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Resource())
	// verify the first and second quotes
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, resource1, iter.Resource())

	// verify the second page
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, resource2, iter.Resource())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, resource3, iter.Resource())

	// verify the third page (end of list)
	assert.False(t, iter.Next()) // this should be false since the third page has no results
	assert.Nil(t, iter.Err())
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
	registerResponder(404, "https://api.polygon.io/resource/ticker1", ListResourceResponse{
		BaseResponse: baseRes,
	})

	iter := c.ListResource(context.Background(), &ListResourceParams{
		Ticker: "ticker1",
	})

	// iter.Next() should return false and the error should not be nil
	assert.NotNil(t, iter.Err())
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Err())
	assert.Equal(t, expectedErr.Error(), iter.Err().Error())

	// subsequent calls to iter.Next() should be false, item should be not nil, page should be an empty response
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Resource())
}

func TestListResourceEncodeError(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	iter := c.ListResource(context.Background(), nil)

	// iter.Next() should return false and the error should not be nil
	assert.NotNil(t, iter.Err())
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Err())

	// subsequent calls to iter.Next() should be false, item should be not nil, page should be an empty response
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Resource())
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
