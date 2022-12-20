package main

import (
	"fmt"
	"log"
	"net/http"
	"portfolio-service/data"
)

const webPort = "80"

type AppConfig struct {
	m_postgreSQL data.PostgreSQL
}

func main() {
	app := AppConfig{}

	postgreSQL, err := data.ConnectToPostgresSQL()
	if err != nil {
		log.Panic(err)
	}
	app.m_postgreSQL = postgreSQL

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
