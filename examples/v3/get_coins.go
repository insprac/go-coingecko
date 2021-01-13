package main

import (
	"fmt"

	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	coins, err := coingecko.GetCoins()

	if err == nil {
		fmt.Println(coins)
	} else {
		fmt.Println("Failed to list coins", err)
	}
}
