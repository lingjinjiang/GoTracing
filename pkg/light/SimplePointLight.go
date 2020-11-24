package light

import (
	geo "GoTracing/pkg/geometry"
	"GoTracing/pkg/util"
	"errors"
	"image/color"
	"log"
	"strconv"
)

type SimplePointLight struct {
	Position geo.Point3D
	Color    color.RGBA
	Ls       float64
}

func (light SimplePointLight) GetDirection(point geo.Point3D) geo.Vector3D {
	return point.Sub(light.Position)
}

func (light SimplePointLight) GetColor() color.RGBA {
	return light.Color
}

func NewSimplePointLight(args map[string]string) (Light, error) {
	light := SimplePointLight{}
	if position, err := geo.ParsePoint(args["position"]); err == nil {
		light.Position = *position
	} else {
		log.Fatal("The postion is illegal: ", args["position"])
		return nil, err
	}

	if ls, err := strconv.ParseFloat(args["ls"], 64); err == nil {
		if ls <= 0.0 {
			return nil, errors.New("The ls should be a positive value" + args["ls"])
		}
		light.Ls = ls
	} else {
		log.Fatal("The ls is illegal: ", args["ls"])
		return nil, err
	}

	color, err := util.ParseColor(args["color"])
	if err != nil {
		log.Fatal("Error when pharse specular phong argments color:", args["color"])
		return nil, err
	} else {
		light.Color = *color
	}

	return light, nil
}
