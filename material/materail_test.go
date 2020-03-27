package material

import "testing"

func TestParseColor(t *testing.T) {
	color, err := ParseColor("123,234,212,122")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if color.R != 123 {
		t.Log("R", color.R)
		t.Fail()
	}

	if color.G != 234 {
		t.Log("G", color.G)
		t.Fail()
	}

	if color.B != 212 {
		t.Log("B", color.B)
		t.Fail()
	}

	if color.A != 122 {
		t.Log("A", color.A)
		t.Fail()
	}
}
