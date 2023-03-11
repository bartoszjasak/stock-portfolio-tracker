package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type StockPriceHistory struct {
	Ticker       string             `json:"ticker"`
	QueryCount   int                `json:"queryCount"`
	ResultsCount int                `json:"resultsCount"`
	Adjusted     bool               `json:"adjusted"`
	Results      []StockPriceResult `json:"results"`
	Status       string             `json:"status"`
	RequestId    string             `json:"request_id"`
}

type StockPriceResult struct {
	Volume               float32 `json:"v"`
	VolumeWeighted       float32 `json:"vw"`
	Open                 float32 `json:"o"`
	Close                float32 `json:"c"`
	High                 float32 `json:"h"`
	Low                  float32 `json:"l"`
	Timestamp            int64   `json:"t"`
	NumberOfTransactions int     `json:"n"`
}

type StockDividendHistory struct {
	Status  string                `json:"status"`
	Results []StockDividendResult `json:"results"`
}

type StockDividendResult struct {
	CashAmount      float32 `json:"cash_amount"`
	Currency        string  `json:"currency"`
	DeclarationDate string  `json:"declaration_date"`
	DividenType     string  `json:"dividend_type"`
	ExDividendDate  string  `json:"ex_dividend_date"`
	Frequency       int     `json:"frequency"`
	PayDate         string  `json:"pay_date"`
	RecordDate      string  `json:"record_date"`
	Ticker          string  `json:"ticker"`
}

func getStockPriceHistory(symbol string, dateFrom time.Time, dateTo time.Time) StockPriceHistory {
	log.Printf("getStockPriceHistory")
	var ret StockPriceHistory

	resp, err := http.Get(fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/%s/range/1/day/%s/%s?adjusted=true&sort=desc&limit=720&apiKey=0dtTmunAmp5FDWKfpzL4YxKRbSpReuqW",
		symbol, dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02")))
	if err != nil {
		log.Fatal(err)
		return ret
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret
}

func getStockDividendHistory(symbol string, dateFrom time.Time, dateTo time.Time) StockDividendHistory {
	log.Println("getStockDividendHistory")
	var ret StockDividendHistory

	resp, err := http.Get(fmt.Sprintf("https://api.polygon.io/v3/reference/dividends?ticker=%s&ex_dividend_date.gt=%s&ex_dividend_date.lt=%s&apiKey=LOT8HqEp0sTHVRQflrDnKDaKzzUg1dwm",
		symbol, dateFrom.Format("2006-01-02"), dateTo.Format("2006-01-02")))
	if err != nil {
		log.Fatal(err)
		return ret
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret
}
