package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.ListMarketsParams{VSCurrency: "nzd"}
	markets, err := coingecko.ListMarkets(params)

	if err == nil {
		fmt.Println(markets)
	} else {
		fmt.Println("Failed to list markets", err)
	}
}
