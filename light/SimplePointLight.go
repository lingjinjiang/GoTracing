package light

import (
	geo "GoTracing/geometry"
)

type SimplePointLight struct {
	Position geo.Point3D
}

func (light SimplePointLight) GetDirection(point geo.Point3D) geo.Vector3D {
	return point.Sub(light.Position)
}
