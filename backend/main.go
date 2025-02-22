package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	FlockSize = 300
)

var flock = CreateFlock(FlockSize)

func getBoidsHandler(w http.ResponseWriter, r *http.Request) {
	// update the math
	flock = UpdateFlock(flock)

	// send the response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(flock)

	if err != nil {
		return
	}
}

func main() {
	// init front-end
	http.Handle("/", http.FileServer(http.Dir("../frontend")))

	// handler func at /boids
	http.HandleFunc("/boids", getBoidsHandler)

	log.Println("Server started on :8080")
	go func() {
		for {
			time.Sleep(50 * time.Millisecond) // Update boids every 100ms
			flock = UpdateFlock(flock)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
