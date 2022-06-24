package encoder

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/polygon-io/client-go/rest/models"
)

// Encoder defines a path and query param encoder that plays nicely with the Polygon REST API.
type Encoder struct {
	validate     *validator.Validate
	pathEncoder  *form.Encoder
	queryEncoder *form.Encoder
}

// New returns a new path and query param encoder.
func New() *Encoder {
	return &Encoder{
		validate:     validator.New(),
		pathEncoder:  newEncoder("path"),
		queryEncoder: newEncoder("query"),
	}
}

// EncodeParams encodes path and query params and returns a valid request URI.
func (e *Encoder) EncodeParams(path string, params any) (string, error) {
	if err := e.validateParams(params); err != nil {
		return "", err
	}

	uri, err := e.encodePath(path, params)
	if err != nil {
		return "", err
	}

	query, err := e.encodeQuery(params)
	if err != nil {
		return "", err
	} else if query != "" {
		uri += "?" + query
	}

	return uri, nil
}

func (e *Encoder) validateParams(params any) error {
	if err := e.validate.Struct(params); err != nil {
		return fmt.Errorf("invalid request params: %w", err)
	}
	return nil
}

func (e *Encoder) encodePath(uri string, params any) (string, error) {
	val, err := e.pathEncoder.Encode(&params)
	if err != nil {
		return "", fmt.Errorf("error encoding path params: %w", err)
	}

	pathParams := map[string]string{}
	for k, v := range val {
		pathParams[k] = v[0] // only accept the first one for a given key
	}

	for k, v := range pathParams {
		uri = strings.ReplaceAll(uri, fmt.Sprintf("{%s}", k), url.PathEscape(v))
	}

	return uri, nil
}

func (e *Encoder) encodeQuery(params any) (string, error) {
	query, err := e.queryEncoder.Encode(&params)
	if err != nil {
		return "", fmt.Errorf("error encoding query params: %w", err)
	}
	return query.Encode(), nil
}

func newEncoder(tag string) *form.Encoder {
	e := form.NewEncoder()
	e.SetMode(form.ModeExplicit)
	e.SetTagName(tag)

	e.RegisterCustomTypeFunc(func(x any) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Time)).Format("2006-01-02T15:04:05.000Z"))}, nil
	}, models.Time{})
	e.RegisterCustomTypeFunc(func(x any) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Date)).Format("2006-01-02"))}, nil
	}, models.Date{})
	e.RegisterCustomTypeFunc(func(x any) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Millis)).UnixMilli())}, nil
	}, models.Millis{})
	e.RegisterCustomTypeFunc(func(x any) ([]string, error) {
		if isDay(time.Time(x.(models.Nanos))) {
			// endpoints that have nanosecond timestamp query parameters are expected to
			// also work with date strings if a user wants all data from a specific day
			return []string{fmt.Sprint(time.Time(x.(models.Nanos)).Format("2006-01-02"))}, nil
		}
		return []string{fmt.Sprint(time.Time(x.(models.Nanos)).UnixNano())}, nil
	}, models.Nanos{})

	return e
}

func isDay(t time.Time) bool {
	if t.Hour() != 0 || t.Minute() != 0 || t.Second() != 0 || t.Nanosecond() != 0 {
		return false
	}
	return true
}
