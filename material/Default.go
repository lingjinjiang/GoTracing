package material

import "image/color"

type Default struct {
}

func (d Default) Shade(shadeRec ShadeRec, hitLight bool, diffuseColor color.RGBA) color.RGBA {
	return color.RGBA{0, 0, 0, 255}
}

func NewDefaultMaterial(args map[string]string) (Material, error) {
	return Default{}, nil
}
