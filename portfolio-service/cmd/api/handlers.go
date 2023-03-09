package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"portfolio-service/data"
	"sort"
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
	portfolioTimestampToValueMap := make(map[int64]float32)

	var portfolio data.Portfolio
	portfolio, err := app.m_postgreSQL.GetPortfolioByUserId(1)
	if err != nil {
		return
	}

	transactions, err := app.m_postgreSQL.GetTransactionHostoryByUserId(1)
	if err != nil {
		return
	}
	log.Println(transactions)

	for _, position := range portfolio.Positions {
		stockPriceHistory := getStockPriceHistory(position.Symbol, time.Now().AddDate(-1, 0, 0), time.Now())
		for _, stockQuote := range stockPriceHistory.Results {
			val, ok := portfolioTimestampToValueMap[stockQuote.Timestamp]
			time := time.Unix(stockQuote.Timestamp/1000, 0)
			if ok {
				portfolioTimestampToValueMap[stockQuote.Timestamp] = val + stockQuote.Close*float32(GetQuantityByDate(time, position.Symbol, transactions))
			} else {
				portfolioTimestampToValueMap[stockQuote.Timestamp] = stockQuote.Close * float32(GetQuantityByDate(time, position.Symbol, transactions))
			}
		}
	}

	keys := make([]int64, 0, len(portfolioTimestampToValueMap))
	for key := range portfolioTimestampToValueMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		time := time.Unix(key/1000, 0)
		response.Quotes = append(response.Quotes, portfolioValueQuote{
			Date:  time.Format("2006-01-02"),
			Value: fmt.Sprintf("%f", portfolioTimestampToValueMap[key]),
		})
	}

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
