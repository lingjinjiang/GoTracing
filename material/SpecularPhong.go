package material

import (
	"GoTracing/material/brdf"

	"image/color"
)

type SpecularPhong struct {
	ambient  brdf.BRDF
	diffuse  brdf.BRDF
	specular brdf.GlossySpecular
	Color    color.RGBA
}

// phong with simple mirror reflection
func NewSpecularPhong(ks float64, exp float64, kd float64, color color.RGBA) SpecularPhong {
	phong := SpecularPhong{
		ambient: brdf.Lambertian{
			Kd: kd,
		},
		diffuse: brdf.Lambertian{
			Kd: kd,
		},
		specular: brdf.GlossySpecular{
			Ks:  ks,
			Exp: exp,
		},
		Color: color,
	}

	return phong
}

func (sp SpecularPhong) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	normal := shadeRec.Normal
	vOut := shadeRec.VOut
	vIn := shadeRec.VIn
	hitPoint := shadeRec.HitPoint
	lightColor := shadeRec.Light.GetColor()
	ambientColor := lightColor

	reflectR := float64(ambientColor.R) * sp.ambient.Rho() * normal.Dot(vOut)
	reflectG := float64(ambientColor.G) * sp.ambient.Rho() * normal.Dot(vOut)
	reflectB := float64(ambientColor.B) * sp.ambient.Rho() * normal.Dot(vOut)
	reflectA := float64(ambientColor.A) * sp.ambient.Rho() * normal.Dot(vOut)

	if hitLight {
		reflectR = reflectR + (sp.diffuse.F(hitPoint, normal, vOut, vIn)+sp.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.R)
		reflectG = reflectG + (sp.diffuse.F(hitPoint, normal, vOut, vIn)+sp.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.G)
		reflectB = reflectB + (sp.diffuse.F(hitPoint, normal, vOut, vIn)+sp.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.B)
		reflectA = reflectA + (sp.diffuse.F(hitPoint, normal, vOut, vIn)+sp.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.A)
	}

	reflectR = FixRGBA(float64(sp.Color.R) * reflectR / 255)
	reflectG = FixRGBA(float64(sp.Color.G) * reflectG / 255)
	reflectB = FixRGBA(float64(sp.Color.B) * reflectB / 255)
	reflectA = FixRGBA(float64(sp.Color.A) * reflectB / 255)

	diffuse := sp.diffuse.F(hitPoint, normal, vOut, vIn) * normal.Dot(vOut)

	finalR := reflectR + float64(diffuseColor.R)*diffuse
	if finalR > 255 {
		finalR = 255
	}

	finalG := reflectG + float64(diffuseColor.G)*diffuse
	if finalG > 255 {
		finalG = 255
	}

	finalB := reflectB + float64(diffuseColor.B)*diffuse
	if finalB > 255 {
		finalB = 255
	}

	finalA := reflectA + float64(diffuseColor.A)*diffuse
	if finalA > 255 {
		finalA = 255
	}
	return color.RGBA{
		R: uint8(finalR),
		G: uint8(finalG),
		B: uint8(finalB),
		A: uint8(finalA),
	}
	// return diffuseColor
}
