package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/model"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	upCurrency := strings.ToUpper(currency)

	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		json := string(bodyBytes)

		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	rate := model.Rate{Currency: currency, Price: 20}
	return &rate, nil
}