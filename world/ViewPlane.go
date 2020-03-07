package world

type ViewPlane struct {
	Height            float64
	Width             float64
	s                 float64
	Samples           int
	gamma             float64
	inv_gamma         float64
	show_out_of_gamut bool
}
