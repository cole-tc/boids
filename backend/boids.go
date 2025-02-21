package main

import (
	"math/rand"
)

type Boid struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	VX float64 `json:"vx"`
	VY float64 `json:"vy"`
}

// Initialize boids randomly within a 800x600 space
func GenerateBoids(num int) []Boid {
	boids := make([]Boid, num)
	for i := range boids {
		boids[i] = Boid{
			X:  rand.Float64() * 800,
			Y:  rand.Float64() * 600,
			VX: (rand.Float64() * 2) - 1,
			VY: (rand.Float64() * 2) - 1,
		}
	}
	return boids
}

// Update boid positions (basic movement logic for now)
func UpdateBoids(boids []Boid) []Boid {
	for i := range boids {
		boids[i].X += boids[i].VX * 5
		boids[i].Y += boids[i].VY * 5

		// Keep boids inside bounds (wrap around)
		if boids[i].X < 0 {
			boids[i].X = 800
		} else if boids[i].X > 800 {
			boids[i].X = 0
		}
		if boids[i].Y < 0 {
			boids[i].Y = 600
		} else if boids[i].Y > 600 {
			boids[i].Y = 0
		}
	}
	return boids
}
