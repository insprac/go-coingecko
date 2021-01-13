package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinOHLCParams{
		VSCurrency: "nzd",
		Days:       7,
	}

	ohlc, err := coingecko.GetCoinOHLC("ethereum", params)

	if err == nil {
		fmt.Println(ohlc)
	} else {
		fmt.Println(err)
	}
}
