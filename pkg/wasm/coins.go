//go:build wasm
// +build wasm

package wasm

import (
	"crypto/ecdsa"
	enchex "encoding/hex"
	"fmt"
	"math/big"
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/ethereum/go-ethereum/crypto"
)

func account() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "account", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			acct := coin.GetAccount()
			tmp := map[string]interface{}{
				"address": coin.GetAddress(acct),
				"balance": acct.Balance.String(),
				"path":    acct.Path.String(),
			}
			return []any{tmp}, nil
		})
	})
}

func accounts() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "accounts", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			accts := coin.GetAccounts()
			tmp := make([]interface{}, len(accts))
			for i, acct := range accts {
				tmp[i] = map[string]interface{}{
					"address": coin.GetAddress(acct),
					"balance": acct.Balance.String(),
					"path":    acct.Path.String(),
				}
			}
			return []any{tmp}, nil
		})
	})
}

func balance() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "balance", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			balance, err := _sdk.Balance(coin)
			if err != nil {
				return nil, err
			}

			return []any{balance.Int64()}, nil
		})
	})
}

func coins() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		log.D("wasm", "coins", "coins")

		return promisify(func() ([]any, error) {
			coins, err := _sdk.Coins()
			if err != nil {
				return nil, err
			}

			tmp := make([]interface{}, len(coins))
			for i, coin := range coins {
				tmp[i] = map[string]interface{}{
					"name":   coin.GetName(),
					"symbol": coin.GetSymbol(),
				}
			}
			return []any{tmp}, nil
		})
	})
}

func gas() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "gas", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			gas, err := _sdk.Gas(coin)
			if err != nil {
				return nil, err
			}

			return []any{gas.Int64()}, nil
		})
	})
}

func generate() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "generate", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			acct, err := _sdk.Generate(coin)
			if err != nil {
				return nil, err
			}

			prv, err := acct.PrivateKey()
			if err != nil {
				return nil, err
			}

			pub := prv.Public().(*ecdsa.PublicKey)
			pubBytes := crypto.FromECDSAPub(pub)
			pubHex := enchex.EncodeToString(pubBytes)

			tmp := map[string]interface{}{
				"address":   coin.GetAddress(acct),
				"balance":   acct.Balance.String(),
				"path":      acct.Path.String(),
				"publicKey": pubHex,
			}
			return []any{tmp}, nil
		})
	})
}

func send() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		to := args[1].String()
		amountStr := args[2].String()
		data := args[3].String()
		log.D("wasm", "coins", "send", coinName, to, amountStr, data)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			var ok bool
			amount := big.NewInt(0)
			if amount, ok = amount.SetString(amountStr, 10); !ok {
				return nil, fmt.Errorf("unable to parse amount `%s`", amountStr)
			}
			log.D("wasm", "send", "amount", amount.String())

			tx, err := _sdk.Send(coin, to, amount, []byte(data))
			if err != nil {
				return nil, err
			}

			return []any{tx}, nil
		})
	})
}

func subscribe() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		callback := args[1]
		log.D("wasm", "coins", "subscribe", coinName, callback)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			sub := make(chan interface{})
			if err = coin.Subscribe(sub); err != nil {
				return nil, err
			}

			go func() {
				for v := range sub {
					callback.Invoke(v.(map[string]interface{}))
				}
			}()

			return []any{}, nil
		})
	})
}

func unsubscribe() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		coinName := args[0].String()
		log.D("wasm", "coins", "unsubscribe", coinName)

		return promisify(func() ([]any, error) {
			coin, err := _sdk.Coin(coinName)
			if err != nil {
				return nil, err
			}

			if err = coin.Unsubscribe(); err != nil {
				return nil, err
			}

			return []any{}, nil
		})
	})
}
