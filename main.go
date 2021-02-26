package main

import (
	"log"
	"net/http"

	"github.com/angelicahill/shakespearescanner/shakespeare"
)

func main() {

	http.HandleFunc("/appendix/", shakespeare.AppendixHandler)
	http.HandleFunc("/run2", shakespeare.Run2)

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
