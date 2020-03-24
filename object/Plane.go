package object

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
)

type Plane struct {
	Position geo.Point3D
	Normal   geo.Vector3D
	material material.Material
}

func (plane Plane) Hit(ray geo.Ray) (bool, geo.Point3D) {
	t := ray.Endpoint.Sub(plane.Position).Dot(plane.Normal) / ray.Direction.Dot(plane.Normal)

	if t < 0.001 {
		return false, geo.Point3D{}
	}
	hitPoint := geo.Point3D{
		X: ray.Endpoint.X + t*ray.Direction.X,
		Y: ray.Endpoint.Y + t*ray.Direction.Y,
		Z: ray.Endpoint.Z + t*ray.Direction.Z,
	}
	return true, hitPoint
}

func (plane Plane) NormalVector(point geo.Point3D) geo.Vector3D {
	return plane.Normal
}

func (plane *Plane) SetMaterial(material material.Material) {
	plane.material = material
}

func (plane Plane) GetMaterial() material.Material {
	return plane.material
}

func (plane Plane) GetPosition() geo.Point3D {
	return geo.Point3D{}
}

func (plane Plane) GetObjX() geo.Vector3D {
	return geo.Vector3D{}
}

func (plane Plane) GetObjZ() geo.Vector3D {
	return geo.Vector3D{}
}
