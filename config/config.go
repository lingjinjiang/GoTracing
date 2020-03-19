package config

import "github.com/go-ini/ini"

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
