package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ConsumedLicenses struct {
    TotalSeatsConsumed int `json:"total_seats_consumed"`
    TotalSeatsPurchased int `json:"total_seats_purchased"`
}

func GetConsumedLicenses(enterprise string, token string) *ConsumedLicenses {

		requestURL := fmt.Sprintf("https://api.github.com/enterprises/%s/consumed-licenses", enterprise)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)

		if err != nil {
			fmt.Println("Error creating request")
			os.Exit(1)
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res, err := http.DefaultClient.Do(req)
		
		if err != nil {
			fmt.Println("Error sending request", err)
			os.Exit(1)
		}

		defer res.Body.Close()

		resBody, _ := ioutil.ReadAll(res.Body)
		response := string(resBody)

		var consumedLicenses ConsumedLicenses
		json.Unmarshal([]byte(response), &consumedLicenses)
		
		return &consumedLicenses
}