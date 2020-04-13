package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	"container/list"
	"image/color"
)

var BACKGROUND color.RGBA = color.RGBA{20, 20, 20, 255}

type Tracer interface {
	Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec
	GetColor(shadeRec material.ShadeRec, objList list.List, light light.Light) color.RGBA
	Tracing2(objList list.List, shadeRec *material.ShadeRec) color.RGBA
}
