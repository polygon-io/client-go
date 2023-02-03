# Polygon Go Client
![Coverage](https://img.shields.io/badge/Coverage-40.9%25-yellow)

<!-- todo: add a codecov badge -->
<!-- todo: figure out a way to show all build statuses -->
<!-- todo: consider moving some stuff into separate readmes -->

[![docs][doc-img]][doc] [![Build][build-img]][build] [![Go Report Card][report-card-img]][report-card]

The official Go client library for the [Polygon](https://polygon.io/) REST and WebSocket API.

`go get github.com/polygon-io/client-go`

This client makes use of Go generics and thus requires Go 1.18. See the [docs](https://polygon.io/docs/stocks/getting-started) for more details on our API.

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

Or create a client with a custom HTTP client implementation.

```golang
hc := http.Client{} // some custom HTTP client
c := polygon.NewWithClient("YOUR_API_KEY", hc)
```

### Using the client

After creating the client, making calls to the Polygon API is simple.

```golang
params := models.GetTickerDetailsParams{
    Ticker: "AAPL",
}.WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)))

res, err := c.GetTickerDetails(context.Background(), params)
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

We also provide a builder method to make it easier to retrieve all trades and quotes for a specific day.

```golang
params := models.ListQuotesParams{Ticker: "AAPL"}.
    WithDay(2021, 7, 22). // get all quotes for July 22, 2021
    WithOrder(models.Asc)
iter := c.ListQuotes(context.Background(), params)

for iter.Next() {
    log.Print(iter.Item())
}
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

#### Launchpad Usage

Users of the Launchpad product will need to pass in certain headers in order to make API requests.
Example can be found [here](./rest/example/launchpad).

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
    log.Error(err)
    return
}
```

The client automatically reconnects to the server when the connection is dropped. By default, it will attempt to reconnect indefinitely but the number of retries is configurable. When the client successfully reconnects, it automatically resubscribes to any topics that were set before the disconnect.

### Using the client

After creating a client, subscribe to one or more topics and start accessing data. Currently, all of the data is pushed to a single output channel.

```golang
// passing a topic by itself will subscribe to all tickers
if err := c.Subscribe(polygonws.StocksSecAggs); err != nil {
    log.Fatal(err)
}
if err := c.Subscribe(polygonws.StocksTrades, "TSLA", "GME"); err != nil {
    log.Fatal(err)
}

for {
    select {
    case err := <-c.Error(): // check for any fatal errors (e.g. auth failed)
        log.Fatal(err)
    case out, more := <-c.Output(): // read the next data message
        if !more {
            return
        }

        switch out.(type) {
        case models.EquityAgg:
            log.Print(out) // do something with the agg
        case models.EquityTrade:
            log.Print(out) // do something with the trade
        }
    }
}
```

See the [full example](./websocket/example/main.go) for more details on how to use this client effectively.

## Release planning

This client will attempt to follow the release cadence of our API. When endpoints are deprecated and newer versions are added, the client will maintain two methods in a backwards compatible way (e.g. `ListTrades` and `ListTradesV4(...)`). When deprecated endpoints are removed from the API, we'll rename the versioned method (e.g. `ListTradesV4(...)` -> `ListTrades(...)`), remove the old method, and release a new major version of the client. The goal is to give users ample time to upgrade to newer versions of our API _before_ we bump the major version of the client, and in general, we'll try to bundle breaking changes like this to avoid frequent major version bumps.

There are a couple exceptions to this. When we find small breaking issues with this client library (e.g. incorrect response types), we may decide to release them under the same major version. These changes will be clearly outlined in the release notes. Also, methods that fall under the VX client are considered experimental and may be modified or deprecated as needed. We'll call out any breaking changes to VX endpoints in our release notes to make using them easier.

## Contributing

If you found a bug or have an idea for a new feature, please first discuss it with us by [submitting a new issue](https://github.com/polygon-io/client-go/issues/new/choose). We will respond to issues within at most 3 weeks. We're also open to volunteers if you want to submit a PR for any open issues but please discuss it with us beforehand. PRs that aren't linked to an existing issue or discussed with us ahead of time will generally be declined. If you have more general feedback or want to discuss using this client with other users, feel free to reach out on our [Slack channel](https://polygon-io.slack.com/archives/C03FCSBSAFL).

-------------------------------------------------------------------------------

[doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go
[doc]: https://pkg.go.dev/github.com/polygon-io/client-go
[rest-doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go/rest
[rest-doc]: https://pkg.go.dev/github.com/polygon-io/client-go/rest
[ws-doc-img]: https://pkg.go.dev/badge/github.com/polygon-io/client-go/websocket
[ws-doc]: https://pkg.go.dev/github.com/polygon-io/client-go/websocket
[build-img]: https://github.com/polygon-io/client-go/actions/workflows/test.yml/badge.svg
[build]: https://github.com/polygon-io/client-go/actions
[report-card-img]: https://goreportcard.com/badge/github.com/polygon-io/client-go
[report-card]: https://goreportcard.com/report/github.com/polygon-io/client-go