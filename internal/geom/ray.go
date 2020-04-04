package geom

type Ray struct {
	Origin, Direction Vec3d
}

func (r *Ray) at(t float64) Vec3d {
	return r.Origin.Add(r.Direction.Multiply(t))
}
