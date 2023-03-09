package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	var data TemplateData

	data.StockList, _ = GetPortfolio()
	data.HistoricalValue, _ = GetPortfolioValue()

	render(w, "main.page.gohtml", data)
}

func Transactions(w http.ResponseWriter, r *http.Request) {
	var data TemplateData
	data.TransactionHistory, _ = GetTransactionHistory()

	render(w, "transactions.page.gohtml", data)
}
