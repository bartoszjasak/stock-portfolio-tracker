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

func render(w http.ResponseWriter, t string, data templateData) {

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

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
