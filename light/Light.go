package light

import (
	geo "GoTracing/geometry"
)

type Light interface {
	GetDirection(point geo.Point3D) geo.Vector3D
}
