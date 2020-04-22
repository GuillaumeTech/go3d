package hit

import (
	"math"
	"math/rand"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) Scatter(rayIn geom.Ray, hitRecord *HitRecord, attenutaion *geom.Vec3d, scattered *geom.Ray) bool {
	*attenutaion = geom.Vec3d{1, 1, 1}
	var etaTOverEtaI float64
	if hitRecord.FrontFace {
		etaTOverEtaI = 1 / d.RefractionIndex
	} else {
		etaTOverEtaI = d.RefractionIndex
	}
	unitDir := geom.UnitVector(rayIn.Direction)

	cosTheta := math.Min(geom.DotProduct(unitDir.Negate(), hitRecord.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	reflectProbability := schlick(cosTheta, etaTOverEtaI)
	randomFloat := rand.Float64()

	if etaTOverEtaI*sinTheta > 1 || randomFloat < reflectProbability {
		refelected := refelect(unitDir, hitRecord.Normal)
		scattered.Direction = refelected
		scattered.Origin = hitRecord.P
		return true
	}

	refracted := refract(unitDir, hitRecord.Normal, etaTOverEtaI)
	scattered.Direction = refracted
	scattered.Origin = hitRecord.P
	return true
}
