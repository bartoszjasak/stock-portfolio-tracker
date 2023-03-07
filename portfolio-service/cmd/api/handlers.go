package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"portfolio-service/data"
	"time"
)

type getPortfolioValueResp struct {
	Quotes []portfolioValueQuote `json:"quote"`
}

type portfolioValueQuote struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func (app *AppConfig) GetPortfolioValue(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetPortfolioValue")
	var response getPortfolioValueResp
	stockHistoricalDatasMap := make(map[string][]portfolioValueQuote)
	portfolioSymbols := []string{"AAPL", "MSFT"}

	for _, symbol := range portfolioSymbols {
		history := getStockPriceHistory(symbol, time.Now().AddDate(0, -1, 0), time.Now())
		stockHistoricalDatasMap[symbol] = scrapeStockHistoricalData(symbol)
	}

	// portfolioHistoricalValuation := make(map[string]string)
	// for symbol, quoteSlice := range stockHistoricalDatasMap {
	// 	for _, quote
	// 	portfolioHistoricalValuation[portfolioValueQuoteSlice] =
	// }

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

func (app *AppConfig) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetPortfolio")
	var response data.Portfolio

	response, err := app.m_postgreSQL.GetPortfolioByUserId(1)
	if err != nil {
		return
	}

	fmt.Println(response)

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
