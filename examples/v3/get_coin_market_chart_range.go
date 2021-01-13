package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinMarketChartRangeParams{
		VSCurrency: "nzd",
		From:       1609412400,
		To:         1609498800,
	}

	marketChart, err := coingecko.GetCoinMarketChartRange("polkadot", params)

	if err == nil {
		fmt.Println(marketChart)
	} else {
		fmt.Println(err)
	}
}
