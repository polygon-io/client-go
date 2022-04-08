package models

import (
	"context"
)

// Query defines a closure that domain specific iterators must implement. The implementation should
// include a call to the API and should return the API response with a separate slice of the results.
type Query func(string) (ListResponse, []interface{}, error)

// Iter defines an iterator type that should be used when implementing list methods. It's intended to
// be embedded in a domain specific iterator struct.
type Iter struct {
	ctx   context.Context
	query Query

	page    ListResponse
	item    interface{}
	results []interface{}

	err error
}

// NewIter returns a new initialized iterator. This method automatically makes the first query to populate
// the results. List methods should use this helper method when building domain specific iterators.
func NewIter(ctx context.Context, uri string, query Query) Iter {
	it := Iter{
		ctx:   ctx,
		query: query,
	}
	it.page, it.results, it.err = it.query(uri)
	return it
}

// Next moves the iterator to the next result.
func (it *Iter) Next() bool {
	if len(it.results) == 0 && it.page.NextPageURL() != "" {
		it.page, it.results, it.err = it.query(it.page.NextPageURL())
	}

	if it.err != nil || len(it.results) == 0 {
		return false
	}

	it.err = it.ctx.Err()
	if it.err != nil {
		return false
	}

	it.item = it.results[0]
	it.results = it.results[1:]
	return true
}

// Item returns the result that the iterator is currently pointing to.
func (it *Iter) Item() interface{} {
	return it.item
}

// Err returns any errors that occur during iteration.
func (it *Iter) Err() error {
	return it.err
}
