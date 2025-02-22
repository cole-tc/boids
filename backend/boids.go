package main

import (
	"math"
	"math/rand"
)

// TODO: dynamically scale this for each screen
const (
	MaxX        float64 = 2500
	MaxY        float64 = 1200
	SpeedLimit  float64 = 200
	MinDistance float64 = 20
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
	FlockSize float64

	SpeedLimit float64
}

// CreateFlock
// Initialize boids randomly within a 800x600 space
func CreateFlock(flockSize int) Flock {
	// init random pos + vel
	boids := make([]Boid, flockSize)
	for i := range boids {
		boids[i] = Boid{
			// 800x600 space
			X: rand.Float64() * MaxX,
			Y: rand.Float64() * MaxY,
			// rand gives random between 0.0 and 1.0
			VX: (rand.Float64() * 2) - 1,
			VY: (rand.Float64() * 2) - 1,
		}
	}

	flock := Flock{
		Boids:     boids,
		FlockSize: FlockSize,
	}

	return flock
}

// UpdateFlock
// Update boid positions (basic movement logic for now)
func UpdateFlock(flock Flock) Flock {
	// StackOverflow: "Structs are copied in range loops. You need to access by index."
	for i := range flock.Boids {
		offsetX, offsetY := flock.Boids[i].FlyTowardsCenter(flock)
		flock.Boids[i].VX += offsetX
		flock.Boids[i].VY += offsetY

		offsetX, offsetY = flock.Boids[i].FlyAwayFromOtherBoids(flock)
		flock.Boids[i].VX += offsetX
		flock.Boids[i].VY += offsetY

		offsetX, offsetY = flock.Boids[i].MatchBoidVelocity(flock)
		flock.Boids[i].VX += offsetX
		flock.Boids[i].VY += offsetY

		offsetX, offsetY = flock.Boids[i].StayWithinBounds(flock)
		flock.Boids[i].VX += offsetX
		flock.Boids[i].VY += offsetY

		// limit speed after calculations, before committing
		flock.Boids[i].LimitSpeed()

		flock.Boids[i].X += flock.Boids[i].VX
		flock.Boids[i].Y += flock.Boids[i].VY
	}
	return flock
}

// FlyTowardsCenter
// Rule 1:
// Boids try to fly towards the centre of mass of neighbouring boids.
func (boid *Boid) FlyTowardsCenter(flock Flock) (offsetX float64, offsetY float64) {
	pcX := 0.0
	pcY := 0.0

	for i := range flock.Boids {
		// "perceived" center, excluding oneself
		if *boid == flock.Boids[i] {
			continue
		}

		pcX += flock.Boids[i].X
		pcY += flock.Boids[i].Y
	}

	pcX = pcX / (FlockSize - 1)
	pcY = pcY / (FlockSize - 1)

	offsetX = (pcX - boid.X) / 200
	offsetY = (pcY - boid.Y) / 200
	return offsetX, offsetY
}

// FlyAwayFromOtherBoids
// Rule 2:
// Boids try to keep a small distance from other Boids
func (boid *Boid) FlyAwayFromOtherBoids(flock Flock) (offsetX float64, offsetY float64) {
	for i := range flock.Boids {
		if *boid == flock.Boids[i] {
			continue
		}

		// distance between the two boids
		distance := Distance(*boid, flock.Boids[i])

		if distance < MinDistance {
			offsetX -= flock.Boids[i].X - boid.X
			offsetY -= flock.Boids[i].Y - boid.Y
		}
	}
	return offsetX, offsetY
}

// MatchBoidVelocity
// Rule 3:
// Boids try to match velocity of neighbouring boids.
func (boid *Boid) MatchBoidVelocity(flock Flock) (offsetX float64, offsetY float64) {
	pvX := 0.0
	pvY := 0.0

	for i := range flock.Boids {
		// "perceived" center, excluding oneself
		if *boid == flock.Boids[i] {
			continue
		}

		pvX += flock.Boids[i].VX
		pvY += flock.Boids[i].VY
	}

	pvX = pvX / (FlockSize - 1)
	pvY = pvY / (FlockSize - 1)

	offsetX = (pvX - boid.VX) / 8
	offsetY = (pvY - boid.VY) / 8
	return offsetX, offsetY
}

// StayWithinBounds
// Softly coax the Boid back within the canvas bounds
func (boid *Boid) StayWithinBounds(flock Flock) (offsetX float64, offsetY float64) {
	if boid.X < 0 {
		offsetX = 10
	} else if boid.X > MaxX {
		offsetX = -10
	}

	if boid.Y < 0 {
		offsetY = 10
	} else if boid.Y > MaxY {
		offsetY = -10
	}
	return offsetX, offsetY
}

// LimitSpeed
// Limit the speed of the Boid
// This func should run last, as it directly modifies the Boid velocity
func (boid *Boid) LimitSpeed() {
	speed := math.Sqrt(
		math.Pow(boid.VX, 2) + math.Pow(boid.VY, 2))

	if speed > SpeedLimit {
		boid.VX = (boid.VX / speed) * SpeedLimit
		boid.VY = (boid.VY / speed) * SpeedLimit
	}
}

// Distance
// Euclidean distance between two x,y points
func Distance(boid1 Boid, boid2 Boid) float64 {
	return math.Sqrt(
		math.Pow(boid1.X-boid2.X, 2) +
			math.Pow(boid1.Y-boid2.Y, 2))
}
