package brdf

import (
	geo "GoTracing/geometry"

	"image/color"
	"math"
)

type SpecularPhong struct {
	Ks    float64
	Kd    float64
	Cd    float64
	Color color.RGBA
}

func (sp SpecularPhong) Shade(vIn geo.Vector3D, vOut geo.Vector3D, normal geo.Vector3D, hitPoint geo.Point3D, hitLight bool, diffuseColor color.RGBA) color.RGBA {

	reflect := sp.ambient() * normal.Dot(vOut)
	if hitLight {
		reflect = reflect + sp.specular(hitPoint, normal, vOut, vIn)
	}

	if reflect > 1 {
		reflect = 1
	}

	diffuse := sp.lambertion() * normal.Dot(vOut)

	finalR := float64(sp.Color.R)*reflect + float64(diffuseColor.R)*diffuse
	if finalR > 255 {
		finalR = 255
	}

	finalG := float64(sp.Color.G)*reflect + float64(diffuseColor.G)*diffuse
	if finalG > 255 {
		finalG = 255
	}

	finalB := float64(sp.Color.B)*reflect + float64(diffuseColor.B)*diffuse
	if finalB > 255 {
		finalB = 255
	}

	return color.RGBA{
		R: uint8(finalR),
		G: uint8(finalG),
		B: uint8(finalB),
		A: 255,
	}
	// return diffuseColor

}

func (sp SpecularPhong) ambient() float64 {
	return sp.Kd * sp.Cd
}

func (sp SpecularPhong) lambertion() float64 {
	return sp.Kd * sp.Cd * invPi
}

func (sp SpecularPhong) specular(hitPoint geo.Point3D, normal geo.Vector3D, vOut geo.Vector3D, vIn geo.Vector3D) float64 {
	var result float64 = 0.0
	NorDotIn := normal.Dot(vIn)
	r := geo.Vector3D{
		X: -vIn.X + 2.0*NorDotIn*normal.X,
		Y: -vIn.Y + 2.0*NorDotIn*normal.Y,
		Z: -vIn.Z + 2.0*NorDotIn*normal.Z,
	}

	RDotOut := r.Dot(vOut)

	if RDotOut > 0.0 {
		result = sp.Ks * math.Pow(RDotOut, 2)
	}

	return result
}
