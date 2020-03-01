package brdf

import (
	geo "GoTracing/geometry"

	"image/color"
)


type BRDF interface {
	Shade(vIn geo.Vector3D, vOut geo.Vector3D, normal geo.Vector3D, hitPoint geo.Point3D) color.RGBA
}