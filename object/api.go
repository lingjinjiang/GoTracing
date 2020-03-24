package object

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
)

type Object interface {
	// return true and the hit point if the ray hit the object
	Hit(ray geo.Ray) (bool, geo.Point3D)
	// get normal vector of the hit point
	NormalVector(point geo.Point3D) geo.Vector3D
	// SetMaterial(material color.RGBA)
	GetMaterial() material.Material

	GetPosition() geo.Point3D
	GetObjX() geo.Vector3D
	GetObjZ() geo.Vector3D
}
