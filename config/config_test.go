package config

import (
	"GoTracing/tracer"
	"testing"
)

func TestYamlConfig(t *testing.T) {

	conf, _ := NewConfiguration("config.yaml")
	if conf.Main.Output != "/home/example.jpg" {
		t.Log(conf.Main.Output)
		t.Fail()
	}

	if conf.Camera.VPlane.Sample != "16" {
		t.Log(conf.Camera.VPlane.Sample)
		t.Fail()
	}

	if len(conf.Objects) != 2 {
		t.Log(len(conf.Objects))
		t.Fail()
	}

	if conf.Objects[0].Name != "whitesphere" {
		t.Log(conf.Objects[0].Name)
		t.Fail()
	}

	if conf.Objects[1].Args["position"] != "1,1,1" {
		t.Log(conf.Objects[1].Args["position"])
		t.Fail()
	}

	objList := GenerateObjects(conf, tracer.SimpleTracer{})

	if objList.Len() != 1 {
		t.Log(objList.Len())
		t.Fail()
	}

	// obj := objList.Front().Value.(object.Object)
	// if o, isType := obj.(object.Sphere); isType {
	// 	if o.Radius() != 50.0 {
	// 		t.Log(o.Radius())
	// 		t.Fail()
	// 	}
	// } else {
	// 	t.Log("the type is not sphere")
	// 	t.Fail()
	// }
}
