package brdf

import (
	geo "GoTracing/geometry"
	"math"
)

type GlossySpecular struct {
	Ks  float64
	Exp float64
}

func (g GlossySpecular) F(hitPoint geo.Point3D, normal geo.Vector3D, vOut geo.Vector3D, vIn geo.Vector3D) float64 {
	var result float64 = 0.0
	NorDotIn := normal.Dot(vIn)
	r := geo.Vector3D{
		X: -vIn.X + 2.0*NorDotIn*normal.X,
		Y: -vIn.Y + 2.0*NorDotIn*normal.Y,
		Z: -vIn.Z + 2.0*NorDotIn*normal.Z,
	}

	RDotOut := r.Dot(vOut)

	if RDotOut > 0.0 {
		result = g.Ks * math.Pow(RDotOut, g.Exp)
	}

	return result
}

func (g GlossySpecular) Rho() float64 {
	return g.Ks
}
