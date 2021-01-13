package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinMarketChartParams{
		VSCurrency: "nzd",
		Days:       5,
	}

	marketChart, err := coingecko.GetCoinMarketChart("stellar", params)

	if err == nil {
		fmt.Println(marketChart)
	} else {
		fmt.Println(err)
	}
}
