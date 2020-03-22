package light

import (
	geo "GoTracing/geometry"
	"image/color"
)

type SimplePointLight struct {
	Position geo.Point3D
	Color    color.RGBA
	Ls       float64
}

func (light SimplePointLight) GetDirection(point geo.Point3D) geo.Vector3D {
	return point.Sub(light.Position)
}

func (light SimplePointLight) GetColor() color.RGBA {
	return light.Color
}
