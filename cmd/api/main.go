package main

import (
	"fmt"
	"time"

	"github.com/cmarkh/errs"
	"github.com/cmarkh/td"
)

func main() {
	errs.SetLogFlags()

	fmt.Println(td.QuoteEquities("TSLA", "AAPL"))
	fmt.Println()

	hist1()
}

func hist1() {
	yday := time.Now().Add(time.Hour * -24)

	request := td.HistoryRequest{
		Ticker:                "AAPL",
		Frequency:             1,
		FrequencyType:         "minute",
		Period:                10,
		PeriodType:            "day",
		StartDate:             time.Date(yday.Year(), yday.Month(), yday.Day(), 9, 00, 0, 0, time.Local),
		EndDate:               time.Date(yday.Year(), yday.Month(), yday.Day(), 17, 00, 0, 0, time.Local),
		NeedExtendedHoursData: false,
	}

	history, err := td.PriceHistory(request)
	if err != nil {
		errs.Log(err)
		return
	}

	count := 0
	for _, candle := range history.Candles {
		fmt.Println(candle)
		count++
	}

	fmt.Println(count)
}
