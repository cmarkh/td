package td

//https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory#

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cmarkh/errs"
)

func QuoteEquity(ticker string) (quote Equity, err error) {
	quotes, err := QuoteEquities(ticker)
	if err != nil {
		return quote, errs.WrapErr(err)
	}
	return quotes[ticker], nil
}

// QuoteEquities gets the quotes for the given tickers - map[Ticker]Equity
func QuoteEquities(tickers ...string) (quotes map[string]Equity, err error) {
	err = quote(tickers, &quotes)
	if err != nil {
		return quotes, errs.WrapErr(err)
	}
	return
}

// quote return value in quotes interface
func quote(tickers []string, quotes interface{}) (err error) {
	u := quoteurl(tickers)

	resp, err := http.Get(u)
	if err != nil {
		return errs.WrapErr(err, u)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errs.WrapErr(fmt.Errorf("unexpected http GET status: %s", resp.Status), u)
	}

	//body, err := io.ReadAll(resp.Body)
	//fmt.Println(string(body))

	err = json.NewDecoder(resp.Body).Decode(quotes)
	if err != nil {
		return errs.WrapErr(err, u)
	}

	return
}

func quoteurl(tickers []string) (u string) {
	u = "https://api.tdameritrade.com/v1/marketdata/quotes?"

	v := url.Values{}
	v.Add("apikey", apiKey)

	t := ""
	for i, ticker := range tickers { //comma separated tickers
		if i > 0 {
			t += ","
		}
		t += ticker
	}
	v.Add("symbol", t)

	return u + v.Encode()
}
