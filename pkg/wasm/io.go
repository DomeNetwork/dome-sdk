//go:build wasm
// +build wasm

package wasm

import (
	"bytes"
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

func read() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		key := args[0].String()
		log.D("wasm", "io", "read", key)

		return promisify(func() ([]any, error) {
			w := bytes.NewBufferString("")
			if err := _sdk.Read(key, w); err != nil {
				return nil, err
			}

			return []any{w.String()}, nil
		})
	})
}

func write() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		key := args[0].String()
		value := args[1].String()
		log.D("wasm", "io", "write", key, value)

		return promisify(func() ([]any, error) {
			r := bytes.NewBufferString(value)
			if err := _sdk.Write(key, r); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}
