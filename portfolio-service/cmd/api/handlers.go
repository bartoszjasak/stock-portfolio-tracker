package main

import (
	"encoding/json"
	"net/http"
)

type getPortfolioValueResp struct {
	Quotes []portfolioValueQuote `json:"quote"`
}

type portfolioValueQuote struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func (app *AppConfig) GetPortfolioValue(w http.ResponseWriter, r *http.Request) {
	var response getPortfolioValueResp
	response.Quotes = append(response.Quotes, portfolioValueQuote{
		Date:  "11.11",
		Value: "3",
	})
	response.Quotes = append(response.Quotes, portfolioValueQuote{
		Date:  "12.11",
		Value: "6",
	})
	response.Quotes = append(response.Quotes, portfolioValueQuote{
		Date:  "13.11",
		Value: "3",
	})
	response.Quotes = append(response.Quotes, portfolioValueQuote{
		Date:  "14.11",
		Value: "9",
	})
	response.Quotes = append(response.Quotes, portfolioValueQuote{
		Date:  "15.11",
		Value: "5",
	})

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(responseJSON)
	if err != nil {
		return
	}
}
