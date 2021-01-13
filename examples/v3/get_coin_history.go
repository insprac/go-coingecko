package main

import (
	"fmt"
	coingecko "github.com/insprac/go-coingecko/v3"
)

func main() {
	params := coingecko.GetCoinHistoryParams{
		Date:         "01-01-2021",
		Localization: true,
	}

	history, err := coingecko.GetCoinHistory("cardano", params)

	if err == nil {
		fmt.Println(history)
	} else {
		fmt.Println("Failed to list coin history", err)
	}
}
