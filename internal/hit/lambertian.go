package hit

import (
	"math"
	"math/rand"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Lambertian struct {
	Albedo geom.Vec3d
}

func (l Lambertian) Scatter(rayIn geom.Ray, hitRecord *HitRecord, attenutaion *geom.Vec3d, scattered *geom.Ray) bool {
	var etaTOverEtaI float64
	if hitRecord.FrontFace {
		etaTOverEtaI = 1 / 1.1
	} else {
		etaTOverEtaI = 1.1
	}
	unitDir := geom.UnitVector(rayIn.Direction)
	cosTheta := math.Min(geom.DotProduct(unitDir.Negate(), hitRecord.Normal), 1)
	reflectProbability := schlick(cosTheta, etaTOverEtaI)
	randomFloat := rand.Float64()

	if randomFloat < reflectProbability {
		refelected := refelect(unitDir, hitRecord.Normal)
		scattered.Direction = refelected
		scattered.Origin = hitRecord.P
		*attenutaion = l.Albedo
		return true
	}

	scatterDir := geom.AddTwoVec(geom.RandomUnitVector(), geom.AddTwoVec(hitRecord.P, hitRecord.Normal))
	scattered.Direction = scatterDir
	scattered.Origin = hitRecord.P
	*attenutaion = l.Albedo
	return true
}
