package client

// Iter defines an iterator type that should be used when implementing list methods. It's intended to
// be embedded in a domain specific iterator struct.
type Iter struct {
	query Query

	page    ListResponse
	current interface{}
	results []interface{}

	err error
}

// Query defines a closure that domain specific iterators must implement. The implementation should
// include a call to the API and should return the API response with a separate slice of the results.
type Query func(string) (ListResponse, []interface{}, error)

// GetIter returns a new initialized iterator. This method automatically makes the first query to
// populate the results. List methods should use this helper method when building domain specific
// iterators.
func GetIter(url string, query Query) *Iter {
	it := &Iter{query: query}
	it.page, it.results, it.err = it.query(url)
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
	it.current = it.results[0]
	it.results = it.results[1:]
	return true
}

// Page returns the current results page in the form of an API response.
func (it *Iter) Page() ListResponse {
	return it.page
}

// Current returns the result that the iterator is currently pointing to.
func (it *Iter) Current() interface{} {
	return it.current
}

// Err returns any errors that occur during iteration.
func (it *Iter) Err() error {
	return it.err
}
