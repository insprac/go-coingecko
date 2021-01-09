package main

import (
	"encoding/json"
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	coin, err := coingecko.GetCoin("bitcoin")

	if err != nil {
		fmt.Println("Failed to get coin", err)
		return
	}

	data, err := json.Marshal(coin)

	if err != nil {
		fmt.Println("Failed to encode JSON", err)
		return
	}

	fmt.Println(string(data))
}
