package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From Go-Server running inside Docker Container!")
}

func main() {
	http.HandleFunc("/server-api", serverAPIHandler) // Update this line of code

	fmt.Printf("Starting Go-Server at port 8091 in Docker\n")

	if err := http.ListenAndServe("0.0.0.0:8091", nil); err != nil {
		log.Fatal(err)
	}
}
