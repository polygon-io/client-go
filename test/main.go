package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	p := polygon.New("FQeWYfnKTxXIv1uiIb0P4KTsrdCAnQaj")
	//from := time.UnixMilli(1609773010000)

	fromDate, _ := time.Parse("2006-01-02", "2021-01-04")
	to, _ := time.Parse("2006-01-02", "2021-02-01")
	aggs, _ := p.Aggs.GetAggs(context.Background(),
		models.GetAggsParams{
			Ticker:     "AAPL",
			Multiplier: 1,
			Resolution: "minute",
			From:       fromDate,
			To:         to,
		},
	)

	for i := range aggs.Results[:10] {
		fmt.Println(aggs.Results[i])
		fmt.Println(aggs.Results[i].Timestamp.ToTime(), reflect.TypeOf(aggs.Results[i].Timestamp.ToTime()))
		fmt.Println(aggs.Results[i].Timestamp.ToTime(), reflect.TypeOf(aggs.Results[i].Timestamp))

	}
}
