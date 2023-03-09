package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type getPortfolioValueResp struct {
	Quotes []portfolioValueQuote `json:"quote"`
}

type portfolioValueQuote struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

type getPortfolioResp struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	StockName string  `json:"stock_name"`
	Symbol    string  `json:"symbol"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
	Value     float32 `json:"value"`
}

type GetTransactionHistoryResp struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	StockName string    `json:"stock_name"`
	Symbol    string    `json:"symbol"`
	Price     float32   `json:"price"`
	Quantity  int       `json:"quantity"`
	Date      time.Time `json:"date"`
	UserId    int       `json:"user_id"`
}

func GetPortfolioValue() (historicalValue, error) {
	var ret historicalValue

	request, err := http.NewRequest("GET", "http://localhost:8080/portfolio-value", nil)
	if err != nil {
		return ret, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ret, err
	}
	defer response.Body.Close()

	var portfolioValue getPortfolioValueResp
	json.NewDecoder(response.Body).Decode(&portfolioValue)

	for _, quote := range portfolioValue.Quotes {
		ret.Dates = append(ret.Dates, quote.Date)
		ret.Values = append(ret.Values, quote.Value)
	}

	return ret, nil
}

func GetPortfolio() ([]Position, error) {
	var portfolio getPortfolioResp

	request, err := http.NewRequest("GET", "http://localhost:8080/portfolio", nil)
	if err != nil {
		return portfolio.Positions, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return portfolio.Positions, err
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&portfolio)

	return portfolio.Positions, nil
}

func GetTransactionHistory() ([]Transaction, error) {
	var getTransactionHistoryResp GetTransactionHistoryResp

	resp, err := http.Get("http://localhost:8080/transaction-history")
	if err != nil {
		return getTransactionHistoryResp.Transactions, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&getTransactionHistoryResp)

	return getTransactionHistoryResp.Transactions, nil
}
