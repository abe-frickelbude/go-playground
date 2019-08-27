package vecmath

import "math"

type Vec4f struct {
	X, Y, Z, W float32
}

func (vec *Vec4f) Length() float64 {
	return math.Sqrt(float64(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z + vec.W*vec.W))
}
