
Users of the Launchpad product will need to pass in certain headers in order to make API requests.

```golang
params := &models.GetGroupedDailyAggsParams{
    Locale:     models.US,
    MarketType: models.Stocks,
    Date:       models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)),
}

res, err := c.GetGroupedDailyAggs(context.Background(), params,
    models.RequiredEdgeHeaders("EDGE_USER_ID", "EDGE_USER_IP_ADDRESS"),
)
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```

Launchpad users can also provide the optional User Agent value describing their Edge User's origination request.

```golang
params := &models.GetGroupedDailyAggsParams{
    Locale:     models.US,
    MarketType: models.Stocks,
    Date:       models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)),
}

res, err := c.GetGroupedDailyAggs(context.Background(), params,
    models.RequiredEdgeHeaders("EDGE_USER_ID", "EDGE_USER_IP_ADDRESS"),
    models.EdgeUserAgent("EDGE_USER_AGENT"),
)
if err != nil {
    log.Fatal(err)
}
log.Print(res) // do something with the result
```