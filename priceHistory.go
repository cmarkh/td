package td

//https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory#

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cmarkh/errs"
)

// HistoryRequest TD API
// Tick every Frequency FrequencyType for Period PeriodTypes
// Tick every 1 Minute for 10 Days (for example)
type HistoryRequest struct {
	Ticker                string
	Frequency             int
	FrequencyType         string
	Period                int //If startDate and endDate are provided, period should not be provided
	PeriodType            string
	StartDate             time.Time
	EndDate               time.Time
	NeedExtendedHoursData bool
	Authorization         string //Supply an access token to make an authenticated request. The format is Bearer <access token>.
}

// History TD API
type History struct {
	Symbol  string
	Empty   bool
	Candles []Price
}

// PriceHistory gets the price history of the given ticker for the given date range
func PriceHistory(request HistoryRequest) (history History, err error) {
	u := historyurl(request)

	resp, err := http.Get(u)
	if err != nil {
		return history, errs.WrapErr(err, u)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return history, errs.WrapErr(fmt.Errorf("unexpected http GET status: %s", resp.Status), u)
	}

	//body, err := io.ReadAll(resp.Body)
	//fmt.Println(body)

	err = json.NewDecoder(resp.Body).Decode(&history)
	if err != nil {
		return history, errs.WrapErr(err, u)
	}

	return
}

func historyurl(request HistoryRequest) (u string) {
	u = "https://api.tdameritrade.com/v1/marketdata/"

	ticker := strings.ReplaceAll(request.Ticker, "/", ".")
	u += url.PathEscape(ticker)

	u += "/pricehistory?"

	v := url.Values{}
	v.Add("apikey", apiKey)

	if request.PeriodType != "" {
		v.Add("periodType", request.PeriodType)
	}
	if request.Period != 0 {
		v.Add("period", fmt.Sprint(request.Period))
	}

	if request.FrequencyType != "" {
		v.Add("frequencyType", request.FrequencyType)
	}
	if request.Frequency != 0 {
		v.Add("frequency", fmt.Sprint(request.Frequency))
	}

	if !request.EndDate.IsZero() { //If startDate and endDate are provided, period should not be provided
		v.Add("endDate", fmt.Sprint(ToDatetime(request.EndDate))) //milliseconds since epoch
		v.Del("period")
	}
	if !request.StartDate.IsZero() {
		v.Add("startDate", fmt.Sprint(ToDatetime(request.StartDate))) //milliseconds since epoch
		v.Del("period")
	}

	v.Add("needExtendedHoursData", fmt.Sprint(request.NeedExtendedHoursData))

	return u + v.Encode()
}

func DailyHistory(ticker string, start time.Time, end time.Time) (history History, err error) {
	request := HistoryRequest{
		Ticker:                ticker,
		Frequency:             1,
		FrequencyType:         "daily",
		PeriodType:            "year",
		StartDate:             start,
		EndDate:               end,
		NeedExtendedHoursData: false,
	}

	history, err = PriceHistory(request)
	if err != nil {
		return history, errs.WrapErr(err)
	}

	return
}
