package brdf

import geo "GoTracing/geometry"

type BRDF interface {
	F(hitPoint geo.Point3D, normal geo.Vector3D, vOut geo.Vector3D, vIn geo.Vector3D) float64
	Rho() float64
}
