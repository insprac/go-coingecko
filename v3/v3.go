package v3

func ListCoins() ([]ListedCoin, error) {
	var coins []ListedCoin
	err := get("/coins/list", &coins)
	return coins, err
}

func ListCoinMarkets(params ListCoinMarketsParams) ([]CoinMarket, error) {
	var markets []CoinMarket
	err := getWithParams("/coins/markets", params, &markets)
	return markets, err
}

func GetCoin(id string, params GetCoinParams) (Coin, error) {
	var coin Coin
	err := getWithParams("/coins/"+id, params, &coin)
	return coin, err
}

func ListCoinTickers(
	id string,
	params ListCoinTickersParams,
) (CoinTickers, error) {
	var tickers CoinTickers
	err := getWithParams("/coins/"+id+"/tickers", params, &tickers)
	return tickers, err
}
