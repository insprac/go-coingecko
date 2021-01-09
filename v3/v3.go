package v3

func ListCoins() ([]ListedCoin, error) {
	var coins []ListedCoin
	err := get("/coins/list", &coins)
	return coins, err
}

func ListMarkets() ([]Market, error) {
	var markets []Market
	err := get("/coins/markets", &markets)
	return markets, err
}

func GetCoin(id string) (Coin, error) {
	var coin Coin
	err := get("/coins/"+id, &coin)
	return coin, err
}
