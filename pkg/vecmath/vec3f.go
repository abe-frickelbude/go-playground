package vecmath

import "math"

type Vec3f struct {
	X float32
	Y float32
	Z float32
}

func (vec *Vec3f) Length() float64 {
	return math.Sqrt(float64(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z))
}

func (vec *Vec3f) Scale(sc float32) {
	vec.X *= sc
	vec.Y *= sc
	vec.Z *= sc
}

func (vec *Vec3f) Scaled(sc float32) Vec3f {
	return Vec3f{vec.X * sc, vec.Y * sc, vec.Z * sc}
}

func (vec *Vec3f) Dot(vec2 *Vec3f) float32 {
	return vec.X*vec2.X + vec.Y*vec2.Y + vec.Z*vec2.Z
}
