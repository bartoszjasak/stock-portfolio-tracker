package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	var data templateData

	data.StockList, _ = GetPortfolio()
	data.HistoricalValue, _ = GetPortfolioValue()

	render(w, "test.page.gohtml", data)
}
