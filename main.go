package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	fs := http.FileServer(http.Dir("./json"))
	http.Handle("/json/", http.StripPrefix("/json/", fs))

	//Heroku ens diu a quin port
	port := os.Getenv("PORT")
	//Si no ens ho especifica triem nosaltres
	if port == "" {
		port = "80"
	}
	fmt.Println(port)
	_ = http.ListenAndServe(":"+port, nil)
}
