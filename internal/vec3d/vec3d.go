package vec3d

import (
	"fmt"
	"math"
)

type Vec3d struct {
	X, Y, Z float64
}

func (v *Vec3d) Negate() Vec3d {
	return Vec3d{-v.X, -v.Y, -v.Z}
}

func (v *Vec3d) Add(v2 Vec3d) Vec3d {
	return Vec3d{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v *Vec3d) Substract(v2 Vec3d) Vec3d {
	return v.Add(v2.Negate())
}

func (v *Vec3d) Multiply(factor float64) Vec3d {
	return Vec3d{factor * v.X, factor * v.Y, factor * v.Z}
}

func (v *Vec3d) Divide(factor float64) Vec3d {
	return v.Multiply(1 / factor)
}

func (v *Vec3d) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3d) GetColor() string {
	return fmt.Sprintf("%d %d %d\n",
		int(255.99*v.X),
		int(255.99*v.Y),
		int(255.99*v.Z))
}
func DotProduct(v1 Vec3d, v2 Vec3d) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func crossProduct(v1 Vec3d, v2 Vec3d) Vec3d {
	return Vec3d{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y + v1.Y*v2.X}
}

func unitVector(v1 Vec3d) Vec3d {
	return v1.Divide(v1.Length())
}
