package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Docker Container!")
}

func main() {
	http.HandleFunc("/hi", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
