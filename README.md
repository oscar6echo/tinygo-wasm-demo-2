# TinyGo to WASM Demo 2

Follow up from repo [github.com/oscar6echo/tinygo-wasm-demo](https://github.com/oscar6echo/tinygo-wasm-demo).

- [Go package syscall/js doc](https://pkg.go.dev/syscall/js)
- ObservableHQ notebook [WASM from TinyGo - demo 2](https://observablehq.com/@oscar6echo/wasm-from-tinygo-demo-2)

# Overview

This repo contains to sample packages:

- [primes](./primes)
- [arrays](./arrays)

The `main.go` script is the source to be compiled to `html/main_js.wasm`.  
It essentially exposes functions from Go packages (primes, arrays), and local functions too.  
Most of the work is about converting:

- incoming JavaScript variable to Go types
- outgoing Go variable to JavaScript types

This is done with the `js.Value` type in [package syscall/js](https://pkg.go.dev/syscall/js).

_NOTE_: There is a persistent error: `wasm_exec.js:305 syscall/js.finalizeRef not implemented`. It seems a symptom of a potential memory leak - discussed in [TinyGo issue #1140](https://github.com/tinygo-org/tinygo/issues/1140).

## Test

```bash
# package primes
cd primes
go test -v

# package arrays
cd arrays
go test -v
```

## Build

```bash
# compile - tinygo
tinygo build -o ./html/main_js.wasm -target wasm ./main_js.go

# boiler plate
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js ./html
mv ./html/wasm_exec.js ./html/wasm_exec_mod.js

# edit wasm_exec_mod.js
# tiny adjustment: change IIFE to module
# i.e. :
# change (() => {FUNC_BODY})() to FUNC_BODY
# change global.Go = class{...} to export const Go = class{...}
```

## Serve

Serve folder `/html`:

```bash
go run serve/serve.go
```

Open <http://localhost:8082>
