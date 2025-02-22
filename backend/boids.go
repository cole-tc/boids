package main

import (
	"math/rand"
)

type Boid struct {

	// grid position
	X float64 `json:"x"`
	Y float64 `json:"y"`

	// velocities
	VX float64 `json:"vx"`
	VY float64 `json:"vy"`
}

type Flock struct {
	// the flock members
	Boids []Boid `json:"boids"`

	// float64 for less type conversion
	FlockSize float64 `json:"flockSize"`
}

// CreateFlock
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

	flock := Flock{
		Boids:     boids,
		FlockSize: float64(len(boids)),
	}

	return flock
}

// UpdateFlock
// Update boid positions (basic movement logic for now)
func UpdateFlock(flock Flock) Flock {
	// StackOverflow: "Structs are copied in range loops. You need to access by index."
	for i := range flock.Boids {
		Rule1X, Rule1Y := FlyTowardsCenter(flock, &flock.Boids[i])
		flock.Boids[i].X += Rule1X
		flock.Boids[i].Y += Rule1Y

		flock.Boids[i].X += flock.Boids[i].VX
		flock.Boids[i].Y += flock.Boids[i].VY

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

// FlyTowardsCenter
// Rule 1:
// Boids try to fly towards the centre of mass of neighbouring boids.
func FlyTowardsCenter(flock Flock, boid *Boid) (offsetX float64, offsetY float64) {
	PerceivedCenterX := 0.0
	PerceivedCenterY := 0.0

	for i := range flock.Boids {
		// "perceived" center, excluding oneself
		if *boid == flock.Boids[i] {
			continue
		}

		PerceivedCenterX += flock.Boids[i].X
		PerceivedCenterY += flock.Boids[i].Y
	}

	PerceivedCenterX /= flock.FlockSize - 1
	PerceivedCenterY /= flock.FlockSize - 1

	//fmt.Printf("FlyTowardsCenter -> [%v, %v]\n", PerceivedCenterX, PerceivedCenterY)

	offsetX = (PerceivedCenterX - boid.X) / 100
	offsetY = (PerceivedCenterY - boid.Y) / 100
	return
}
