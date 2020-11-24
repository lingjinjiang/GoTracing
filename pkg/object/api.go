package object

import (
	geo "GoTracing/pkg/geometry"
	"GoTracing/pkg/material"
)

type Object interface {
	// return true and the hit point if the ray hit the object
	Hit(ray geo.Ray) (bool, geo.Point3D)
	// get normal vector of the hit point
	NormalVector(point geo.Point3D) geo.Vector3D
	// SetMaterial(material color.RGBA)
	GetMaterial() material.Material

	GetPosition() geo.Point3D
	GetLocalX() geo.Vector3D
	GetLocalY() geo.Vector3D
	GetLocalZ() geo.Vector3D
}
