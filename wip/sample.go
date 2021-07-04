package main

import (
	"fmt"

	"github.com/oscar6echo/tinygo-wasm-demo-2/arrays"
	"github.com/oscar6echo/tinygo-wasm-demo-2/primes"
)

func main() {
	fmt.Println("start main - sample calls from primes and arrays sub packages")

	min := 10
	max := 100
	primes, err := primes.CalcPrimes(min, max)
	if err != nil {
		fmt.Println("CalcPrimes crashed: ", err)
		return
	}
	fmt.Printf("primes from %d to %d = %v\n", min, max, primes)

	P := len(primes)
	var primesFloat = make([]float32, P)
	for i := 0; i < P; i++ {
		primesFloat[i] = float32(primes[i])
	}

	primesSquared, err := arrays.ArrayOperation("mult", [][]float32{primesFloat, primesFloat})
	if err != nil {
		fmt.Println("ArrayOperation crashed: ", err)
		return
	}
	fmt.Printf("primes from %d to %d squared = %v\n", min, max, primesSquared)

	fmt.Println("end main")
}
