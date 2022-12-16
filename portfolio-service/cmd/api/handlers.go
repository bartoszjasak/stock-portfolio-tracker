package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
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
	response.Quotes = scrapeStockHistoricalData()

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

func scrapeStockHistoricalData() []portfolioValueQuote {
	var ret []portfolioValueQuote
	scrapeUrl := "https://finance.yahoo.com/quote/AAPL/history?p=AAPL"

	collector := colly.NewCollector(colly.AllowedDomains("finance.yahoo.com", "www.finance.yahoo.com"))

	collector.OnHTML("tr", func(h *colly.HTMLElement) {
		selection := h.DOM
		childNodes := selection.Children().Nodes
		if len(childNodes) == 7 {
			date := selection.FindNodes(childNodes[0]).Text()
			value := selection.FindNodes(childNodes[5]).Text()
			fmt.Printf("Date: %s, value %s\n", date, value)
			if date != "Date" {
				ret = append([]portfolioValueQuote{portfolioValueQuote{
					Date:  date,
					Value: value,
				}}, ret...)
			}
		}
	})

	collector.Visit(scrapeUrl)
	return ret
}
