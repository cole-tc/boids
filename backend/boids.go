package main

import (
	"fmt"
	"math/rand"
)

type Boid struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	VX float64 `json:"vx"`
	VY float64 `json:"vy"`
}

type Flock struct {
	// the flock members
	Boids []Boid `json:"boids"`
}

// Initialize boids randomly within a 800x600 space
func CreateFlock(flockSize int) Flock {
	// init random pos + vel
	boids := make([]Boid, flockSize)
	for i := range boids {
		boids[i] = Boid{
			// 800x600 space
			X: rand.Float64() * 800,
			Y: rand.Float64() * 600,
			// rand gives random between 0.0 and 1.0
			VX: (rand.Float64() * 2) - 1,
			VY: (rand.Float64() * 2) - 1,
		}
	}

	flock := Flock{Boids: boids}

	return flock
}

// Update boid positions (basic movement logic for now)
func UpdateFlock(flock Flock) Flock {
	fmt.Println("updatingFlock...")

	// StackOverflow: "Structs are copied in range loops. You need to access by index."
	for i := range flock.Boids {
		flock.Boids[i].X += flock.Boids[i].VX * 5
		flock.Boids[i].Y += flock.Boids[i].VY * 5

		// Keep flock.Boids[i] inside bounds (wrap around)
		if flock.Boids[i].X < 0 {
			flock.Boids[i].X = 800
		} else if flock.Boids[i].X > 800 {
			flock.Boids[i].X = 0
		}
		if flock.Boids[i].Y < 0 {
			flock.Boids[i].Y = 600
		} else if flock.Boids[i].Y > 600 {
			flock.Boids[i].Y = 0
		}
	}
	return flock
}
