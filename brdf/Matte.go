package brdf

import (
	"image/color"
)

type Matte struct {
	Color color.RGBA
}

func (matte Matte) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	return matte.Color
}
