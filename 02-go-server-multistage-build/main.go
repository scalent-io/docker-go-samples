package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Multi-Stage Docker Container!")
}

func main() {
	http.HandleFunc("/hi", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 80 in Docker\n")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
