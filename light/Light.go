package light

import (
	geo "GoTracing/geometry"
	"image/color"
)

type Light interface {
	// get the direction from hit point to light
	GetDirection(point geo.Point3D) geo.Vector3D
	// return the color of light
	GetColor() color.RGBA
}
