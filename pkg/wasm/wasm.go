//go:build wasm
// +build wasm

package wasm

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/codec"
	"github.com/domenetwork/dome-lib/pkg/coin/ethereum"
	"github.com/domenetwork/dome-lib/pkg/io"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-sdk/sdk"
)

var _sdk *sdk.SDK

func config() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "config", jsonStr)

		return promisify(func() ([]any, error) {
			r := strings.NewReader(jsonStr)
			if err := cfg.Setup(r, "json"); err != nil {
				return nil, err
			}

			// Ethereum
			eth, err := ethereum.New()
			if err != nil {
				return nil, err
			}

			_sdk.AddCoins(eth)

			return nil, nil
		})
	})
}

// Run is the entry point for WASM and will attach the SDK into
// the "dome" global JS variable.
func Run() {
	codec := codec.NewHex()
	io := io.NewJS()

	_sdk = sdk.New(codec, io)

	dome := map[string]interface{}{
		"config": config(),

		// codec
		"base64": base64(),
		"hex":    hex(),

		// coins
		"account":     account(),
		"accounts":    accounts(),
		"balance":     balance(),
		"coins":       coins(),
		"gas":         gas(),
		"generate":    generate(),
		"send":        send(),
		"subscribe":   subscribe(),
		"unsubscribe": unsubscribe(),

		// crypto
		"decrypt":   decrypt(),
		"encrypt":   encrypt(),
		"hash":      hash(),
		"publicKey": publicKey(),

		// depot
		"download": download(),
		"metadata": metadata(),
		"upload":   upload(),

		// io
		"read":  read(),
		"write": write(),

		// nym
		"forgot":   forgot(),
		"key":      key(),
		"keys":     keys(),
		"login":    login(),
		"lookup":   lookup(),
		"register": register(),
		"reset":    reset(),
		"search":   search(),
		"update":   update(),
		"user":     user(),

		// validate
		"address":    address(),
		"password":   password(),
		"privateKey": privateKey(),
		"secret":     secret(),
		"username":   username(),
		"uuidv4":     uuidv4(),

		// wallet
		"check":    check(),
		"load":     load(),
		"mnemonic": mnemonic(),
		"open":     open(),
		"save":     save(),
	}

	js.Global().Set("dome", dome)

	fmt.Println("Running the DOME WASM!")
	select {}
}
