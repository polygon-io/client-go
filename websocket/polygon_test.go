package polygonws_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	polygonws "github.com/polygon-io/client-go/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)

func TestMain(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return // skip in CI for now
	}

	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	c := polygonws.New(apiKey)
	conn, aggs, err := c.StreamSecondAggs(ctx, models.Stocks, "*")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	for agg := range aggs {
		fmt.Println(agg)
	}
}
