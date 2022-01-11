package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Song struct {
	Title  string `json:"name"`
	Lyrics string `json:"lyrics"`
	Artist string `json:"artist"`
}

type Songs struct {
	Number int    `json:"numSongs"`
	Songs  []Song `json:"songs"`
}

func main() {
	//home web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	//Lyrics json - GET
	fs := http.FileServer(http.Dir("./json"))
	http.Handle("/json/", http.StripPrefix("/json/", fs))

	//Add song
	http.HandleFunc("/songs/", handlerSongs)

	//Heroku ens diu a quin port
	port := os.Getenv("PORT")

	//Si no ens ho especifica triem nosaltres
	if port == "" {
		port = "80"
	}

	fmt.Println(port)
	_ = http.ListenAndServe(":"+port, nil)
}

func handlerSongs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./json/lyrics.json")
	case "POST":

		//Llegim la song de la peticio post
		var newSong Song
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newSong)
		if err != nil {
			panic(err)
		}

		//Llegim el fitxer json
		file, _ := ioutil.ReadFile("./json/lyrics.json")
		songsFile := Songs{}
		_ = json.Unmarshal([]byte(file), &songsFile)

		//Afegim la canço
		songsFile.Number++
		songsFile.Songs = append(songsFile.Songs, newSong)

		//Tornem a escriure el fitxer
		file, _ = json.MarshalIndent(songsFile, "", "    ")
		_ = ioutil.WriteFile("./json/lyrics.json", file, 0644)

		/*
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			name := r.FormValue("name")
			address := r.FormValue("address")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Address = %s\n", address)
		*/

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
