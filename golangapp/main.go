package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!\n")
		fmt.Println("Request received!")
	})

	fmt.Println("Server started! Waiting requests...")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}