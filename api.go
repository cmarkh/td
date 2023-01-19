package td

import "time"


//Price is the price schema for TD API
type Price struct {
	Close    float64
	Datetime datetime
	High     float64
	Low      float64
	Open     float64
	Volume   int64
}

func ToDatetime(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

type datetime int64

func (t datetime) Time() time.Time {
	return time.Unix(int64(t)/1000, 0)
}

func (t datetime) UnixSeconds() int64 {
	return int64(t) / 1000
}
