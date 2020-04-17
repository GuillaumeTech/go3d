package hit

import (
	"math"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Sphere struct {
	Center geom.Vec3d
	Radius float64
	Mat    Material
}

func (sphere Sphere) hit(ray geom.Ray, tMin float64, tMax float64, record *HitRecord) bool {
	oc := geom.SubstractTwoVec(ray.Origin, sphere.Center)
	a := geom.DotProduct(ray.Direction, ray.Direction)
	halfB := geom.DotProduct(oc, ray.Direction)
	c := geom.DotProduct(oc, oc) - sphere.Radius*sphere.Radius
	discriminant := halfB*halfB - a*c
	if discriminant > 0 {

		root := math.Sqrt(discriminant)
		temp := (-halfB - root) / a

		if temp < tMax && temp > tMin {
			record.T = temp
			record.P = ray.At(temp)
			outwardNormal := geom.DivideVec(geom.SubstractTwoVec(record.P, sphere.Center), sphere.Radius)
			record.setFaceNormal(ray, outwardNormal)
			record.Mat = sphere.Mat
			return true
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			record.T = temp
			record.P = ray.At(temp)
			outwardNormal := geom.DivideVec(geom.SubstractTwoVec(record.P, sphere.Center), sphere.Radius)
			record.setFaceNormal(ray, outwardNormal)
			record.Mat = sphere.Mat
			return true
		}
	}
	return false
}
