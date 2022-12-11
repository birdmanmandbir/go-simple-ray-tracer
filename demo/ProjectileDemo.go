package demo

import (
	"fmt"
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/canvas"
	"github.com/birdmanmandbir/go-simple-ray-tracer/internal/pkg/mat"
	"os"
	"path"
)

const OutputDir = "./output"

func ProjectileDemo() {
	velocityScale := 11.25
	velocity := mat.MultiplyByScalar(*mat.Normalize(*mat.NewVector(1, 1.8, 0)), velocityScale)
	p := newProjectile(mat.NewPoint(0, 1, 0), velocity)
	e := newEnvironment(mat.NewVector(0, -0.1, 0), mat.NewVector(-0.01, 0, 0))
	c := canvas.NewCanvas(900, 550)
	pointColor := mat.NewColor(1, 0.2, 0.1)
	for p.position.GetY() > 0 {
		p = tick(*e, *p)
		c.WritePixel(int(p.position.GetX()), c.H-int(p.position.GetY()), pointColor)
		fmt.Printf("The Object is at %v at height %v with velocity %v\n", *p.position, p.position.GetY(), *p.velocity)
	}
	println("Falling to ground!")
	_ = os.Mkdir(OutputDir, os.FileMode(0755))
	err := os.WriteFile(path.Join(OutputDir, "ProjectileDemoPicture.ppm"), []byte(c.ToPPM()), os.FileMode(0755))
	if err != nil {
		fmt.Printf("Write file error: %v", err)
	}
}

type Projectile struct {
	position *mat.Tuple4
	velocity *mat.Tuple4
}

type Environment struct {
	gravity *mat.Tuple4
	wind    *mat.Tuple4
}

func newEnvironment(gravity *mat.Tuple4, wind *mat.Tuple4) *Environment {
	return &Environment{gravity: gravity, wind: wind}
}

func newProjectile(position *mat.Tuple4, velocity *mat.Tuple4) *Projectile {
	return &Projectile{position: position, velocity: velocity}
}

func tick(env Environment, proj Projectile) *Projectile {
	newPosition := mat.Add(*proj.position, *proj.velocity)
	newVelocity := mat.Add(*mat.Add(*proj.velocity, *env.gravity), *env.wind)
	return newProjectile(newPosition, newVelocity)
}
