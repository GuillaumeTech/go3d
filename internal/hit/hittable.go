package hit

import "github.com/GuillaumeTech/3dgo/internal/geom"

type HitRecord struct {
	P, Normal geom.Vec3d
	T         float64
	FrontFace bool
}

func (h *HitRecord) setFaceNormal(ray geom.Ray, outwardNormal geom.Vec3d) {
	h.FrontFace = geom.DotProduct(ray.Direction, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}

}

type Hittable interface {
	hit(ray geom.Ray, tMin float64, tMax float64, record *HitRecord) bool
}
