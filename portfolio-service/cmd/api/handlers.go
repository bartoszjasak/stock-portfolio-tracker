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
	stockHistoricalDatasMap := make(map[string][]portfolioValueQuote)
	portfolioSymbols := []string{"AAPL", "MSFT"}

	for _, symbol := range portfolioSymbols {
		stockHistoricalDatasMap[symbol] = scrapeStockHistoricalData(symbol)
	}

	portfolioHistoricalValuation := make(map[string]string)
	for symbol, portfolioValueQuoteSlice := range stockHistoricalDatasMap {

	}

	response.Quotes = stockHistoricalDatasMap["AAPL"]

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
