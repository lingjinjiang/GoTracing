package brdf

import (
	geo "GoTracing/geometry"
)

type ShadeRec struct {
	IsHit    bool
	Ray      geo.Ray
	HitPoint geo.Point3D
	Material BRDF
	Normal   geo.Vector3D
	VIn      geo.Vector3D
	VOut     geo.Vector3D
}
