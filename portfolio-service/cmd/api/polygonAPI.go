package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type StockPriceHistory struct {
	Ticker       string   `json:"ticker"`
	QueryCount   int      `json:"queryCount"`
	ResultsCount int      `json:"resultsCount"`
	Adjusted     bool     `json:"adjusted"`
	Results      []Result `json:"results"`
	Status       string   `json:"status"`
	RequestId    string   `json:"request_id"`
}

type Result struct {
	Volume               float32 `json:"v"`
	VolumeWeighted       float32 `json:"vw"`
	Open                 float32 `json:"o"`
	Close                float32 `json:"c"`
	High                 float32 `json:"h"`
	Low                  float32 `json:"l"`
	Timestamp            int     `json:"t"`
	NumberOfTransactions int     `json:"n"`
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
