package config

import (
	"GoTracing/light"
	"GoTracing/material"
	"GoTracing/object"
	"container/list"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func NewConfiguration(filename string) (Configuration, error) {
	var conf Configuration = Configuration{}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf, nil
}

type Configuration struct {
	Main    MainConfig   `yaml:"main"`
	Camra   CamraConfig  `yaml:"camra"`
	Objects []ObjectInfo `yaml:"objects"`
	Lights  []LightInfo  `yaml:"lights"`
}

type MainConfig struct {
	Width         float64 `yaml:"width"`
	Height        float64 `yaml:"height"`
	Output        string  `yaml:"output"`
	RenderThreads int     `yaml:"renderThreads"`
}

type CamraConfig struct {
	Positsion string `yaml:"position"`
	Sample    int    `yaml:"sample"`
}

type ObjectInfo struct {
	Name     string            `yaml:"name"`
	Kind     string            `yaml:"kind"`
	Args     map[string]string `yaml:"args"`
	Material MaterialInfo      `yaml:"material"`
}

type LightInfo struct {
	Name string            `yaml:"name"`
	Kind string            `yaml:"kind"`
	Args map[string]string `yaml:"args"`
}

type MaterialInfo struct {
	Kind string            `yaml:"kind"`
	Args map[string]string `yaml:"args"`
}

func GenerateObjects(conf Configuration) *list.List {
	objects := conf.Objects
	if objects == nil {
		log.Println("No object is defined in configuration.")
		return nil
	}

	objectsInit := newObjectsInitializers()
	materialInit := newMaterailInitializers()

	objList := list.New()
	for _, objInfo := range objects {
		if objectsInit[objInfo.Kind] == nil {
			continue
		}
		materialInitFunc := materialInit[objInfo.Material.Kind]
		if materialInitFunc == nil {
			log.Println("[WARN] Unknown material: ", objInfo.Material.Kind)
			materialInitFunc = material.NewDefaultMaterial
		}
		material, err := materialInitFunc(objInfo.Material.Args)
		if err != nil {
			log.Fatal("Can't parse the object material: ", objInfo.Name, err)
			continue
		}
		obj, err := objectsInit[objInfo.Kind](material, objInfo.Args)
		if err != nil {
			log.Fatal("Can't initialize the object: ", objInfo.Name, err)
			continue
		}
		objList.PushBack(obj)
	}

	return objList
}

func GenerateLights(conf Configuration) *list.List {
	lights := conf.Lights
	if lights == nil {
		log.Println("No light is defined in configuration.")
		return nil
	}

	lightsInit := newLightInitializers()
	lightList := list.New()
	for _, lightInfo := range lights {
		if lightsInit[lightInfo.Kind] == nil {
			continue
		}
		l, err := lightsInit[lightInfo.Kind](lightInfo.Args)
		if err != nil {
			log.Fatal("Can't initialize the light: ", lightInfo.Name, err)
			continue
		}
		lightList.PushBack(l)
	}

	return lightList
}

type ObjInit func(material material.Material, args map[string]string) (object.Object, error)

func newObjectsInitializers() map[string]ObjInit {
	objectsInit := map[string]ObjInit{}

	objectsInit["Sphere"] = object.NewSphere
	objectsInit["Rect"] = object.NewRect

	return objectsInit
}

type MaterialInit func(args map[string]string) (material.Material, error)

func newMaterailInitializers() map[string]MaterialInit {
	materialInit := map[string]MaterialInit{}
	materialInit["Default"] = material.NewDefaultMaterial
	materialInit["Phong"] = material.NewPhong
	materialInit["SpecularPhong"] = material.NewSpecularPhong
	materialInit["SV_Matte"] = material.NewSVMatte
	return materialInit
}

type LightInit func(args map[string]string) (light.Light, error)

func newLightInitializers() map[string]LightInit {
	lightInit := map[string]LightInit{}
	lightInit["SimplePointLight"] = light.NewSimplePointLight
	return lightInit
}
