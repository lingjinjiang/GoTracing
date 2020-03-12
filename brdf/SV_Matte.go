package brdf

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
		shadeColor = sv.Color1
	} else {
		shadeColor = sv.Color2
	}

	return color.RGBA{
		R: uint8(float64(shadeColor.R) * delta),
		G: uint8(float64(shadeColor.G) * delta),
		B: uint8(float64(shadeColor.B) * delta),
	}
}
