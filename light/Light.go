package light

import (
	geo "GoTracing/geometry"
	"image/color"
)

type Light interface {
	GetDirection(point geo.Point3D) geo.Vector3D
	GetColor() color.RGBA
}
