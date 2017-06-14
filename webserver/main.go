package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! ~10Pearls")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
