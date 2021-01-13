package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinParams{}
	coin, err := coingecko.GetCoin("bitcoin", params)

	if err == nil {
		fmt.Println(coin)
	} else {
		fmt.Println("Failed to get coin", err)
	}
}
