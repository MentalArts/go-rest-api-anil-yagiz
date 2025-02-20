package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /ping", handlePing)

	log.Println("Server listening...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved")
}