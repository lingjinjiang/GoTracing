package brdf

import (
	"image/color"
)

type BRDF interface {
	Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA
}
