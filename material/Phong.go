package material

import (
	"GoTracing/material/brdf"
	"GoTracing/util"
	"log"
	"strconv"

	"image/color"
)

type Phong struct {
	ambient  brdf.BRDF
	diffuse  brdf.BRDF
	specular brdf.GlossySpecular
	Color    color.RGBA
}

func NewPhong(args map[string]string) (Material, error) {
	kd, err := strconv.ParseFloat(args["kd"], 64)
	if err != nil {
		log.Fatal("Error when pharse specular phong argments kd:", args["kd"])
		return nil, err
	}

	ks, err := strconv.ParseFloat(args["ks"], 64)
	if err != nil {
		log.Fatal("Error when pharse specular phong argments ks:", args["ks"])
		return nil, err
	}

	exp, err := strconv.ParseFloat(args["exp"], 64)
	if err != nil {
		log.Fatal("Error when pharse specular phong argments exp:", args["exp"])
		return nil, err
	}

	color, err := util.ParseColor(args["color"])
	if err != nil {
		log.Fatal("Error when pharse specular phong argments color:", args["color"])
		return nil, err
	}

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
		Color: *color,
	}

	return &phong, nil
}

func (p Phong) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	normal := shadeRec.Normal
	vOut := shadeRec.VOut
	vIn := shadeRec.VIn
	hitPoint := shadeRec.HitPoint
	lightColor := shadeRec.Light.GetColor()
	ambientColor := lightColor

	reflectR := float64(ambientColor.R) * p.ambient.Rho() * normal.Dot(vOut)
	reflectG := float64(ambientColor.G) * p.ambient.Rho() * normal.Dot(vOut)
	reflectB := float64(ambientColor.B) * p.ambient.Rho() * normal.Dot(vOut)
	reflectA := float64(ambientColor.A) * p.ambient.Rho() * normal.Dot(vOut)

	if hitLight {
		reflectR = reflectR + (p.diffuse.F(hitPoint, normal, vOut, vIn)+p.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.R)
		reflectG = reflectG + (p.diffuse.F(hitPoint, normal, vOut, vIn)+p.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.G)
		reflectB = reflectB + (p.diffuse.F(hitPoint, normal, vOut, vIn)+p.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.B)
		reflectA = reflectA + (p.diffuse.F(hitPoint, normal, vOut, vIn)+p.specular.F(hitPoint, normal, vOut, vIn))*float64(lightColor.A)
	}

	reflectR = FixRGBA(float64(p.Color.R) * reflectR / 255)
	reflectG = FixRGBA(float64(p.Color.G) * reflectG / 255)
	reflectB = FixRGBA(float64(p.Color.B) * reflectB / 255)
	reflectA = FixRGBA(float64(p.Color.A) * reflectB / 255)

	return color.RGBA{uint8(reflectR), uint8(reflectG), uint8(reflectB), uint8(reflectA)}
}

func (p *Phong) SetTraceFunc(TraceFunc) {
}
