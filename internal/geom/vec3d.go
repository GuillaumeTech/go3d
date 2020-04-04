package geom

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

func (v *Vec3d) Add(vectorB Vec3d) Vec3d {
	return Vec3d{v.X + vectorB.X, v.Y + vectorB.Y, v.Z + vectorB.Z}
}

func (v *Vec3d) Substract(vectorB Vec3d) Vec3d {
	return v.Add(vectorB.Negate())
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
func DotProduct(vectorA Vec3d, vectorB Vec3d) float64 {
	return vectorA.X*vectorB.X + vectorA.Y*vectorB.Y + vectorA.Z*vectorB.Z
}

func CrossProduct(vectorA Vec3d, vectorB Vec3d) Vec3d {
	return Vec3d{
		vectorA.Y*vectorB.Z - vectorA.Z*vectorB.Y,
		vectorA.Z*vectorB.X - vectorA.X*vectorB.Z,
		vectorA.X*vectorB.Y + vectorA.Y*vectorB.X}
}

func UnitVector(vectorA Vec3d) Vec3d {
	return vectorA.Divide(vectorA.Length())
}

func AddTwoVec(vectorA Vec3d, vectorB Vec3d) Vec3d {
	return vectorA.Add(vectorB)
}

func SubstractTwoVec(vectorA Vec3d, vectorB Vec3d) Vec3d {
	return vectorA.Substract(vectorB)
}

func MultiplyVec(factor float64, vector Vec3d) Vec3d {
	return vector.Multiply(factor)
}

func DivideVec(vector Vec3d, factor float64) Vec3d {
	return vector.Divide(factor)
}
