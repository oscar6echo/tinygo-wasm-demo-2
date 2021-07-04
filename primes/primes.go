package primes

import (
	"errors"
	"math"
)

func CalcPrimes(a, b int) ([]int, error) {
	if a < 2 || a >= b {
		return nil, errors.New("Must have: a>=2 and b>a")
	}

	var primes = make([]int, 0)

	for a <= b {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
			if a%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, a)

		}
		a++
	}
	return primes, nil

}

func CalcPrimeFactors(n int) map[int]int {
	var pf = make(map[int]int)

	// Get the number of 2s that divide n
	for n%2 == 0 {
		if _, ok := pf[2]; ok {
			pf[2] += 1
		} else {
			pf[2] = 1
		}
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			if _, ok := pf[i]; ok {
				pf[i] += 1
			} else {
				pf[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pf[n] = 1
	}
	return pf
}
