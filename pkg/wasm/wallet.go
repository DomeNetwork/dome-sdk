//go:build wasm
// +build wasm

package wasm

import (
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

func check() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "check")
		return promisify(func() ([]any, error) {
			err := _sdk.Check()
			found := err == nil
			return []any{found}, nil
		})
	})
}

func load() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		seed := args[0].String()
		log.D("wasm", "wallet", "new", words)

		return promisify(func() ([]any, error) {
			return nil, _sdk.Load(seed)
		})
	})
}

func mnemonic() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "wallet", "mnemonic")

		return promisify(func() ([]any, error) {
			words, err := _sdk.Mnemonic()
			if err != nil {
				return nil, err
			}
			return []any{words}, nil
		})
	})
}

func open() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		secret := args[0].String()
		log.D("wasm", "wallet", "open", secret)

		return promisify(func() ([]any, error) {
			return nil, _sdk.Open(secret)
		})
	})
}

func save() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		secret := args[0].String()
		log.D("wasm", "wallet", "save", secret)

		return promisify(func() ([]any, error) {
			return nil, _sdk.Save(secret)
		})
	})
}
