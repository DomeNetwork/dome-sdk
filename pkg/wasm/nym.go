//go:build wasm
// +build wasm

package wasm

import (
	"encoding/json"
	"strings"
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/domenetwork/dome-lib/pkg/log"
)

func forgot() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "forgot", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			otp, err := _sdk.Forgot(user)
			if err != nil {
				return nil, err
			}

			return []any{otp}, nil
		})
	})
}

func key() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "key", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			key := &common.Key{}
			if err := json.NewDecoder(r).Decode(key); err != nil {
				return nil, err
			}

			if err := _sdk.Key(key); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}

func keys() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "nym", "keys")

		return promisify(func() ([]any, error) {
			keys, err := _sdk.Keys()
			if err != nil {
				return nil, err
			}

			tmp := make([]any, len(keys))
			for i, key := range keys {
				tmp[i] = map[string]interface{}{
					"createdAt": key.CreatedAt.Format("2006-01-02T15:04:05.999999999Z"),
					"name":      key.Name,
					"publicKey": key.PublicKey,
					"used":      key.Used,
				}
			}
			return []any{tmp}, nil
		})
	})
}

func login() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "login", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			if err := _sdk.Login(user); err != nil {
				return nil, err
			}

			return []any{}, nil
		})
	})
}

func lookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "lookup", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			if err := _sdk.Lookup(user); err != nil {
				return nil, err
			}

			return []any{user.PublicKey}, nil
		})
	})
}

func register() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "register", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			if err := _sdk.Register(user); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}

func reset() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		otp := args[0].String()
		jsonStr := args[1].String()
		log.D("wasm", "nym", "reset", otp, jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			if err := _sdk.Reset(otp, user); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}

func search() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "search", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			search := &common.Search{}
			if err := json.NewDecoder(r).Decode(search); err != nil {
				return nil, err
			}

			users, err := _sdk.Search(search.Term)
			if err != nil {
				return nil, err
			}

			tmp := make([]any, len(users))
			for i, user := range users {
				tmp[i] = user.Username
			}
			return []any{tmp}, nil
		})
	})
}

func update() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "nym", "update", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			user := &common.User{}
			if err := json.NewDecoder(r).Decode(user); err != nil {
				return nil, err
			}

			if err := _sdk.Update(user); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}

func user() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "nym", "user")
		return promisify(func() ([]any, error) {
			user, err := _sdk.User()
			if err != nil {
				return nil, err
			}

			return []any{user}, nil
		})
	})
}
