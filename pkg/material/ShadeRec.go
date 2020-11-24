package material

import (
	geo "GoTracing/pkg/geometry"
	"GoTracing/pkg/light"
	"container/list"
)

type ShadeRec struct {
	IsHit       bool
	Ray         geo.Ray
	HitPoint    geo.Point3D
	Material    Material
	Normal      geo.Vector3D
	VIn         geo.Vector3D
	VOut        geo.Vector3D
	Light       light.Light
	ObjPosition geo.Point3D
	ObjX        geo.Vector3D
	ObjY        geo.Vector3D
	ObjZ        geo.Vector3D
	Depth       uint // times for tracing
	ObjList     list.List
}
