package main

import (
	"fmt"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate for %v is %.2f", rate.Currency, rate.Price)
	}
}