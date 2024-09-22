package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/", handleRoot).Methods("GET")
	r.HandleFunc("/hello", handleHello).Methods("GET")
	r.HandleFunc("/health", handleHealth).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Welcome to the secure microservice!",
		Time:    time.Now().Format(time.RFC3339),
	}
	sendJSONResponse(w, response)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, secure world!",
		Time:    time.Now().Format(time.RFC3339),
	}
	sendJSONResponse(w, response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Healthy",
		Time:    time.Now().Format(time.RFC3339),
	}
	sendJSONResponse(w, response)
}

func sendJSONResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
