//go:build wasm
// +build wasm

package wasm

import (
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

func decrypt() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		secret := args[0].String()
		cipher := args[1].String()
		log.D("wasm", "crypto", "decrypt", secret, cipher)

		return promisify(func() ([]any, error) {
			plain, err := _sdk.Decrypt(secret, cipher)
			if err != nil {
				return nil, err
			}

			return []any{plain}, nil
		})
	})
}

func encrypt() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		secret := args[0].String()
		plain := args[1].String()
		log.D("wasm", "crypto", "encrypt", secret, plain)

		return promisify(func() ([]any, error) {
			cipher, err := _sdk.Encrypt(secret, plain)
			if err != nil {
				return nil, err
			}

			return []any{cipher}, nil
		})
	})
}

func hash() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "crypto", "hash", s)

		return promisify(func() ([]any, error) {
			v, err := _sdk.Hash(s)
			if err != nil {
				return nil, err
			}

			return []any{v}, nil
		})
	})
}

func publicKey() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "crypto", "publicKey")

		return promisify(func() ([]any, error) {
			pubHex, err := _sdk.PublicKey()
			if err != nil {
				return nil, err
			}

			return []any{pubHex}, nil
		})
	})
}
