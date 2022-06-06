package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde server GO")
	})

	address := "localhost:8080"

	log.Fatal(http.ListenAndServe(address, nil))
}