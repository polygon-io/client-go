# Polygon Go Client

<!-- todo: add a codecov badge -->
<!-- todo: figure out a way to show all build statuses -->

[![Build][build-img]][build]

The official Go client library for the [Polygon](https://polygon.io/) REST and WebSocket API.

`go get github.com/polygon-io/client-go`

See the [docs](https://polygon.io/docs/stocks/getting-started) for more details on our API. 

This client is still in pre-release. It only supports the REST API but WebSocket support is coming soon. The public interface is relatively stable at this point but is still liable to change slightly until we release v1.

## REST API Client

[![rest-docs][rest-doc-img]][rest-doc]

To get started, you'll need to import two main packages.

```golang
import (
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)
```

Next, create a new client with your [API key](https://polygon.io/dashboard/signup).

```golang
c := polygon.New("YOUR_API_KEY")
```

### Using the client

After creating the client, making calls to the Polygon API is simple.

```golang
params := models.GetAllTickersSnapshotParams{
    Locale:     models.US,
    MarketType: models.Stocks,
}.WithTickers("AAPL,MSFT")

res, err := c.GetAllTickersSnapshot(context.Background(), params)
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```

### Pagination

Our list methods return iterators that handle pagination for you.

```golang
// create a new iterator
params := models.ListTradesParams{Ticker: "AAPL"}.
    WithTimestamp(models.GTE, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
    WithOrder(models.Asc)
iter := c.ListTrades(context.Background(), params)

// iter.Next() advances the iterator to the next value in the list
for iter.Next() {
    log.Print(iter.Item()) // do something with the current value
}

// if the loop breaks, it has either reached the end of the list or an error has occurred
// you can check if something went wrong with iter.Err()
if iter.Err() != nil {
    log.Fatal(iter.Err())
}
```

### Request options

Advanced users may want to add additional headers or query params to a given request.

```golang
params := &models.GetGroupedDailyAggsParams{
    Locale:     models.US,
    MarketType: models.Stocks,
    Date:       models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)),
}

res, err := c.GetGroupedDailyAggs(context.Background(), params,
    models.APIKey("YOUR_OTHER_API_KEY"),
    models.Header("X-CUSTOM-HEADER", "VALUE"),
    models.QueryParam("adjusted", strconv.FormatBool(true)))
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```

## WebSocket Client

[![ws-docs][ws-doc-img]][ws-doc]

Import the WebSocket client and models packages to get started.

```golang
import (
    polygonws "github.com/polygon-io/client-go/websocket"
    "github.com/polygon-io/client-go/websocket/models"
)
```

Next, create a new client with your API key and a couple other config options.

```golang
// create a new client
c, err := polygonws.New(polygonws.Config{
    APIKey:    "YOUR_API_KEY",
    Feed:      polygonws.RealTime,
    Market:    polygonws.Stocks,
})
if err != nil {
    log.Fatal(err)
}
defer c.Close() // the user of this client must close it

// connect to the server
if err := c.Connect(); err != nil {
    log.Fatal(err)
}
```

The client automatically reconnects to the server when the connection is dropped. By default, it will attempt to reconnect indefinitely but the number of retries is configurable. When the client successfully reconnects, it automatically resubscribes to any topics that were set before the disconnect.

### Using the client

After creating a client, subscribe to a topic and start accessing data.

```golang
// subscribe to a specific ticker
if err := c.Subscribe(polygonws.StocksSecAggs, "AAPL"); err != nil {
    log.Fatal(err)
}

for {
    // do something with the first message
    if out := c.Output(); out != nil {
        log.Print(out)
        break
    }
}

// passing a topic by itself will unsubscribe from all tickers
if err := c.Unsubscribe(polygonws.StocksSecAggs); err != nil {
    log.Fatal(err)
}
```

This client lets you subscribe to as many topics as desired. Currently, all of the data is stored in a single output channel.

```golang
// passing a topic by itself will subscribe to all tickers
if err := c.Subscribe(polygonws.StocksSecAggs); err != nil {
    log.Fatal(err)
}
if err := c.Subscribe(polygonws.StocksTrades, "TSLA", "GME"); err != nil {
    log.Fatal(err)
}

for {
    out := c.Output()
    switch out.(type) {
    case models.EquityAgg:
        // do something with the agg
    case models.EquityTrade:
        // do something with the trade
    }
}
```

## Contributing

For now, we're generally not accepting pull requests from outside contributors but we're open to bug reports and feature requests. Or if you have more general feedback, feel free to reach out on our [Slack channel](https://polygon.io/contact).

-------------------------------------------------------------------------------

[rest-doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go/rest
[rest-doc]: https://pkg.go.dev/github.com/polygon-io/client-go/rest
[ws-doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go/websocket
[ws-doc]: https://pkg.go.dev/github.com/polygon-io/client-go/websocket
[build-img]: https://github.com/polygon-io/client-go/actions/workflows/test.yml/badge.svg
[build]: https://github.com/polygon-io/client-go/actions
