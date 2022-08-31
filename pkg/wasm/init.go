//go:build wasm
// +build wasm

package wasm

import "syscall/js"

var (
	// Makes the JS Error type available to Go.
	Error js.Value
	// Allows for the usage of JS Promises from Go.
	Promise js.Value
)

func init() {
	global := js.Global()

	Error = global.Get("Error")
	Promise = global.Get("Promise")
}
