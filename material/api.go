package material

import (
	"image/color"
)

type Material interface {
	Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA
}

func FixRGBA(value float64) float64 {
	if value > 255.0 {
		return 255.0
	}
	return value
}
