package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/stretchr/testify/assert"
)

const (
	listResourcePath = "/resource/{ticker}"
)

type Client struct {
	client.Client
}

type ResourceIter struct {
	client.Iter
}

func (it *ResourceIter) Resource() Resource {
	if it.Item() != nil {
		return it.Item().(Resource)
	}
	return Resource{}
}

type ResourceResponse struct {
	client.BaseResponse
	Results []Resource `json:"results,omitempty"`
}

type Resource struct {
	Price string `json:"price"`
}

type ListResourceParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Timestamp *string `query:"timestamp"`
}

func (c *Client) ListResource(ctx context.Context, params ListResourceParams, options ...client.Option) (*ResourceIter, error) {
	url, err := c.EncodeParams(listResourcePath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}

	return &ResourceIter{
		Iter: client.NewIter(ctx, url, func(url string) (client.ListResponse, []interface{}, error) {
			res := &ResourceResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}

func TestListResource(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	resource1 := Resource{Price: "price1"}
	expectedRes1 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXT",
			},
		},
		Results: []Resource{resource1},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/ticker1", ReqHandler(200, expectedRes1))

	resource2 := Resource{Price: "price2"}
	resource3 := Resource{Price: "price3"}
	expectedRes2 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     2,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXTER",
			},
		},
		Results: []Resource{resource2, resource3},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/ticker1?cursor=NEXT", ReqHandler(200, expectedRes2))

	expectedRes3 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req3",
			Count:     0,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/ticker1?cursor=NEXTER",
			},
		},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/ticker1?cursor=NEXTER", ReqHandler(200, expectedRes3))

	iter, err := c.ListResource(context.Background(), ListResourceParams{
		Ticker: "ticker1",
	})

	// verify the first page
	assert.Nil(t, err)
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

	baseRes := client.BaseResponse{
		Status:       "NOT FOUND",
		RequestID:    "req1",
		ErrorMessage: "resource not found",
	}
	expectedRes := ResourceResponse{
		BaseResponse: baseRes,
	}
	expectedErr := client.ErrorResponse{
		StatusCode:   404,
		BaseResponse: baseRes,
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/ticker1", ReqHandler(404, expectedRes))

	iter, err := c.ListResource(context.Background(), ListResourceParams{
		Ticker: "ticker1",
	})

	// iter.Next() should return false and the error should not be nil
	assert.Nil(t, err)
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Err())
	assert.Equal(t, expectedErr.Error(), iter.Err().Error())

	// subsequent calls to iter.Next() should be false, item should be nil, page should be an empty response
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Item())
}

func ReqHandler(code int, res ResourceResponse) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		resp := httpmock.NewStringResponse(code, string(b))
		resp.Header.Add("Content-Type", "application/json")
		return resp, nil
	}
}
