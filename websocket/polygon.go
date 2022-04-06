package polygonws

import (
	"context"
	"encoding/json"

	"github.com/polygon-io/client-go/websocket/client"
	"github.com/polygon-io/client-go/websocket/models"
)

// PolygonClient defines a client to the Polygon WebSocket API.
type PolygonClient struct {
	client.Client
}

func New(apiKey string) *PolygonClient {
	c := client.New(apiKey)
	return &PolygonClient{
		Client: c,
	}
}

// StreamSecondAggs streams real-time second aggregates for a given ticker symbol. "*" will stream data for
// all tickers of a given market type. For more details see https://polygon.io/docs/stocks/ws_stocks_a.
// Callers of this method must close the connection:
//   conn, aggs, err := c.StreamSecondAggs(context.TODO(), models.Stocks, "AAPL")
//   if err != nil {
//       return err
//   }
//   defer conn.Close()
func (pc *PolygonClient) StreamSecondAggs(ctx context.Context, market models.MarketType, ticker string) (*client.Conn, models.SecondAggs, error) {
	conn, err := pc.Connect(market, "A."+ticker)
	if err != nil {
		return nil, nil, err
	}

	data := make(chan []byte, 10000)
	go conn.Collect(data)

	secondAggs := make(models.SecondAggs, 10000)
	// go conn.Process(ctx, data, secondAggs)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(secondAggs)
				return
			case msgBytes := <-data:
				var aggs []models.Agg
				err := json.Unmarshal(msgBytes, &aggs)
				if err != nil {
					continue // ignore malformed data
				}
				for _, agg := range aggs {
					secondAggs <- agg
				}
			}
		}
	}()

	return conn, secondAggs, nil
}
