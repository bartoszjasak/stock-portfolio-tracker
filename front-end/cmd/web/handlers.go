package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "test.page.gohtml")
}
