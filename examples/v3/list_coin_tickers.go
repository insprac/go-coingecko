package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.ListCoinTickersParams{
		IncludeExchangeLogo: true,
		Order:               "volume_desc",
	}

	tickers, err := coingecko.ListCoinTickers("ethereum", params)

	if err == nil {
		fmt.Println(tickers)
	} else {
		fmt.Println("Failed to list coin tickers", err)
	}
}
