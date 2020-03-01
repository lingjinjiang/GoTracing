package geometry

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func (src Point3D) Sub(dest Point3D) Vector3D {
	return Vector3D {
		X: dest.X - src.X,
		Y: dest.Y - src.Y,
		Z: dest.Z - src.Z,
	}
}

func (this Point3D) Add(v Vector3D) Point3D {
	return Point3D {
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	point := &Point3D {
		X: x,
		Y: y,
		Z: z,
	}
	return point
}