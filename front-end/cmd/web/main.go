package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8081"

type stock struct {
	Name     string
	Symbol   string
	Quantity string
	Price    string
	Value    string
}

type historicalValue struct {
	Date  []string
	Value []string
}

type templateData struct {
	StockList       []stock
	HistoricalValue historicalValue
}

func main() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	fmt.Println("Starting front end service on port 8081")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data templateData
	data.HistoricalValue.Date = []string{"1.12", "2.12", "3.12", "4.12", "5.12"}
	data.HistoricalValue.Value = []string{"1", "2", "2", "8", "5"}

	data.StockList = append(data.StockList, stock{
		Name:     "Apple Inc.",
		Symbol:   "AAPL",
		Quantity: "40",
		Price:    "143.22",
		Value:    "5393.84",
	})
	data.StockList = append(data.StockList, stock{
		Name:     "Microsoft Inc.",
		Symbol:   "MSFT",
		Quantity: "25",
		Price:    "250",
		Value:    "5393.84",
	})

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
