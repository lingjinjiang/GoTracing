package config

import (
	"testing"
)

func TestYamlConfig(t *testing.T) {
	conf := NewYamlConfiguration("config.yaml")
	if conf.Main.Output != "/home/example.jpg" {
		t.Log(conf.Main.Output)
		t.Fail()
	}

	if conf.Camra.Sample != 16 {
		t.Log(conf.Camra.Sample)
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
}
