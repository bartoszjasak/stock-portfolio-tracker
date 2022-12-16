package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type AppConfig struct {
}

func main() {
	app := AppConfig{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
