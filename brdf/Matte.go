package brdf

import (
	geo "GoTracing/geometry"

	"image/color"
)

type Matte struct {
	Color color.RGBA
}

func (matte Matte) Shade(vIn geo.Vector3D, vOut geo.Vector3D, normal geo.Vector3D, hitPoint geo.Point3D, hitLight bool) color.RGBA {
	return matte.Color
}
