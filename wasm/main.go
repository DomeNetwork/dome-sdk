//go:build wasm
// +build wasm

package main

import (
	"github.com/domenetwork/dome-sdk/pkg/wasm"
)

// This is the entrypoint for compilation of the SDK into WASM.
func main() {
	wasm.Run()
}
