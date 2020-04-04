package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	obj "GoTracing/object"
	"container/list"
	"image/color"
)

var BACKGOUND color.RGBA = color.RGBA{20, 20, 20, 255}

// tracing the ray and return the shade info of this ray
func Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec {
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

	shadeRec := material.ShadeRec{}

	if isHit {
		shadeRec.IsHit = true
		shadeRec.Material = hitObject.GetMaterial()
		shadeRec.HitPoint = hitPoint
		shadeRec.Normal = hitObject.NormalVector(hitPoint)
		shadeRec.Ray = ray
		shadeRec.Light = light
		shadeRec.ObjPosition = hitObject.GetPosition()
		shadeRec.ObjX = hitObject.GetLocalX()
		shadeRec.ObjY = hitObject.GetLocalY()
		shadeRec.ObjZ = hitObject.GetLocalZ()
	} else {
		shadeRec.IsHit = false
	}
	return shadeRec
}

// get the color of ray
func GetColor(shadeRec material.ShadeRec, objList list.List, light light.Light) color.RGBA {
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
		lightShadeRec := Tracing(objList, light, lcoalRay)

		// simple diffuse
		diffuseIn := localNormal.Add(shadeRec.Ray.Direction.Normalize())
		diffuseRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: diffuseIn,
		}
		diffuseShadeRec := Tracing(objList, light, diffuseRay)
		diffuseShadeRec.VIn = light.GetDirection(diffuseShadeRec.HitPoint).Normalize()
		diffuseShadeRec.VOut = diffuseIn
		var diffuseColor color.RGBA
		if diffuseShadeRec.IsHit {
			diffuseColor = diffuseShadeRec.Material.Shade(diffuseShadeRec, diffuseShadeRec.IsHit, BACKGOUND)
		} else {
			diffuseColor = BACKGOUND
		}

		return shadeRec.Material.Shade(shadeRec, !lightShadeRec.IsHit, diffuseColor)
	} else {
		return BACKGOUND
	}
}
