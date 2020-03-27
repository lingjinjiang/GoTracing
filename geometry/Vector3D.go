package geometry

import (
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
)

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (this Vector3D) Normalize() Vector3D {
	length := this.Length()
	return Vector3D{
		X: this.X / length,
		Y: this.Y / length,
		Z: this.Z / length,
	}
}

func (this Vector3D) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

func (this Vector3D) Add(v Vector3D) Vector3D {
	return Vector3D{
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

func (this Vector3D) Sub(v Vector3D) Vector3D {
	return Vector3D{
		X: this.X - v.X,
		Y: this.Y - v.Y,
		Z: this.Z - v.Z,
	}
}

func (left Vector3D) Dot(right Vector3D) float64 {
	return left.X*right.X + left.Y*right.Y + left.Z*right.Z
}

func (this Vector3D) Opposite() Vector3D {
	return Vector3D{
		X: -this.X,
		Y: -this.Y,
		Z: -this.Z,
	}
}

func ParseVector(vectorStr string) (*Vector3D, error) {
	cood := strings.Split(vectorStr, ",")
	if len(cood) != 3 {
		log.Fatal("Can't use the elements to convert vector:", vectorStr)
		return nil, errors.New("Not enough element to convert vector")
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

	vector := Vector3D{
		X: x,
		Y: y,
		Z: z,
	}

	return &vector, nil
}
