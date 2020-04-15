package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	obj "GoTracing/object"
	"container/list"
	"image/color"
	"math"
)

type Whitted struct {
	maxDepth uint
}

func (t Whitted) Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec {
	return material.ShadeRec{}
}

func (t Whitted) GetColor(shadeRec material.ShadeRec, objList list.List, light light.Light) color.RGBA {
	return color.RGBA{}
}

func NewWhitted() Tracer {
	return Whitted{
		maxDepth: 2,
	}
}

func (t Whitted) Tracing2(objList list.List, shadeRec *material.ShadeRec) color.RGBA {
	var min float64 = -1.0
	var isHit bool = false
	var hitPoint geo.Point3D
	var hitObject obj.Object = nil
	ray := shadeRec.Ray
	light := shadeRec.Light
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
	shadeRec.Depth -= 1

	var color color.RGBA
	if isHit {
		shadeRec.IsHit = true
		shadeRec.Material = hitObject.GetMaterial()
		shadeRec.HitPoint = hitPoint
		shadeRec.Normal = hitObject.NormalVector(hitPoint)
		shadeRec.ObjPosition = hitObject.GetPosition()
		shadeRec.ObjX = hitObject.GetLocalX()
		shadeRec.ObjY = hitObject.GetLocalY()
		shadeRec.ObjZ = hitObject.GetLocalZ()
		shadeRec.VOut = shadeRec.Ray.Direction.Opposite()
		lightIn := light.GetDirection(shadeRec.HitPoint).Normalize()
		shadeRec.VIn = lightIn
		lightShadeRec := material.ShadeRec{
			IsHit: false,
			Light: shadeRec.Light,
			Ray: geo.Ray{
				Endpoint:  shadeRec.HitPoint,
				Direction: lightIn,
			},
			Depth: 1,
		}

		t.Tracing2(objList, &lightShadeRec)
		color = shadeRec.Material.Shade(*shadeRec, !lightShadeRec.IsHit, BACKGROUND)

		isSpecular, fr := shadeRec.Material.IsSpecular()
		if shadeRec.Depth > 0 && isSpecular {
			rayDirect := shadeRec.Ray.Direction.Normalize()
			reflectIn := geo.Vector3D{
				X: rayDirect.X - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.X,
				Y: rayDirect.Y - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.Y,
				Z: rayDirect.Z - 2*rayDirect.Dot(shadeRec.Normal)*shadeRec.Normal.Z,
			}

			lcoalRay := geo.Ray{
				Endpoint:  shadeRec.HitPoint,
				Direction: reflectIn,
			}

			localShadeRec := material.ShadeRec{
				Light: shadeRec.Light,
				Depth: shadeRec.Depth,
				Ray:   lcoalRay,
			}

			reflectColor := t.Tracing2(objList, &localShadeRec)
			reflectColor.R = uint8(float64(reflectColor.R) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * fr)
			reflectColor.G = uint8(float64(reflectColor.G) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * fr)
			reflectColor.B = uint8(float64(reflectColor.B) * math.Abs(shadeRec.Normal.Dot(shadeRec.VOut.Normalize())) * fr)

			color = material.FixColor(color, reflectColor)
		}
	} else {
		shadeRec.IsHit = false
		color = BACKGROUND
	}

	return color
}

func (t Whitted) GetMaxDepth() uint {
	return t.maxDepth
}
