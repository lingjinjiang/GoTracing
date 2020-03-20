package material

import (
	"GoTracing/material/brdf"

	"image/color"
)

type Phong struct {
	ambient  brdf.BRDF
	diffuse  brdf.BRDF
	specular brdf.GlossySpecular
	Color    color.RGBA
}

func NewPhong(ks float64, exp float64, kd float64, color color.RGBA) Phong {
	phong := Phong{
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

func (p Phong) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	normal := shadeRec.Normal
	vOut := shadeRec.VOut
	vIn := shadeRec.VIn
	hitPoint := shadeRec.HitPoint

	reflect := p.ambient.Rho() * normal.Dot(vOut)

	if hitLight {
		reflect = reflect + p.diffuse.F(hitPoint, normal, vOut, vIn) + p.specular.F(hitPoint, normal, vOut, vIn)
	}

	if reflect > 1 {
		reflect = 1
	}

	return color.RGBA{uint8(float64(p.Color.R) * reflect), uint8(float64(p.Color.G) * reflect), uint8(float64(p.Color.B) * reflect), 255}
}
