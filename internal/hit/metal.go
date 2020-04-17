package hit

import (
	"math/rand"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Metal struct {
	Albedo geom.Vec3d
	Fuzz   float64
}

func (m Metal) Scatter(rayIn geom.Ray, hitRecord *HitRecord, attenutaion *geom.Vec3d, scattered *geom.Ray) bool {
	reflected := refelect(geom.UnitVector(rayIn.Direction), hitRecord.Normal)
	fuzz := geom.AddTwoVec(reflected, geom.MultiplyVec(m.Fuzz*rand.Float64(), geom.RandomUnitVector()))
	scattered.Direction = fuzz
	scattered.Origin = hitRecord.P
	*attenutaion = m.Albedo
	return geom.DotProduct(scattered.Direction, hitRecord.Normal) > 0
}

func refelect(vector geom.Vec3d, normal geom.Vec3d) geom.Vec3d {
	dot := geom.DotProduct(vector, normal)
	return geom.SubstractTwoVec(vector, geom.MultiplyVec(2*dot, normal))
}
