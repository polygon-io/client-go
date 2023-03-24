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
		input  []byte   // input
		expect Response // expected output
		err    error    // expected error
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
		"unexpected format": {
			input: []byte(`{
				"time": "2022-10T22:3"
			}`),
			err: &time.ParseError{Layout: "2006-01-02T15:04:05Z", Value: "2022-10T22:3", LayoutElem: "-", ValueElem: "T22:3", Message: ""},
		},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			var res Response
			err := json.Unmarshal(tc.input, &res)
			assert.Equal(t, tc.err, err)

			expect := time.Time(tc.expect.Time).String()
			actual := time.Time(res.Time).String()
			if expect != actual {
				t.Errorf("%v: expected { %v } got { %v }", desc, expect, actual)
			}
		})
	}
}

func TestTime_MarshalUnmarshal(t *testing.T) {
	ti := models.Time(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC))
	data, err := json.Marshal(&ti)
	assert.NoError(t, err)
	assert.Equal(t, "\"2022-02-02T00:00:00.000Z\"", string(data))

	var unmarshaled models.Time
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, ti, unmarshaled)
}

func TestDate_MarshalUnmarshal(t *testing.T) {
	date := models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC))
	data, err := json.Marshal(&date)
	assert.NoError(t, err)
	assert.Equal(t, "\"2022-02-02\"", string(data))

	var unmarshaled models.Date
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, date, unmarshaled)
}

func TestMillis_MarshalUnmarshal(t *testing.T) {
	millis := models.Millis(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC))
	data, err := json.Marshal(&millis)
	assert.NoError(t, err)

	var unmarshaled models.Millis
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, time.Time(millis).UnixMilli(), time.Time(unmarshaled).UnixMilli())
}

func TestNanos_MarshalUnmarshal(t *testing.T) {
	nanos := models.Nanos(time.Now())
	data, err := json.Marshal(&nanos)
	assert.NoError(t, err)

	var unmarshaled models.Nanos
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, time.Time(nanos).UnixNano(), time.Time(unmarshaled).UnixNano())
}
