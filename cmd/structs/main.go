package main

import (
	// local package import
	"de.frickelbude/go_playground/pkg/vecmath"

	// example of putting all exported identifiers into the current file, removing
	// the need for explicit package name "qualifier" in statements
	. "fmt"
)

func main() {
	// func main() {
	// fmt.Println(Vec3f{0.5, 0.7, 0.6})

	vec1 := &vecmath.Vec3f{1.0, 0.0, 0.0}
	vec2 := &vecmath.Vec3f{1.0, 1.0, 1.0}

	Println(vec1.Dot(vec2))
	vec1.Scale(2.0)

	vectors := []vecmath.Vec3f{
		{1.0, 0.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0},
	}

	for _, item := range vectors[:2] {
		Println(item)
	}

	// slice
	model := make([]vecmath.Vec3f, 4)
	for i := range model {
		model[i] = vecmath.Vec3f{X: 1.0, Y: 1.0, Z: 1.0}
		model[i].Scale(float32(i))
		Println(model[i].Dot(&model[1]))
	}
	Println(model)
	Println(vec1.Scaled(15.0))

	// initializing multiple variables in a single line
	x0, y0, z0 := vec1.X, vec1.Y, vec1.Z
	Printf("%.3f %.3f %.3f\n", x0, y0, z0)

	Printf("%+v -> %p", vec1, vec1) // ?!
}
