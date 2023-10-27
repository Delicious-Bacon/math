package algo

import (
	"math/rand"
	"testing"
)

func TestEuclid(t *testing.T) {

	type tcase struct {
		a, b, res int
	}

	cases := []tcase{
		{a: 2, b: 3, res: 1},
		{a: 3, b: 5, res: 1},
		{a: 3, b: 3, res: 3},
		{a: 24, b: 48, res: 24},
		{a: 24, b: 36, res: 12},
		{a: 64, b: 236, res: 4},
	}

	for _, c := range cases {
		if Euclid(c.a, c.b) != c.res {
			t.Errorf("Euclid(%d, %d) = %d, want %d", c.a, c.b, Euclid(c.a, c.b), c.res)
		}
	}
}

func FuzzEuclid(f *testing.F) {

	// Do fuzzy test on euclid algo

	for i := 0; i < 1000; i++ {
		f.Add(i, i+(rand.Intn(100)*(-1*rand.Intn(2))))
	}

	f.Fuzz(func(t *testing.T, a, b int) {

		var (
			gcf int
		)

		aa := a
		bb := b
		if aa < 0 {
			aa *= -1
		}
		if bb < 0 {
			bb *= -1
		}

		if aa < bb {
			gcf = aa
		} else {
			gcf = bb
		}

		for gcf > 0 && !(a%gcf == 0 && b%gcf == 0) {
			gcf--
		}

		res := Euclid(a, b)
		if res != gcf {
			t.Errorf("Euclid(%d, %d) = %d, want %d", a, b, res, gcf)
		}
	})
}
