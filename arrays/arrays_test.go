package arrays

import (
	"reflect"
	"testing"
)

func TestArrayOperation(t *testing.T) {

	cases := []struct {
		op  string
		arr [][]float32
		res []float32
	}{
		{"sum", [][]float32{{1, 2, 3}, {10, 20, 30}}, []float32{11, 22, 33}},
		{"sum", [][]float32{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}}, []float32{111, 222, 333}},
		{"mult", [][]float32{{1, 2, 3}, {10, 20, 30}}, []float32{10, 40, 90}},
		{"mult", [][]float32{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}}, []float32{1000, 8000, 27000}},
	}
	for _, c := range cases {
		got, err := ArrayOperation(c.op, c.arr)
		want := c.res
		same := reflect.DeepEqual(got, want)

		if err != nil {
			t.Errorf("ArrayOperation returned an error: %s", err)
		} else if !same {
			t.Errorf("ArrayOperation(%s, %v) = %v, want %v", c.op, c.arr, got, want)
		} else {
			// t.Logf("ArrayOperation(%s, %v) = %v", c.op, c.arr, got)
		}
	}
}
