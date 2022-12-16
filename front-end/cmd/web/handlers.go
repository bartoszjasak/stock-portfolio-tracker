package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	var data templateData
	// data.HistoricalValue.Dates = []string{"1.12", "2.12", "3.12", "4.12", "5.12"}
	// data.HistoricalValue.Values = []string{"1", "2", "2", "8", "5"}

	// data.StockList = append(data.StockList, stock{
	// 	Name:     "Apple Inc.",
	// 	Symbol:   "AAPL",
	// 	Quantity: "40",
	// 	Price:    "143.22",
	// 	Value:    "5393.84",
	// })
	// data.StockList = append(data.StockList, stock{
	// 	Name:     "Microsoft Inc.",
	// 	Symbol:   "MSFT",
	// 	Quantity: "25",
	// 	Price:    "250",
	// 	Value:    "5393.84",
	// })

	data.HistoricalValue, _ = GetPortfolioValue()

	render(w, "test.page.gohtml", data)
}
