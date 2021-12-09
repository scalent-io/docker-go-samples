package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func clientAPIHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n\nHello From Go-Client running inside Docker Container!")

	fmt.Fprintf(w, "\n\nCalling Go-server /server-api endpoint")
	response, err := http.Get("http://goserver.host:8091/server-api")
	if err != nil {
		fmt.Println("Error calling server api")
		fmt.Fprintf(w, "\n\nError while calling /server-api endpoint ", err.Error())
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		fmt.Fprintf(w, "\n\nError while reading response from /server-api endpoint", err.Error())
		return
	}

	fmt.Fprintf(w, "\n\nResponse from /server-api endpoint-------\n\n")
	respText := string(body)
	fmt.Fprintf(w, respText)
}

func main() {
	http.HandleFunc("/client-api", clientAPIHandler)
	fmt.Printf("Starting Go-Client at port 8090 in Docker\n")

	if err := http.ListenAndServe("0.0.0.0:8090", nil); err != nil {
		log.Fatal(err)
	}
}
