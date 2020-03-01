package geometry

import (
	"math"
)

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (this Vector3D) Normalize() Vector3D {
	length := this.Length()
	return Vector3D {
		X: this.X / length,
		Y: this.Y / length,
		Z: this.Z / length,
	}
}

func (this Vector3D) Length() float64 {
	return math.Sqrt(this.X * this.X + this.Y * this.Y + this.Z * this.Z)
}

func (this Vector3D) Add(v Vector3D) Vector3D {
	return Vector3D {
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

func (this Vector3D) Sub(v Vector3D) Vector3D {
	return Vector3D {
		X: this.X - v.X,
		Y: this.Y - v.Y,
		Z: this.Z - v.Z,
	}
}

func (left Vector3D) Dot(right Vector3D) float64 {
	return left.X * right.X + left.Y * right.Y + left.Z * right.Z
}

func (this Vector3D) Opposite() Vector3D {
	return Vector3D {
		X: -this.X,
		Y: -this.Y,
		Z: -this.Z,
	}
}