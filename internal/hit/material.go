package hit

import "github.com/GuillaumeTech/3dgo/internal/geom"

type Material interface {
	Scatter(rayIn geom.Ray, hitRecord *HitRecord, attenutaion *geom.Vec3d, scattered *geom.Ray) bool
}
