package main

import (
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./json"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Heroku ens diu a quin port
	port := os.Getenv("PORT")
	//Si no ens ho especifica triem nosaltres
	if port == "" {
		port = "8000"
	}

	_ = http.ListenAndServe(":"+port, nil)
}
