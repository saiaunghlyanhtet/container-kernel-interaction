package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current Time: %s", currentTime)
}

func main() {
	http.HandleFunc("/time", getTime)
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
