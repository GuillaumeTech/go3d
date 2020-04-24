package hit

import (
	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Triangle struct {
	a, b, c, outwardNormal geom.Vec3d
	Mat                    Material
}

//Möller–Trumbore
func (triangle Triangle) hit(ray geom.Ray, tMin float64, tMax float64, record *HitRecord) bool {

	EPSILON := 0.00001
	edge1 := geom.SubstractTwoVec(triangle.b, triangle.a)
	edge2 := geom.SubstractTwoVec(triangle.c, triangle.a)

	h := geom.CrossProduct(ray.Direction, edge2)
	a := geom.DotProduct(edge1, h)

	if a > -EPSILON && a < EPSILON {
		return false
	}

	f := 1 / a
	s := geom.SubstractTwoVec(ray.Origin, triangle.a)
	u := f * geom.DotProduct(s, h)
	if u < 0 || u > 1 {
		return false
	}

	q := geom.CrossProduct(s, edge1)
	v := f * geom.DotProduct(ray.Direction, q)
	if v < 0 || u+v > 1 {
		return false
	}
	t := f * geom.DotProduct(edge2, q)
	if t > EPSILON {
		record.T = t
		record.P = ray.At(t)
		record.setFaceNormal(ray, triangle.outwardNormal)
		record.Mat = triangle.Mat
		return true
	}

	return false
}

func NewTriangle(a, b, c geom.Vec3d, mat Material) Triangle {
	edge1 := geom.SubstractTwoVec(b, a)
	edge2 := geom.SubstractTwoVec(c, a)
	normal := geom.UnitVector(geom.CrossProduct(edge1, edge2))
	return Triangle{a, b, c, normal, mat}
}
