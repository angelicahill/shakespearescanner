package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/angelicahill/shakespearescanner/shakespeare"
)

func main() {

	http.HandleFunc("/appendix/", shakespeare.AppendixHandler)
	http.HandleFunc("/run2", shakespeare.Run2)

	fmt.Println("listening on http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
