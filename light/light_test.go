package light

import "testing"

func TestSimplePointLight(t *testing.T) {
	args := make(map[string]string)

	args["position"] = "0,0,0"
	args["color"] = "255,255,255,255"

	light, err := NewSimplePointLight(args)
	if err != nil {
		t.Fail()
	}
	lightColor := light.GetColor()
	if lightColor.R != 255 && lightColor.G != 255 && lightColor.B != 255 && lightColor.A != 255 {
		t.Fail()
	}
}
