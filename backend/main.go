package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var boids = GenerateBoids(50)

func getBoidsHandler(w http.ResponseWriter, r *http.Request) {
	boids = UpdateBoids(boids)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(boids)
	if err != nil {
		return
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.HandleFunc("/boids", getBoidsHandler)

	log.Println("Server started on :8080")
	go func() {
		for {
			time.Sleep(100 * time.Millisecond) // Update boids every 100ms
			boids = UpdateBoids(boids)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
