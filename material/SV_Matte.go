package material

import (
	"image/color"
	"log"
	"math"
	"strconv"
)

// simple two color materail for test
type SV_Matte struct {
	Color1 color.RGBA
	Color2 color.RGBA
	Size   int
}

func (sv SV_Matte) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	hitPoint := shadeRec.HitPoint

	objPos := shadeRec.ObjPosition
	objX := shadeRec.ObjX
	objZ := shadeRec.ObjZ
	x := ((hitPoint.X-objPos.X)*objZ.Z - (hitPoint.Z-objPos.Z)*objZ.X) / (objX.X*objZ.Z - objX.Z*objZ.X)
	z := ((hitPoint.X-objPos.X)*objX.Z - (hitPoint.Z-objPos.Z)*objX.X) / (objZ.X*objX.Z - objX.X*objZ.Z)

	lightColor := shadeRec.Light.GetColor()

	delta := 0.3
	if hitLight {
		delta = 1
	}
	var shadeColor color.RGBA
	if (int(math.Floor(x/float64(sv.Size)))+int(math.Floor(z/float64(sv.Size))))%2 == 0 {
		shadeColor = color.RGBA{
			R: uint8(FixRGBA(float64(sv.Color1.R) * float64(lightColor.R) / 255.0)),
			G: uint8(FixRGBA(float64(sv.Color1.G) * float64(lightColor.G) / 255.0)),
			B: uint8(FixRGBA(float64(sv.Color1.B) * float64(lightColor.B) / 255.0)),
			A: uint8(FixRGBA(float64(sv.Color1.A) * float64(lightColor.A) / 255.0)),
		}
	} else {
		shadeColor = color.RGBA{
			R: uint8(FixRGBA(float64(sv.Color2.R) * float64(lightColor.R) / 255.0)),
			G: uint8(FixRGBA(float64(sv.Color2.G) * float64(lightColor.G) / 255.0)),
			B: uint8(FixRGBA(float64(sv.Color2.B) * float64(lightColor.B) / 255.0)),
			A: uint8(FixRGBA(float64(sv.Color2.A) * float64(lightColor.A) / 255.0)),
		}
	}

	return color.RGBA{
		R: uint8(float64(shadeColor.R) * delta),
		G: uint8(float64(shadeColor.G) * delta),
		B: uint8(float64(shadeColor.B) * delta),
	}
}

func NewSVMatte(args map[string]string) (Material, error) {
	size, err := strconv.Atoi(args["size"])
	if err != nil {
		log.Fatal("Error when parseing argment size:", args["size"])
		return nil, err
	}

	color1, err := ParseColor(args["color1"])
	if err != nil {
		log.Fatal("Error when parseing argment color1:", args["color2"])
		return nil, err
	}

	color2, err := ParseColor(args["color2"])
	if err != nil {
		log.Fatal("Error when parseing argment color2:", args["color2"])
		return nil, err
	}

	matte := SV_Matte{
		Color1: *color1,
		Color2: *color2,
		Size:   size,
	}

	return matte, nil
}
