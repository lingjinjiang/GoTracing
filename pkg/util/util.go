package util

import (
	"image/color"
	"log"
	"strconv"
	"strings"
)

func ParseColor(colorStr string) (*color.RGBA, error) {
	rgb := strings.Split(colorStr, ",")
	if len(rgb) != 4 {
		log.Fatal("not enough element to convert color:", colorStr)
		return nil, nil
	}
	r, err := strconv.ParseUint(strings.Trim(rgb[0], " "), 10, 8)
	if err != nil {
		log.Fatal("Error when parse color Red element:", rgb[0])
		return nil, err
	}
	g, err := strconv.ParseUint(strings.Trim(rgb[1], " "), 10, 8)
	if err != nil {
		log.Fatal("Error when parse color Green element:", rgb[1])
		return nil, err
	}
	b, err := strconv.ParseUint(strings.Trim(rgb[2], " "), 10, 8)
	if err != nil {
		log.Fatal("Error when parse color Blue element:", rgb[2])
		return nil, err
	}
	a, err := strconv.ParseUint(strings.Trim(rgb[3], " "), 10, 8)
	if err != nil {
		log.Fatal("Error when parse color Alpha element:", rgb[3])
		return nil, err
	}

	return &color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}
