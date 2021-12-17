package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./json"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	_ = http.ListenAndServe(":3000", nil)
}
