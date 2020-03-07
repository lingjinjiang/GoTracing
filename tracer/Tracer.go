package tracer

import (
	geo "GoTracing/geometry"
	obj "GoTracing/object"
	"container/list"
	"image/color"
)

func Tracing(objList list.List, ray geo.Ray) (bool, *obj.Object, *geo.Point3D) {
	var min float64 = -1.0
	var isHit bool = false
	var hitPoint geo.Point3D
	var hitObject obj.Object = nil

	// if the ray hit multi objects, so return the nearest one
	for i := objList.Front(); i != nil; i = i.Next() {
		obj := i.Value.(obj.Object)
		currentHit, currentHitPoint := obj.Hit(ray)
		distance := ray.Endpoint.Sub(currentHitPoint).Length()
		if currentHit {
			isHit = true
			if min == -1.0 || (distance < min && distance > 0) {
				min = distance
				hitPoint = currentHitPoint
				hitObject = obj
			}
		}
	}

	if isHit {
		return true, &hitObject, &hitPoint
	} else {
		return false, nil, nil
	}
}

func GetColor(isHit bool, hitObject obj.Object, hitPoint geo.Point3D, ray geo.Ray) color.RGBA {
	if isHit {
		vOut := ray.Direction.Opposite()

		// temporary light direction
		vIn := geo.Vector3D{
			X: 1,
			Y: 1,
			Z: 1,
		}.Normalize()

		return hitObject.GetMaterial().Shade(vIn, vOut, hitObject.NormalVector(hitPoint), hitPoint)
	} else {
		return color.RGBA{20, 20, 20, 255}
	}
}
