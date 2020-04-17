package hit

import (
	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Lambertian struct {
	Albedo geom.Vec3d
}

func (l Lambertian) Scatter(rayIn geom.Ray, hitRecord *HitRecord, attenutaion *geom.Vec3d, scattered *geom.Ray) bool {
	scatterDir := geom.AddTwoVec(geom.RandomUnitVector(), geom.AddTwoVec(hitRecord.P, hitRecord.Normal))
	scattered.Direction = scatterDir
	scattered.Origin = hitRecord.P
	*attenutaion = l.Albedo
	return true
}
