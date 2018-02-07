package mandelbrot

import "testing"

func TestIsMandelbrot(t *testing.T) {
	cases := []struct {
		answer, expected bool
		point            complex128
	}{
		//basic values of z that are in the set
		{expected: true, point: 0 + 0i},
		{expected: true, point: -1 + 0i},
		{expected: true, point: 0 + 1i},
		{expected: true, point: -2 + 0i},
		//basic values of z that are not in the set, the corners
		{expected: false, point: -2 + 1i},
		{expected: false, point: 1 + 1i},
		{expected: false, point: 1 - 1i},
		{expected: false, point: -2 - 1i},
		//testing floats
		{expected: true, point: -.2 - .1i},
		{expected: true, point: 0 + .1i},
		{expected: true, point: -1 + .0001i},
		{expected: true, point: -.2 - .2i},
		//zooming in and starting to look for detail
		{expected: true, point: -.75 + 0i},
		{expected: false, point: -.784 + .344i},
		{expected: false, point: -.784 + .275i},
		{expected: false, point: -.767 + .189i},
		{expected: true, point: -.802 + .12i},
	}
	for _, c := range cases {
		c.answer = IsMandelbrot(c.point)
		if c.answer != c.expected {
			t.Errorf("IsMandelbrot(%f) == %t, expected %t", c.point, c.answer, c.expected)
		}
	}
}
