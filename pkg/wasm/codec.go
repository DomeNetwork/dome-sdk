//go:build wasm
// +build wasm

package wasm

import (
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

func base64() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "codec", "base64", s)

		return promisify(func() ([]any, error) {
			v, err := _sdk.Base64(s)
			if err != nil {
				return nil, err
			}

			return []any{v}, nil
		})
	})
}

func hex() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "codec", "hex", s)

		return promisify(func() ([]any, error) {
			v, err := _sdk.Hex(s)
			if err != nil {
				return nil, err
			}

			return []any{v}, nil
		})
	})
}
