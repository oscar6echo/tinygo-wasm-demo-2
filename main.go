package main

import (
	"strconv"
	"strings"
	"syscall/js"

	"github.com/oscar6echo/tinygo-wasm-demo-2/arrays"
	"github.com/oscar6echo/tinygo-wasm-demo-2/primes"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WASM initialized")

	js.Global().Set("splitter", js.FuncOf(splitter))
	js.Global().Set("calcPrimes", js.FuncOf(calcPrimes))
	js.Global().Set("calcPrimeFactors", js.FuncOf(calcPrimeFactors))
	js.Global().Set("arrayOperation", js.FuncOf(arrayOperation))

	js.Global().Set("wasmDoc", js.FuncOf(wasmDoc))

	<-c
}

func wasmDoc(this js.Value, args []js.Value) interface{} {
	// return name of functions set to global scope in main()
	result := make([]interface{}, 0)
	var funcNames = []string{"splitter", "calcPrimes", "calcPrimeFactors", "arrayOperation"}
	for _, e := range funcNames {
		result = append(result, e)
	}
	return js.ValueOf(result)
}

func splitter(this js.Value, args []js.Value) interface{} {
	values := strings.Split(args[0].String(), ",")

	result := make([]interface{}, 0)
	for _, each := range values {
		result = append(result, each)
	}
	return js.ValueOf(result)
}

func calcPrimes(this js.Value, args []js.Value) interface{} {
	min := args[0].Int()
	max := args[1].Int()

	res, _ := primes.CalcPrimes(min, max)

	result := make([]interface{}, 0)
	for _, e := range res {
		result = append(result, e)
	}
	return js.ValueOf(result)
}

func calcPrimeFactors(this js.Value, args []js.Value) interface{} {
	N := args[0].Int()
	res := primes.CalcPrimeFactors(N)

	result := make(map[string]interface{})
	for k, v := range res {
		s := strconv.Itoa(k)
		result[s] = v
	}

	return js.ValueOf(result)
}

func arrayOperation(this js.Value, args []js.Value) interface{} {

	op := args[0].String()
	arrIn := args[1]
	A := arrIn.Length()
	var arr = make([][]float32, A)

	for i := 0; i < A; i++ {
		sarrIn := arrIn.Index(i)
		B := sarrIn.Length()
		arr[i] = make([]float32, B)
		for j := 0; j < B; j++ {
			arr[i][j] = float32(sarrIn.Index(j).Float())
		}
	}

	res, _ := arrays.ArrayOperation(op, arr)

	result := make([]interface{}, 0)
	for _, e := range res {
		result = append(result, e)
	}
	return js.ValueOf(result)

}
