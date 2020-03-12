package tracer

import (
	"GoTracing/brdf"
	geo "GoTracing/geometry"
	"GoTracing/light"
	obj "GoTracing/object"
	"container/list"
	"image/color"
)

var BACKGOUND color.RGBA = color.RGBA{20, 20, 20, 255}

func Tracing(objList list.List, ray geo.Ray) brdf.ShadeRec {
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

	shadeRec := brdf.ShadeRec{}

	if isHit {
		shadeRec.IsHit = true
		shadeRec.Material = hitObject.GetMaterial()
		shadeRec.HitPoint = hitPoint
		shadeRec.Normal = hitObject.NormalVector(hitPoint)
		shadeRec.Ray = ray
	} else {
		shadeRec.IsHit = false
	}
	return shadeRec
}

func GetColor(shadeRec brdf.ShadeRec, objList list.List, light light.Light) color.RGBA {
	if shadeRec.IsHit {
		localNormal := shadeRec.Normal
		shadeRec.VOut = shadeRec.Ray.Direction.Opposite()

		// temporary light direction
		lightIn := light.GetDirection(shadeRec.HitPoint).Normalize()
		shadeRec.VIn = lightIn

		lcoalRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: lightIn,
		}

		// simple shadow, if the ray from object hit point to light hit some other objects, then the point is in shadow
		lightShadeRec := Tracing(objList, lcoalRay)

		// simple diffuse
		diffuseIn := localNormal.Add(shadeRec.Ray.Direction.Normalize())
		diffuseRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: diffuseIn,
		}
		diffuseShadeRec := Tracing(objList, diffuseRay)
		diffuseShadeRec.VIn = lightIn
		diffuseShadeRec.VOut = diffuseIn
		var diffuseColor color.RGBA
		if diffuseShadeRec.IsHit {
			diffuseColor = diffuseShadeRec.Material.Shade(diffuseShadeRec, true, BACKGOUND)
		} else {
			diffuseColor = BACKGOUND
		}

		return shadeRec.Material.Shade(shadeRec, !lightShadeRec.IsHit, diffuseColor)
	} else {
		return BACKGOUND
	}
}
