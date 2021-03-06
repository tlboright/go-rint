package rint_test

import (
	"github.com/tlboright/go-rint"
	"testing"
)

func checkInRange(num, start, end int) bool {
	return num >= start && num <= end
}

func TestGen(t *testing.T) {
	rint.Init()
	cases := []struct {
		tLimit int
	}{
		{
			tLimit: 10,
		},
		{
			tLimit: 0,
		},
	}

	for _, c := range cases {
		actual := rint.Gen(c.tLimit)
		if !checkInRange(actual, 0, c.tLimit) {
			t.Errorf("The Gen test has generated an illegal value outside of the range 0, 9.  tLimit: %d, actual: %d", c.tLimit, actual)
		}
	}
}

func TestGenRange(t *testing.T) {
	rint.Init()
	cases := []struct {
		tMin, tMax int
	}{
		{
			tMin: 5,
			tMax: 6,
		},
		{
			tMin: 0,
			tMax: 10,
		},
		{
			tMin: 4,
			tMax: 10,
		},
		{
			tMin: 4000,
			tMax: 5000,
		},
		{
			tMin: 0,
			tMax: 0,
		},
	}

	for _, c := range cases {
		if !checkInRange(rint.GenRange(c.tMin, c.tMax), c.tMin, c.tMax) {
			t.Errorf("The GenRange test has generated a value outside of the range. Range: Min: %d, Max: %d", c.tMin, c.tMax)
		}
	}
}
