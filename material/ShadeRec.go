package material

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
)

type ShadeRec struct {
	IsHit    bool
	Ray      geo.Ray
	HitPoint geo.Point3D
	Material Material
	Normal   geo.Vector3D
	VIn      geo.Vector3D
	VOut     geo.Vector3D
	Light    light.Light
}
