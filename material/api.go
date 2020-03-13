package material

import (
	"image/color"
)

type Material interface {
	Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA
}
