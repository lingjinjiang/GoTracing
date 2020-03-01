package geometry

type Ray struct {
	Endpoint Point3D
	Direction Vector3D
}

func NewRay(endpoint Point3D, direction Vector3D) *Ray {
	ray := Ray {}
	ray.Endpoint = endpoint
	ray.Direction = direction
	return &ray
}