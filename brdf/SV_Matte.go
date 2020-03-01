package brdf

import (
	geo "GoTracing/geometry"

	"image/color"
	"math"
)

type SV_Matte struct {
	Color1 color.RGBA
	Color2 color.RGBA
	Size int
}


func (sv SV_Matte) Shade(vIn geo.Vector3D, vOut geo.Vector3D, normal geo.Vector3D, hitPoint geo.Point3D) color.RGBA {
	x := hitPoint.X;
	// y := hitPoint.Y;
	z := hitPoint.Z;
	
	// if (int(math.Floor(x / float64(sv.Size))) + int(math.Floor(y / float64(sv.Size))) + int(math.Floor(z / float64(sv.Size)))) % 2 == 0 {
	// 	return sv.Color1;
	// } else {
	// 	return sv.Color2;
	// }
	if (int(math.Floor(x / float64(sv.Size)))  + int(math.Floor(z / float64(sv.Size)))) % 2 == 0 {
		return sv.Color1;
	} else {
		return sv.Color2;
	}
}