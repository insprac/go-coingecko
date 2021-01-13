package v3

func GetCoins() ([]ListedCoin, error) {
	var coins []ListedCoin
	err := get("/coins/list", &coins)
	return coins, err
}

func GetCoinMarkets(params GetCoinMarketsParams) ([]CoinMarket, error) {
	var markets []CoinMarket
	err := getWithParams("/coins/markets", params, &markets)
	return markets, err
}

func GetCoin(id string, params GetCoinParams) (Coin, error) {
	var coin Coin
	err := getWithParams("/coins/"+id, params, &coin)
	return coin, err
}

func GetCoinTickers(id string, params GetCoinTickersParams) (CoinTickers, error) {
	var tickers CoinTickers
	err := getWithParams("/coins/"+id+"/tickers", params, &tickers)
	return tickers, err
}

func GetCoinHistory(id string, params GetCoinHistoryParams) (CoinHistory, error) {
	var history CoinHistory
	err := getWithParams("/coins/"+id+"/history", params, &history)
	return history, err
}

func GetCoinMarketChart(id string, params GetCoinMarketChartParams) (CoinMarketChart, error) {
	var marketChart CoinMarketChart
	err := getWithParams("/coins/"+id+"/market_chart", params, &marketChart)
	return marketChart, err
}

func GetCoinMarketChartRange(id string, params GetCoinMarketChartRangeParams) (CoinMarketChart, error) {
	var marketChart CoinMarketChart
	err := getWithParams("/coins/"+id+"/market_chart/range", params, &marketChart)
	return marketChart, err
}

func GetCoinStatusUpdates(id string, params GetCoinStatusUpdatesParams) (CoinStatusUpdates, error) {
	var statusUpdates CoinStatusUpdates
	err := getWithParams("/coins/"+id+"/status_updates", params, &statusUpdates)
	return statusUpdates, err
}
