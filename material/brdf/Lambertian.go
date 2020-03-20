package brdf

import geo "GoTracing/geometry"

type Lambertian struct {
	Kd float64
}

var invPi float64 = 0.3183098861837906715

func NewLambertian() Lambertian {
	return Lambertian{}
}

func (l Lambertian) F(hitPoint geo.Point3D, normal geo.Vector3D, vOut geo.Vector3D, vIn geo.Vector3D) float64 {
	return l.Kd * invPi
}

func (l Lambertian) Rho() float64 {
	return l.Kd
}
