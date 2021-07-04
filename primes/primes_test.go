package primes

import (
	"reflect"
	"testing"
)

func TestCalcPrimes(t *testing.T) {

	cases := []struct {
		min    int
		max    int
		primes []int
	}{
		{2, 20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{15, 25, []int{17, 19, 23}},
		{100, 110, []int{101, 103, 107, 109}},
	}
	for _, c := range cases {
		got, err := CalcPrimes(c.min, c.max)
		want := c.primes
		same := reflect.DeepEqual(got, want)

		if err != nil {
			t.Errorf("CalcPrimes returned an error: %s", err)
		} else if !same {
			t.Errorf("CalcPrimes(%d, %d) == %v, want %v", c.min, c.max, got, want)
		} else {
			// t.Logf("CalcPrimes(%d, %d) == %v", c.min, c.max, got)
		}
	}
}

func TestCalcPrimeFactors(t *testing.T) {

	cases := []struct {
		nb      int
		factors map[int]int
	}{
		{2, map[int]int{2: 1}},
		{3, map[int]int{3: 1}},
		{5, map[int]int{5: 1}},
		{8, map[int]int{2: 3}},
		{14, map[int]int{2: 1, 7: 1}},
		{15, map[int]int{3: 1, 5: 1}},
		{27, map[int]int{3: 3}},
		{99, map[int]int{3: 2, 11: 1}},
		{1605240, map[int]int{2: 3, 3: 2, 5: 1, 7: 3, 13: 1}},
	}
	for _, c := range cases {
		got := CalcPrimeFactors(c.nb)
		want := c.factors
		same := reflect.DeepEqual(got, want)

		if !same {
			t.Errorf("CalcPrimeFactors(%d) = %v, want %v", c.nb, got, want)
		} else {
			t.Logf("CalcPrimeFactors(%d) = %v", c.nb, got)
		}
	}

}
