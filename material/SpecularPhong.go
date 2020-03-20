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

	reflect := sp.ambient.Rho() * normal.Dot(vOut)
	if hitLight {
		reflect = reflect + sp.specular.F(hitPoint, normal, vOut, vIn)
	}

	if reflect > 1 {
		reflect = 1
	}

	diffuse := sp.diffuse.F(hitPoint, normal, vOut, vIn) * normal.Dot(vOut)

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
