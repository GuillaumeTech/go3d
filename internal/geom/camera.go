package geom

import "math"

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin Vec3d
}

func (cam *Camera) GetRay(u float64, v float64) Ray {
	direction := cam.LowerLeftCorner.Add(cam.Horizontal.Multiply(u))
	direction = direction.Add(cam.Vertical.Multiply(v))
	direction = direction.Substract(cam.Origin)
	ray := Ray{cam.Origin, direction}
	return ray
}

// func NewCamera(vFov float64, aspect float64) Camera {
// 	theta := degToRadian(vFov)
// 	halfHeight := math.Tan(theta / 2)
// 	halfWidth := aspect * halfHeight
// 	lowerLeftCorner := Vec3d{-halfWidth, -halfHeight, -1.0}
// 	horizontal := Vec3d{2 * halfWidth, 0, 0}
// 	vertical := Vec3d{0, 2 * halfHeight, 0}
// 	origin := Vec3d{0, 0, 0}
// 	return Camera{lowerLeftCorner, horizontal, vertical, origin}

// }

func degToRadian(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func NewCamera(lookFrom Vec3d, lookAt Vec3d, vUp Vec3d, vFov float64, aspect float64) Camera {
	theta := degToRadian(vFov)
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	origin := lookFrom
	w := UnitVector(SubstractTwoVec(lookFrom, lookAt))
	u := UnitVector(CrossProduct(vUp, w))
	v := UnitVector(CrossProduct(w, u))
	halfHeightV := MultiplyVec(halfHeight, v)
	halfWidthU := MultiplyVec(halfWidth, u)

	lowerLeftCorner := SubstractTwoVec(SubstractTwoVec(SubstractTwoVec(lookFrom, halfWidthU), halfHeightV), w)
	horizontal := MultiplyVec(2, halfWidthU)
	vertical := MultiplyVec(2, halfHeightV)
	//print(horizontal.X)
	return Camera{lowerLeftCorner, horizontal, vertical, origin}

}
