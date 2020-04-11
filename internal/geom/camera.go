package geom

type Camera struct {
	LowerLeftCorner, Horizontal, Vertical, Origin Vec3d
}

func (cam *Camera) GetRay(u float64, v float64) Ray {
	direction := cam.LowerLeftCorner.Add(cam.Horizontal.Multiply(u))
	direction = direction.Add(cam.Vertical.Multiply(v))
	ray := Ray{cam.Origin, direction}
	return ray
}
