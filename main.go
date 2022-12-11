package main

import (
	"fmt"
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"
)

func main() {
	p := NewProjectile(mat.NewPoint(0, 1, 0), mat.Normalize(*mat.NewVector(1, 5, 0)))
	e := NewEnvironment(mat.NewVector(0, -0.1, 0), mat.NewVector(-0.01, 0, 0))
	for {
		p = Tick(*e, *p)
		fmt.Printf("The Object is at %v at height %v with velocity %v\n", *p.position, p.position.GetY(), *p.velocity)
		if p.position.GetY() <= 0 {
			break
		}
	}
	println("Falling to ground!")
}

type Projectile struct {
	position *mat.Tuple4
	velocity *mat.Tuple4
}

type Environment struct {
	gravity *mat.Tuple4
	wind    *mat.Tuple4
}

func NewEnvironment(gravity *mat.Tuple4, wind *mat.Tuple4) *Environment {
	return &Environment{gravity: gravity, wind: wind}
}

func NewProjectile(position *mat.Tuple4, velocity *mat.Tuple4) *Projectile {
	return &Projectile{position: position, velocity: velocity}
}

func Tick(env Environment, proj Projectile) *Projectile {
	newPosition := mat.Add(*proj.position, *proj.velocity)
	newVelocity := mat.Add(*mat.Add(*proj.velocity, *env.gravity), *env.wind)
	return NewProjectile(newPosition, newVelocity)
}
