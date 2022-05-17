package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Time models.Time `json:"time"`
}

func TestUnmarshalTime(t *testing.T) {
	expect1, _ := time.Parse("2006-01-02T15:04:05Z", "2022-05-10T22:30:37Z")
	expect2, _ := time.Parse("2006-01-02T15:04:05.000Z", "2022-05-10T22:30:37.546Z")
	expect3, _ := time.Parse("2006-01-02T15:04:05-07:00", "2022-05-10T22:30:37-08:00")
	expect4, _ := time.Parse("2006-01-02T15:04:05.000-0700", "2022-05-10T22:30:37.546-0800")

	tests := map[string]struct {
		input  []byte
		expect Response
	}{
		"2006-01-02T15:04:05Z": {
			input: []byte(`{
				"time": "2022-05-10T22:30:37Z"
			}`),
			expect: Response{
				Time: models.Time(expect1),
			},
		},
		"2006-01-02T15:04:05.000Z": {
			input: []byte(`{
				"time": "2022-05-10T22:30:37.546Z"
			}`),
			expect: Response{
				Time: models.Time(expect2),
			},
		},
		"2006-01-02T15:04:05-07:00": {
			input: []byte(`{
				"time": "2022-05-10T22:30:37-08:00"
			}`),
			expect: Response{
				Time: models.Time(expect3),
			},
		},
		"2006-01-02T15:04:05.000-0700": {
			input: []byte(`{
				"time": "2022-05-10T22:30:37.546-0800"
			}`),
			expect: Response{
				Time: models.Time(expect4),
			},
		},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			var res Response
			err := json.Unmarshal(tc.input, &res)
			assert.Nil(t, err)
			expect := time.Time(tc.expect.Time).String()
			actual := time.Time(res.Time).String()
			if expect != actual {
				t.Errorf("%v: expected { %v } got { %v }", desc, expect, actual)
			}
		})
	}
}
