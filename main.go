package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)
	router.HandleFunc("/movies", ListMovies)
	router.HandleFunc("/movies/{name}", ShowMovie)

	address := "localhost:8080"

	log.Fatal(http.ListenAndServe(address, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from GO server")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact page")
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Movie list")
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_name := params["name"]
	fmt.Fprintf(w, "Movie: %s", movie_name)
}
