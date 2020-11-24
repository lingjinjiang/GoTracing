package world

import (
	"GoTracing/pkg/config"
	"errors"
	"log"
	"strconv"
)

type ViewPlane struct {
	Height            float64
	Width             float64
	s                 float64
	Samples           int
	gamma             float64
	inv_gamma         float64
	show_out_of_gamut bool
}

func NewViewPlane(vPlaneInfo config.ViewPlaneInfo) (ViewPlane, error) {
	vp := ViewPlane{}

	if width, err := strconv.ParseFloat(vPlaneInfo.Width, 64); err == nil {
		if width <= 0.0 {
			return vp, errors.New("The width should be a positive value" + vPlaneInfo.Width)
		}
		vp.Width = width
	} else {
		log.Fatal("The width is illegal: ", vPlaneInfo.Width)
		return vp, err
	}

	if height, err := strconv.ParseFloat(vPlaneInfo.Height, 64); err == nil {
		if height <= 0.0 {
			return vp, errors.New("The height should be a positive value" + vPlaneInfo.Height)
		}
		vp.Height = height
	} else {
		log.Fatal("The heigth is illegal: ", vPlaneInfo.Height)
		return vp, err
	}

	if sample, err := strconv.Atoi(vPlaneInfo.Sample); err == nil {
		if sample <= 0 {
			return vp, errors.New("The sample should be a positive value" + vPlaneInfo.Sample)
		}
		vp.Samples = sample
	} else {
		log.Fatal("The sample is illegal: ", vPlaneInfo.Sample)
		return vp, err
	}

	return vp, nil
}
