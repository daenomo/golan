package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := HelloResponse{Message: "Hello, Go HTTP Server!"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func startBackgroundLogger() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		log.Printf("[goroutine] heartbeat at %v", t)
	}
}

func main() {
	// ✅ バックグラウンドでログを流す goroutine
	go startBackgroundLogger()

	http.HandleFunc("/hello", helloHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
