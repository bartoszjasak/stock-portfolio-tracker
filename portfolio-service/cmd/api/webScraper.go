package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

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
