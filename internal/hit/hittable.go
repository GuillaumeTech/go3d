package hit

import (
	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Hittable interface {
	hit(ray geom.Ray, tMin float64, tMax float64, record *HitRecord) bool
}

type HitRecord struct {
	P, Normal geom.Vec3d
	T         float64
	FrontFace bool
	Mat       Material
}

func (h *HitRecord) setFaceNormal(ray geom.Ray, outwardNormal geom.Vec3d) {
	h.FrontFace = geom.DotProduct(ray.Direction, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}
}

type HittableList struct {
	ObjectList []Hittable
}

func (hl *HittableList) Add(obj Hittable) {
	hl.ObjectList = append(hl.ObjectList, obj)
}

func (hl *HittableList) Hit(ray geom.Ray, tMin float64, tMax float64, record *HitRecord) bool {
	var tempRecord HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, obj := range hl.ObjectList {
		if obj.hit(ray, tMin, closestSoFar, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			*record = tempRecord
		}
	}

	return hitAnything
}
