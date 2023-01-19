package td

//MutualFund quote
type MutualFund struct {
	Symbol          string  `json:"symbol"`
	Description     string  `json:"description"`
	Closeprice      float64 `json:"closePrice"`
	Netchange       float64 `json:"netChange"`
	Totalvolume     float64 `json:"totalVolume"`
	Tradetimeinlong float64 `json:"tradeTimeInLong"`
	Exchange        string  `json:"exchange"`
	Exchangename    string  `json:"exchangeName"`
	Digits          float64 `json:"digits"`
	Five2Wkhigh     float64 `json:"52WkHigh"`
	Five2Wklow      float64 `json:"52WkLow"`
	Nav             float64 `json:"nAV"`
	Peratio         float64 `json:"peRatio"`
	Divamount       float64 `json:"divAmount"`
	Divyield        float64 `json:"divYield"`
	Divdate         string  `json:"divDate"`
	Securitystatus  string  `json:"securityStatus"`
}

//Future quote
type Future struct {
	Symbol                string  `json:"symbol"`
	Bidpriceindouble      float64 `json:"bidPriceInDouble"`
	Askpriceindouble      float64 `json:"askPriceInDouble"`
	Lastpriceindouble     float64 `json:"lastPriceInDouble"`
	Bidid                 string  `json:"bidId"`
	Askid                 string  `json:"askId"`
	Highpriceindouble     float64 `json:"highPriceInDouble"`
	Lowpriceindouble      float64 `json:"lowPriceInDouble"`
	Closepriceindouble    float64 `json:"closePriceInDouble"`
	Exchange              string  `json:"exchange"`
	Description           string  `json:"description"`
	Lastid                string  `json:"lastId"`
	Openpriceindouble     float64 `json:"openPriceInDouble"`
	Changeindouble        float64 `json:"changeInDouble"`
	Futurepercentchange   float64 `json:"futurePercentChange"`
	Exchangename          string  `json:"exchangeName"`
	Securitystatus        string  `json:"securityStatus"`
	Openfloat64erest      float64 `json:"openfloat64erest"`
	Mark                  float64 `json:"mark"`
	Tick                  float64 `json:"tick"`
	Tickamount            float64 `json:"tickAmount"`
	Product               string  `json:"product"`
	Futurepriceformat     string  `json:"futurePriceFormat"`
	Futuretradinghours    string  `json:"futureTradingHours"`
	Futureistradable      bool    `json:"futureIsTradable"`
	Futuremultiplier      float64 `json:"futureMultiplier"`
	Futureisactive        bool    `json:"futureIsActive"`
	Futuresettlementprice float64 `json:"futureSettlementPrice"`
	Futureactivesymbol    string  `json:"futureActiveSymbol"`
	Futureexpirationdate  string  `json:"futureExpirationDate"`
}

//FutureOption quote
type FutureOption struct {
	Symbol                          string  `json:"symbol"`
	Bidpriceindouble                float64 `json:"bidPriceInDouble"`
	Askpriceindouble                float64 `json:"askPriceInDouble"`
	Lastpriceindouble               float64 `json:"lastPriceInDouble"`
	Highpriceindouble               float64 `json:"highPriceInDouble"`
	Lowpriceindouble                float64 `json:"lowPriceInDouble"`
	Closepriceindouble              float64 `json:"closePriceInDouble"`
	Description                     string  `json:"description"`
	Openpriceindouble               float64 `json:"openPriceInDouble"`
	Netchangeindouble               float64 `json:"netChangeInDouble"`
	Openfloat64erest                float64 `json:"openfloat64erest"`
	Exchangename                    string  `json:"exchangeName"`
	Securitystatus                  string  `json:"securityStatus"`
	Volatility                      float64 `json:"volatility"`
	Moneyfloat64rinsicvalueindouble float64 `json:"moneyfloat64rinsicValueInDouble"`
	Multiplierindouble              float64 `json:"multiplierInDouble"`
	Digits                          float64 `json:"digits"`
	Strikepriceindouble             float64 `json:"strikePriceInDouble"`
	Contracttype                    string  `json:"contractType"`
	Underlying                      string  `json:"underlying"`
	Timevalueindouble               float64 `json:"timeValueInDouble"`
	Deltaindouble                   float64 `json:"deltaInDouble"`
	Gammaindouble                   float64 `json:"gammaInDouble"`
	Thetaindouble                   float64 `json:"thetaInDouble"`
	Vegaindouble                    float64 `json:"vegaInDouble"`
	Rhoindouble                     float64 `json:"rhoInDouble"`
	Mark                            float64 `json:"mark"`
	Tick                            float64 `json:"tick"`
	Tickamount                      float64 `json:"tickAmount"`
	Futureistradable                bool    `json:"futureIsTradable"`
	Futuretradinghours              string  `json:"futureTradingHours"`
	Futurepercentchange             float64 `json:"futurePercentChange"`
	Futureisactive                  bool    `json:"futureIsActive"`
	Futureexpirationdate            float64 `json:"futureExpirationDate"`
	Expirationtype                  string  `json:"expirationType"`
	Exercisetype                    string  `json:"exerciseType"`
	Inthemoney                      bool    `json:"intheMoney"`
}

//Index quote
type Index struct {
	Symbol          string  `json:"symbol"`
	Description     string  `json:"description"`
	Lastprice       float64 `json:"lastPrice"`
	Openprice       float64 `json:"openPrice"`
	Highprice       float64 `json:"highPrice"`
	Lowprice        float64 `json:"lowPrice"`
	Closeprice      float64 `json:"closePrice"`
	Netchange       float64 `json:"netChange"`
	Totalvolume     float64 `json:"totalVolume"`
	Tradetimeinlong float64 `json:"tradeTimeInLong"`
	Exchange        string  `json:"exchange"`
	Exchangename    string  `json:"exchangeName"`
	Digits          float64 `json:"digits"`
	Five2Wkhigh     float64 `json:"52WkHigh"`
	Five2Wklow      float64 `json:"52WkLow"`
	Securitystatus  string  `json:"securityStatus"`
}

//Option quote
type Option struct {
	Symbol                  string  `json:"symbol"`
	Description             string  `json:"description"`
	Bidprice                float64 `json:"bidPrice"`
	Bidsize                 float64 `json:"bidSize"`
	Askprice                float64 `json:"askPrice"`
	Asksize                 float64 `json:"askSize"`
	Lastprice               float64 `json:"lastPrice"`
	Lastsize                float64 `json:"lastSize"`
	Openprice               float64 `json:"openPrice"`
	Highprice               float64 `json:"highPrice"`
	Lowprice                float64 `json:"lowPrice"`
	Closeprice              float64 `json:"closePrice"`
	Netchange               float64 `json:"netChange"`
	Totalvolume             float64 `json:"totalVolume"`
	Quotetimeinlong         float64 `json:"quoteTimeInLong"`
	Tradetimeinlong         float64 `json:"tradeTimeInLong"`
	Mark                    float64 `json:"mark"`
	Openfloat64erest        float64 `json:"openfloat64erest"`
	Volatility              float64 `json:"volatility"`
	Moneyfloat64rinsicvalue float64 `json:"moneyfloat64rinsicValue"`
	Multiplier              float64 `json:"multiplier"`
	Strikeprice             float64 `json:"strikePrice"`
	Contracttype            string  `json:"contractType"`
	Underlying              string  `json:"underlying"`
	Timevalue               float64 `json:"timeValue"`
	Deliverables            string  `json:"deliverables"`
	Delta                   float64 `json:"delta"`
	Gamma                   float64 `json:"gamma"`
	Theta                   float64 `json:"theta"`
	Vega                    float64 `json:"vega"`
	Rho                     float64 `json:"rho"`
	Securitystatus          string  `json:"securityStatus"`
	Theoreticaloptionvalue  float64 `json:"theoreticalOptionValue"`
	Underlyingprice         float64 `json:"underlyingPrice"`
	Uvexpirationtype        string  `json:"uvExpirationType"`
	Exchange                string  `json:"exchange"`
	Exchangename            string  `json:"exchangeName"`
	Settlementtype          string  `json:"settlementType"`
}

//Forex quote
type Forex struct {
	Symbol              string  `json:"symbol"`
	Bidpriceindouble    float64 `json:"bidPriceInDouble"`
	Askpriceindouble    float64 `json:"askPriceInDouble"`
	Lastpriceindouble   float64 `json:"lastPriceInDouble"`
	Highpriceindouble   float64 `json:"highPriceInDouble"`
	Lowpriceindouble    float64 `json:"lowPriceInDouble"`
	Closepriceindouble  float64 `json:"closePriceInDouble"`
	Exchange            string  `json:"exchange"`
	Description         string  `json:"description"`
	Openpriceindouble   float64 `json:"openPriceInDouble"`
	Changeindouble      float64 `json:"changeInDouble"`
	Percentchange       float64 `json:"percentChange"`
	Exchangename        string  `json:"exchangeName"`
	Digits              float64 `json:"digits"`
	Securitystatus      string  `json:"securityStatus"`
	Tick                float64 `json:"tick"`
	Tickamount          float64 `json:"tickAmount"`
	Product             string  `json:"product"`
	Tradinghours        string  `json:"tradingHours"`
	Istradable          bool    `json:"isTradable"`
	Marketmaker         string  `json:"marketMaker"`
	Five2Wkhighindouble float64 `json:"52WkHighInDouble"`
	Five2Wklowindouble  float64 `json:"52WkLowInDouble"`
	Mark                float64 `json:"mark"`
}

//ETF quote
type ETF struct {
	Symbol                       string  `json:"symbol"`
	Description                  string  `json:"description"`
	Bidprice                     float64 `json:"bidPrice"`
	Bidsize                      float64 `json:"bidSize"`
	Bidid                        string  `json:"bidId"`
	Askprice                     float64 `json:"askPrice"`
	Asksize                      float64 `json:"askSize"`
	Askid                        string  `json:"askId"`
	Lastprice                    float64 `json:"lastPrice"`
	Lastsize                     float64 `json:"lastSize"`
	Lastid                       string  `json:"lastId"`
	Openprice                    float64 `json:"openPrice"`
	Highprice                    float64 `json:"highPrice"`
	Lowprice                     float64 `json:"lowPrice"`
	Closeprice                   float64 `json:"closePrice"`
	Netchange                    float64 `json:"netChange"`
	Totalvolume                  float64 `json:"totalVolume"`
	Quotetimeinlong              float64 `json:"quoteTimeInLong"`
	Tradetimeinlong              float64 `json:"tradeTimeInLong"`
	Mark                         float64 `json:"mark"`
	Exchange                     string  `json:"exchange"`
	Exchangename                 string  `json:"exchangeName"`
	Marginable                   bool    `json:"marginable"`
	Shortable                    bool    `json:"shortable"`
	Volatility                   float64 `json:"volatility"`
	Digits                       float64 `json:"digits"`
	Five2Wkhigh                  float64 `json:"52WkHigh"`
	Five2Wklow                   float64 `json:"52WkLow"`
	Peratio                      float64 `json:"peRatio"`
	Divamount                    float64 `json:"divAmount"`
	Divyield                     float64 `json:"divYield"`
	Divdate                      string  `json:"divDate"`
	Securitystatus               string  `json:"securityStatus"`
	Regularmarketlastprice       float64 `json:"regularMarketLastPrice"`
	Regularmarketlastsize        float64 `json:"regularMarketLastSize"`
	Regularmarketnetchange       float64 `json:"regularMarketNetChange"`
	Regularmarkettradetimeinlong float64 `json:"regularMarketTradeTimeInLong"`
}

//Equity quote
type Equity struct {
	Symbol                       string  `json:"symbol"`
	Description                  string  `json:"description"`
	Bidprice                     float64 `json:"bidPrice"`
	Bidsize                      float64 `json:"bidSize"`
	Bidid                        string  `json:"bidId"`
	Askprice                     float64 `json:"askPrice"`
	Asksize                      float64 `json:"askSize"`
	Askid                        string  `json:"askId"`
	Lastprice                    float64 `json:"lastPrice"`
	Lastsize                     float64 `json:"lastSize"`
	Lastid                       string  `json:"lastId"`
	Openprice                    float64 `json:"openPrice"`
	Highprice                    float64 `json:"highPrice"`
	Lowprice                     float64 `json:"lowPrice"`
	Closeprice                   float64 `json:"closePrice"`
	Netchange                    float64 `json:"netChange"`
	Totalvolume                  float64 `json:"totalVolume"`
	Quotetimeinlong              float64 `json:"quoteTimeInLong"`
	Tradetimeinlong              float64 `json:"tradeTimeInLong"`
	Mark                         float64 `json:"mark"`
	Exchange                     string  `json:"exchange"`
	Exchangename                 string  `json:"exchangeName"`
	Marginable                   bool    `json:"marginable"`
	Shortable                    bool    `json:"shortable"`
	Volatility                   float64 `json:"volatility"`
	Digits                       float64 `json:"digits"`
	Five2Wkhigh                  float64 `json:"52WkHigh"`
	Five2Wklow                   float64 `json:"52WkLow"`
	Peratio                      float64 `json:"peRatio"`
	Divamount                    float64 `json:"divAmount"`
	Divyield                     float64 `json:"divYield"`
	Divdate                      string  `json:"divDate"`
	Securitystatus               string  `json:"securityStatus"`
	Regularmarketlastprice       float64 `json:"regularMarketLastPrice"`
	Regularmarketlastsize        float64 `json:"regularMarketLastSize"`
	Regularmarketnetchange       float64 `json:"regularMarketNetChange"`
	Regularmarkettradetimeinlong float64 `json:"regularMarketTradeTimeInLong"`
}
