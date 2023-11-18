package api

import (
	"fmt"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/model"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (model.Rate, error) {
	upCurrency := strings.ToUpper(currency)

	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
}