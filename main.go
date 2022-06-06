package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)

	address := "localhost:8080"

	log.Fatal(http.ListenAndServe(address, router))
}

func Index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello world from GO server")
	}

func Contact(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Contact page")
	}