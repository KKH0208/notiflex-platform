package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

const version = "v0.1.1"

var counter atomic.Int64

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	pod, _ := os.Hostname()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"version": version,
		"pod":     pod,
	})
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	id := counter.Add(1)
	pod, _ := os.Hostname()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"id":  id,
		"pod": pod,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /id", idHandler)
	mux.HandleFunc("GET /version", versionHandler)

	log.Println("notiflex-api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
