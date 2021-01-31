package main

import (
	"fmt"
	"log"
	"net/http"
	"routes"
)

func handleRequests() {
	fmt.Println("Init")
	log.Fatal(http.ListenAndServe(":5001", routes.SetRequest()))
}

func main() {
	handleRequests()
}
