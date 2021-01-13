package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.ListCoinMarketsParams{VSCurrency: "nzd"}
	markets, err := coingecko.ListCoinMarkets(params)

	if err == nil {
		fmt.Println(markets)
	} else {
		fmt.Println("Failed to list markets", err)
	}
}
