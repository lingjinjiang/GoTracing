package geometry

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func (src Point3D) Sub(dest Point3D) Vector3D {
	return Vector3D{
		X: dest.X - src.X,
		Y: dest.Y - src.Y,
		Z: dest.Z - src.Z,
	}
}

func (this Point3D) Add(v Vector3D) Point3D {
	return Point3D{
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	point := &Point3D{
		X: x,
		Y: y,
		Z: z,
	}
	return point
}

func ParsePoint(pointStr string) (*Point3D, error) {
	cood := strings.Split(pointStr, ",")
	if len(cood) != 3 {
		log.Fatal("Can't use the elements to convert point:", pointStr)
		return nil, errors.New("Not enough element to convert point")
	}

	x, err := strconv.ParseFloat(strings.Trim(cood[0], " "), 64)
	if err != nil {
		log.Fatal("Error when parse x coordinate:", cood[0])
		return nil, err
	}

	y, err := strconv.ParseFloat(strings.Trim(cood[1], " "), 64)
	if err != nil {
		log.Fatal("Error when parse x coordinate:", cood[1])
		return nil, err
	}

	z, err := strconv.ParseFloat(strings.Trim(cood[2], " "), 64)
	if err != nil {
		log.Fatal("Error when parse x coordinate:", cood[2])
		return nil, err
	}

	point := Point3D{
		X: x,
		Y: y,
		Z: z,
	}

	return &point, nil
}
