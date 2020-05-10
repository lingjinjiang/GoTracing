package material

import (
	geo "GoTracing/geometry"
	"GoTracing/material/brdf"
	"GoTracing/util"
	"log"
	"math"
	"strconv"

	"image/color"
)

type SpecularPhong struct {
	ambient  brdf.BRDF
	diffuse  brdf.BRDF
	specular brdf.GlossySpecular
	Color    color.RGBA
	fr       float64
	tracing  TraceFunc
}

// phong with simple mirror reflection
func NewSpecularPhong(args map[string]string) (Material, error) {
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
		Color: *color,
		fr:    0.2,
	}

	return &phong, nil
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
		reflectR = reflectR + sp.specular.F(hitPoint, normal, vOut, vIn)*float64(lightColor.R)
		reflectG = reflectG + sp.specular.F(hitPoint, normal, vOut, vIn)*float64(lightColor.G)
		reflectB = reflectB + sp.specular.F(hitPoint, normal, vOut, vIn)*float64(lightColor.B)
		reflectA = reflectA + sp.specular.F(hitPoint, normal, vOut, vIn)*float64(lightColor.A)
	}

	diffuse := sp.diffuse.F(hitPoint, normal, vOut, vIn) * normal.Dot(vOut)

	reflectR = reflectR + float64(diffuseColor.R)*diffuse

	reflectG = reflectG + float64(diffuseColor.G)*diffuse

	reflectB = reflectB + float64(diffuseColor.B)*diffuse

	max := math.Max(reflectR, math.Max(reflectG, reflectB))
	if max > 255 {
		reflectR = reflectR / max
		reflectG = reflectG / max
		reflectB = reflectB / max
	}

	reflectA = reflectA + float64(diffuseColor.A)*diffuse
	if reflectA > 255 {
		reflectA = 255
	}

	finalR := float64(sp.Color.R) * reflectR / 255
	finalG := float64(sp.Color.G) * reflectG / 255
	finalB := float64(sp.Color.B) * reflectB / 255
	finalA := float64(sp.Color.A) * reflectA / 255

	finalColor := color.RGBA{
		R: uint8(finalR),
		G: uint8(finalG),
		B: uint8(finalB),
		A: uint8(finalA),
	}

	if shadeRec.Depth > 0 {
		rayDirect := shadeRec.Ray.Direction.Normalize()
		reflectIn := geo.Vector3D{
			X: rayDirect.X - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.X,
			Y: rayDirect.Y - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.Y,
			Z: rayDirect.Z - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.Z,
		}

		lcoalRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: reflectIn,
		}

		localShadeRec := ShadeRec{
			Light:   shadeRec.Light,
			Depth:   shadeRec.Depth,
			Ray:     lcoalRay,
			ObjList: shadeRec.ObjList,
		}

		reflectColor := sp.tracing(shadeRec.ObjList, &localShadeRec)
		reflectColor.R = uint8(float64(reflectColor.R) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * sp.fr)
		reflectColor.G = uint8(float64(reflectColor.G) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * sp.fr)
		reflectColor.B = uint8(float64(reflectColor.B) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * sp.fr)

		finalColor = FixColor(finalColor, reflectColor)
	}

	return finalColor
}

func (sp *SpecularPhong) SetTraceFunc(tracing TraceFunc) {
	sp.tracing = tracing
}
