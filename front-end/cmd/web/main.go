package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8081"

type historicalValue struct {
	Dates  []string
	Values []string
}

type TemplateData struct {
	StockList          []Position
	HistoricalValue    historicalValue
	TransactionHistory []Transaction
}

type Config struct {
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

func render(w http.ResponseWriter, t string, data TemplateData) {

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

	tmpl := template.New("")

	tmpl = tmpl.Funcs(template.FuncMap{"mul": Mul})

	tmpl, err := tmpl.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, t, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Mul(param1 int, param2 float32) float32 {
	return (float32)(param1) * param2
}
