package material

import (
	"image/color"
	"math"
)

type Material interface {
	// return the color of material with lights
	Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA
	IsSpecular() (bool, float64)
}

func FixRGBA(value float64) float64 {
	if value > 255.0 {
		return 255.0
	}
	return value
}

func FixColor(color1 color.RGBA, color2 color.RGBA) color.RGBA {
	r := float64(color1.R) + float64(color2.R)
	g := float64(color1.G) + float64(color2.G)
	b := float64(color1.B) + float64(color2.B)
	a := float64(color1.A) + float64(color2.A)

	max := math.Max(r, math.Max(g, b))

	if max > 255.0 {
		r = r / max
		g = g / max
		b = b / max
	}

	if a > 255 {
		a = 255
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}
