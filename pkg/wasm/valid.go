//go:build wasm
// +build wasm

package wasm

import (
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-lib/pkg/valid"
)

func address() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "address", s)

		return promisify(func() ([]any, error) {
			if err := valid.Address(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func password() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "password", s)

		return promisify(func() ([]any, error) {
			if err := valid.Password(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func privateKey() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "privateKey", s)

		return promisify(func() ([]any, error) {
			if err := valid.PrivateKey(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func secret() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "secret", s)

		return promisify(func() ([]any, error) {
			if err := valid.Secret(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func username() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "username", s)

		return promisify(func() ([]any, error) {
			if err := valid.Username(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func uuidv4() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "uuidv4", s)

		return promisify(func() ([]any, error) {
			if err := valid.UUIDv4(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}

func words() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		s := args[0].String()
		log.D("wasm", "valid", "words", s)

		return promisify(func() ([]any, error) {
			if err := valid.Mnemonic(s); err != nil {
				return nil, err
			}
			return nil, nil
		})
	})
}
