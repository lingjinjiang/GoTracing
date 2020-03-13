package object

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
)

type Object interface {
	Hit(ray geo.Ray) (bool, geo.Point3D)
	NormalVector(point geo.Point3D) geo.Vector3D
	// SetMaterial(material color.RGBA)
	GetMaterial() material.Material
}
