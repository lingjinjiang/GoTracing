package config

import (
	"io/ioutil"
	"log"

	"github.com/go-ini/ini"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Output        string
	Width         float64
	Height        float64
	Sample        int
	RenderThreads int
}

func NewConfiguration(filename string) (*Configuration, error) {
	context, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	config := Configuration{}

	imageInfo := context.Section("image")
	if imageInfo != nil {
		config.Width, err = imageInfo.Key("width").Float64()
		if err != nil {
			return nil, err
		}

		config.Height, err = imageInfo.Key("height").Float64()
		if err != nil {
			return nil, err
		}

		config.Sample, err = imageInfo.Key("sample").Int()
		if err != nil {
			return nil, err
		}

		config.Output = imageInfo.Key("path").String()

		config.RenderThreads, err = imageInfo.Key("renderThreads").Int()
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}

type YamlConfiguration struct {
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

type KeyValue struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func NewYamlConfiguration(filename string) YamlConfiguration {
	var conf YamlConfiguration = YamlConfiguration{}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
