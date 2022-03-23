package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/stretchr/testify/assert"
)

const (
	listResourcePath = "/resource/{param}"
)

type Client struct {
	client.Client
}

func (c *Client) ListResource(ctx context.Context, params ListResourceParams, options ...client.Option) *ResourceIter {
	return &ResourceIter{
		Iter: client.GetIter(ctx, params.String(), func(url string) (client.ListResponse, []interface{}, error) {
			res := &ResourceResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

type ResourceIter struct {
	*client.Iter
}

func (it *ResourceIter) Resource() *Resource {
	if it.Item() != nil {
		return it.Item().(*Resource)
	}
	return nil
}

func (it *ResourceIter) ResourceList() *ResourceResponse {
	if it.Page() != nil {
		return it.Page().(*ResourceResponse)
	}
	return nil
}

type ResourceResponse struct {
	client.BaseResponse
	Results []*Resource `json:"results,omitempty"`
}

type Resource struct {
	Field string `json:"field"`
}

type ListResourceParams struct {
	Param       string
	QueryParams ListResourceQueryParams
}

type ListResourceQueryParams struct {
	Param *string
}

func (p ListResourceParams) Path() map[string]string {
	return map[string]string{
		"param": p.Param,
	}
}

func (p ListResourceParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Param != nil {
		q.Set("timestamp", *p.QueryParams.Param)
	}

	return q
}

func (p ListResourceParams) String() string {
	path := strings.ReplaceAll(listResourcePath, "{param}", url.PathEscape(p.Param))

	q := p.Query().Encode()
	if q != "" {
		path += "?" + q
	}

	return path
}

func TestListResource(t *testing.T) {
	c := Client{Client: client.New("API_KEY")}

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	resource1 := Resource{Field: "field1"}
	resource2 := Resource{Field: "field2"}
	expectedRes1 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     2,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/param1?cursor=NEXT",
			},
		},
		Results: []*Resource{&resource1, &resource2},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/param1", ReqHandler(200, expectedRes1))

	resource3 := Resource{Field: "field3"}
	expectedRes2 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     1,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/param1?cursor=NEXTER",
			},
		},
		Results: []*Resource{&resource3},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/param1?cursor=NEXT", ReqHandler(200, expectedRes2))

	expectedRes3 := ResourceResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req3",
			Count:     0,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/resource/param1?cursor=NEXTER",
			},
		},
	}
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/param1?cursor=NEXTER", ReqHandler(200, expectedRes3))

	iter := c.ListResource(context.Background(), ListResourceParams{
		Param: "param1",
	})

	// verify the first page
	assert.Nil(t, iter.Err())
	assert.Equal(t, &expectedRes1, iter.Page())
	assert.Nil(t, iter.Resource())
	// verify the first and second quotes
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &resource1, iter.Resource())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &resource2, iter.Resource())

	// verify the second page
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	// verify the third quote
	assert.Equal(t, &expectedRes2, iter.ResourceList())
	assert.Equal(t, &resource3, iter.Resource())

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
	httpmock.RegisterResponder("GET", "https://api.polygon.io/resource/param1", ReqHandler(404, expectedRes))

	iter := c.ListResource(context.Background(), ListResourceParams{
		Param: "param1",
	})

	// iter.Next() should return false and the error should not be nil
	assert.False(t, iter.Next())
	assert.NotNil(t, iter.Err())
	assert.Equal(t, expectedErr.Error(), iter.Err().Error())

	// subsequent calls to iter.Next() should be false, item should be nil, page should be an empty response
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Item())
	assert.Equal(t, &ResourceResponse{}, iter.Page())
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
