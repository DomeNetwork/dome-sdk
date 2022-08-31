//go:build wasm
// +build wasm

package wasm

import (
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

func promisify(fn func() ([]any, error)) any {
	promise := js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]

		go func() {
			results, err := fn()
			if err != nil {
				reject.Invoke(Error.New(err))
				return
			}
			if results == nil {
				results = []any{}
			}
			log.D("wasm", "promisify", "results", results)

			resolve.Invoke(results...)
		}()

		return nil
	})
	return Promise.New(promise)
}
