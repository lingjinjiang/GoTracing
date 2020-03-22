package material

import (
	"image/color"
	"math"
)

type SV_Matte struct {
	Color1 color.RGBA
	Color2 color.RGBA
	Size   int
}

func (sv SV_Matte) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	hitPoint := shadeRec.HitPoint
	x := hitPoint.X
	// y := hitPoint.Y;
	z := hitPoint.Z

	lightColor := shadeRec.Light.GetColor()

	// if (int(math.Floor(x / float64(sv.Size))) + int(math.Floor(y / float64(sv.Size))) + int(math.Floor(z / float64(sv.Size)))) % 2 == 0 {
	// 	return sv.Color1;
	// } else {
	// 	return sv.Color2;
	// }
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
