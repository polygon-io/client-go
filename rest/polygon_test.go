package polygon_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	polygon "github.com/polygon-io/client-golang/rest"
	"github.com/polygon-io/client-golang/rest/aggregates"
	"github.com/polygon-io/client-golang/rest/client"
)

// todo: write some tests, just verifying that the client works for now
func TestAggs(t *testing.T) {
	c := polygon.New(client.HTTPBaseConfig{
		URL:        "https://api.polygon.io",
		Key:        os.Getenv("API_KEY"),
		MaxRetries: 3,
	})

	from := time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)
	to := from.Add(10 * 24 * time.Hour)
	res, err := c.Aggregates.Get(context.Background(), "AAPL", 1, "day", from, to, &aggregates.GetQueryParams{
		Adjusted: true,
		Sort:     "asc",
		Limit:    10,
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := res.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
