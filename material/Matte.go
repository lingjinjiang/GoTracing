package material

import (
	"image/color"
)

// simple single color material for test
type Matte struct {
	Color color.RGBA
}

func (matte Matte) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	return matte.Color
}
