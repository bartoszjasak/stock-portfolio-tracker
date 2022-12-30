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

func GetPortfolioValue() (historicalValue, error) {
	var ret historicalValue

	request, err := http.NewRequest("GET", "http://localhost:8080/get-portfolio-value", nil)
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
	var stockList []Position

	request, err := http.NewRequest("GET", "http://localhost:8080/get-portfolio", nil)
	if err != nil {
		return stockList, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return stockList, err
	}
	defer response.Body.Close()

	var portfolio getPortfolioResp
	json.NewDecoder(response.Body).Decode(&portfolio)

	return portfolio.Positions, nil
}
