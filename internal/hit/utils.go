package hit

import (
	"math"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

func refract(vector geom.Vec3d, normal geom.Vec3d, etaTOverEtaI float64) geom.Vec3d {
	cosTheta := geom.DotProduct(vector.Negate(), normal)
	ROutPara := geom.MultiplyVec(etaTOverEtaI, geom.AddTwoVec(vector, geom.MultiplyVec(cosTheta, normal)))
	ROutPerp := geom.MultiplyVec(-math.Sqrt(1-ROutPara.LengthSquared()), normal)
	return geom.AddTwoVec(ROutPara, ROutPerp)
}

func schlick(cos float64, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}

func refelect(vector geom.Vec3d, normal geom.Vec3d) geom.Vec3d {
	dot := geom.DotProduct(vector, normal)
	return geom.SubstractTwoVec(vector, geom.MultiplyVec(2*dot, normal))
}
