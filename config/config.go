package config

import (
	"GoTracing/object"
	"container/list"
	"io/ioutil"
	"log"
	"strconv"

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
	Name string            `yaml:"name"`
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

	objList := list.New()
	for _, objInfo := range objects {
		if objectsInit[objInfo.Kind] == nil {
			continue
		}
		obj := objectsInit[objInfo.Kind]()
		if o, isType := obj.(object.Sphere); isType {
			if r, err := strconv.ParseFloat(objInfo.Args["radius"], 64); err != nil {
				o.SetRadius(r)
			}
			objList.PushBack(o)
		}
	}

	return objList
}

type InitFunc func() object.Object

func newObjectsInitializers() map[string]InitFunc {
	objectsInit := map[string]InitFunc{}

	objectsInit["Sphere"] = object.NewConfigSphere

	return objectsInit
}
