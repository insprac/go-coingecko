package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	markets, err := coingecko.ListMarkets()

	if err == nil {
		fmt.Println(markets)
	} else {
		fmt.Println("Failed to list markets", err)
	}
}
