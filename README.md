# Polygon Go Client

<!-- todo: add a codecov badge -->
<!-- todo: figure out a way to show all build statuses -->

[![Docs][doc-img]][doc] [![Build][build-img]][build]

The official Go client library for the [Polygon](https://polygon.io/) REST and WebSocket API.

`go get github.com/polygon-io/client-go`

See the [docs](https://polygon.io/docs/stocks/getting-started) for more details on our API. 

This client is still in pre-release. It only supports the REST API but WebSocket support is coming soon. The public interface is relatively stable at this point but is still liable to change slightly until we release v1.

## REST API Client

To get started, you'll need to import two main packages.

```golang
import (
	polygon "github.com/polygon-io/client-go"
	"github.com/polygon-io/client-go/rest/models"
)
```

Next, create a new client with your [API key](https://polygon.io/dashboard/signup).

```golang
c := polygon.NewClient("YOUR_API_KEY")
```

### Using the client

After creating the client, making calls to the Polygon API is simple.

```golang
params := models.GetAllTickersSnapshotParams{
    Locale:     "us",
    MarketType: "stocks",
}.WithTickers("AAPL,MSFT")

res, err := c.Snapshot.GetAllTickersSnapshot(context.Background(), params)
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```

### Pagination

Our list methods return iterators that handle pagination for you.

```golang
params := models.ListTradesParams{Ticker: "AAPL"}.
    WithTimestamp(models.GTE, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
    WithOrder(models.Asc)

iter, err := c.Trades.ListTrades(context.Background(), params)
if err != nil {
    log.Fatal(err)
}

for iter.Next() { // iter.Next() advances the iterator to the next value in the list
    log.Print(iter.Trade()) // do something with the current value
}
if iter.Err() != nil {
    log.Fatal(err) // the loop will break if an error occurs while iterating
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

res, err := c.Aggs.GetGroupedDailyAggs(context.Background(), params,
    models.APIKey("YOUR_OTHER_API_KEY"),
    models.Header("X-CUSTOM-HEADER", "VALUE"),
    models.QueryParam("adjusted", strconv.FormatBool(true)))
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```

## Contributing

For now, we're generally not accepting pull requests from outside contributors but we're open to bug reports and feature requests. Or if you have more general feedback, feel free to reach out on our [Slack channel](https://polygon.io/contact).

-------------------------------------------------------------------------------

[doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go
[doc]: https://pkg.go.dev/github.com/polygon-io/client-go
[build-img]: https://github.com/polygon-io/client-go/actions/workflows/test.yml/badge.svg
[build]: https://github.com/polygon-io/client-go/actions
