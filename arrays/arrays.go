package arrays

import "errors"

func ArrayOperation(op string, arr [][]float32) ([]float32, error) {

	A := len(arr)
	if A < 2 {
		return nil, errors.New("Must have: len(arr)>=2")
	}

	B := len(arr[0])
	for i := 1; i < A; i++ {
		if len(arr[i]) != B {
			return nil, errors.New("Must have: all arrays have same length")
		}
	}

	validOp := false
	for _, e := range [2]string{"sum", "mult"} {
		if op == e {
			validOp = true
		}
	}
	if !validOp {
		return nil, errors.New("Must have: op=sum or op=sum, op=mult")
	}

	var i, j int

	var out = make([]float32, B)

	for j = 0; j < B; j++ {
		var s float32
		if op == "sum" {
			s = 0
		} else if op == "mult" {
			s = 1
		}

		for i = 0; i < A; i++ {
			if op == "sum" {
				s += arr[i][j]
			} else if op == "mult" {
				s *= arr[i][j]
			}
		}
		out[j] = s
	}
	return out, nil
}
