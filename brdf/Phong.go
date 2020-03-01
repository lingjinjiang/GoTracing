package brdf

import (
	geo "GoTracing/geometry"

	"math"
	"image/color"
)

type Phong struct {
	Ks float64
	Kd float64
	Cd float64
	Color color.RGBA
}

var invPi float64 = 0.3183098861837906715

// func (p Phong) Shade(vIn geo.Vector3D, vOut geo.Vector3D, hitObject obj.Object, hitPoint geo.Point3D) color.RGBA {
func (p Phong) Shade(vIn geo.Vector3D, vOut geo.Vector3D, normal geo.Vector3D, hitPoint geo.Point3D) color.RGBA {
	// normal := hitObject.NormalVector(hitPoint)
	// vOut := hitPoint.Sub(*scene.ViewPoint).Normalize()
	// vIn := geo.Vector3D {
	// 	X: 1,
	// 	Y: 1,
	// 	Z: 1,
	// }.Normalize()

	reflect := p.lambertion() * normal.Dot(vOut) + p.specular(hitPoint, normal, vOut, vIn) + p.ambient() * normal.Dot(vOut)

	if reflect > 1 {
		reflect = 1
	}

	// material := hitObject.GetMaterial()
	return color.RGBA{uint8(float64(p.Color.R) * reflect), uint8(float64(p.Color.G) * reflect), uint8(float64(p.Color.B) * reflect), 255}
}

func (p Phong) ambient() float64 {
	return p.Kd * p.Cd
}

func (p Phong) lambertion() float64 {
	return p.Kd * p.Cd * invPi
}

func (p Phong) specular(hitPoint geo.Point3D, normal geo.Vector3D, vOut geo.Vector3D, vIn geo.Vector3D) float64 {
	var result float64 = 0.0
	NorDotIn := normal.Dot(vIn)
	r := geo.Vector3D {
		X: -vIn.X + 2.0 * NorDotIn * normal.X,
		Y: -vIn.Y + 2.0 * NorDotIn * normal.Y,
		Z: -vIn.Z + 2.0 * NorDotIn * normal.Z,
	}

	RDotOut := r.Dot(vOut)

	if RDotOut > 0.0 {
		result = p.Ks * math.Pow(RDotOut, 2)
	}

	return result
}