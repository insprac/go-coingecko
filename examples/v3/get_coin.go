package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	coin, err := coingecko.GetCoin("bitcoin")

	if err == nil {
		fmt.Println(coin)
	} else {
		fmt.Println("Failed to get coin", err)
	}
}
