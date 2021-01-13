package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinTickersParams{
		IncludeExchangeLogo: true,
		Order:               "volume_desc",
	}

	tickers, err := coingecko.GetCoinTickers("ethereum", params)

	if err == nil {
		fmt.Println(tickers)
	} else {
		fmt.Println("Failed to list coin tickers", err)
	}
}
