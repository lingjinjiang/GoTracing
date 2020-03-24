package object

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
	"math"
)

type Rect struct {
	Position geo.Point3D
	Normal   geo.Vector3D
	material material.Material
	Width    float64
	WVector  geo.Vector3D // width vector
	Length   float64
	LVector  geo.Vector3D // length vector

}

func (rect Rect) Hit(ray geo.Ray) (bool, geo.Point3D) {
	normal := geo.Vector3D{
		X: rect.LVector.Y*rect.WVector.Z - rect.WVector.Y*rect.LVector.Z,
		Y: rect.LVector.Z*rect.WVector.X - rect.LVector.X*rect.WVector.Z,
		Z: rect.LVector.X*rect.WVector.Y - rect.LVector.Y*rect.WVector.X,
	}

	t := ray.Endpoint.Sub(rect.Position).Dot(normal) / ray.Direction.Dot(normal)

	if t < 0.001 {
		return false, geo.Point3D{}
	}
	hitPoint := geo.Point3D{
		X: ray.Endpoint.X + t*ray.Direction.X,
		Y: ray.Endpoint.Y + t*ray.Direction.Y,
		Z: ray.Endpoint.Z + t*ray.Direction.Z,
	}

	rect.Position.Sub(hitPoint).Dot(rect.WVector)

	// if math.Abs((hitPoint.X-rect.Position.X)*rect.WVector.X+(hitPoint.Y-rect.Position.Y)*rect.WVector.Y+(hitPoint.Z-rect.Position.Z)*rect.WVector.Z) > rect.Width/2 || math.Abs((hitPoint.X-rect.Position.X)*rect.LVector.X+(hitPoint.Y-rect.Position.Y)*rect.LVector.Y+(hitPoint.Z-rect.Position.Z)*rect.LVector.Z) > rect.Length/2 {
	// 	return false, geo.Point3D{}
	// }
	if math.Abs(rect.Position.Sub(hitPoint).Dot(rect.WVector)) > rect.Width/2 || math.Abs(rect.Position.Sub(hitPoint).Dot(rect.LVector)) > rect.Length/2 {
		return false, geo.Point3D{}
	}

	return true, hitPoint
}

func (rect Rect) NormalVector(point geo.Point3D) geo.Vector3D {
	return geo.Vector3D{
		X: rect.LVector.Y*rect.WVector.Z - rect.WVector.Y*rect.LVector.Z,
		Y: rect.LVector.Z*rect.WVector.X - rect.LVector.X*rect.WVector.Z,
		Z: rect.LVector.X*rect.WVector.Y - rect.LVector.Y*rect.WVector.X,
	}
}

func (rect *Rect) SetMaterial(material material.Material) {
	rect.material = material
}

func (rect Rect) GetMaterial() material.Material {
	return rect.material
}

func (rect Rect) GetPosition() geo.Point3D {
	return rect.Position
}

func (rect Rect) GetObjX() geo.Vector3D {
	return rect.LVector
}

func (rect Rect) GetObjZ() geo.Vector3D {
	return rect.WVector
}
