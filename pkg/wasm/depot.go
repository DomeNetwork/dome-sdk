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

func download() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "depot", "download", jsonStr)

		return promisify(func() ([]any, error) {
			blob := &common.Blob{}
			r := strings.NewReader(jsonStr)
			if err := json.NewDecoder(r).Decode(blob); err != nil {
				return nil, err
			}

			if err := _sdk.Download(blob); err != nil {
				return nil, err
			}

			tmp := map[string]interface{}{
				"data":     blob.Data,
				"metadata": blob.Metadata,
				"name":     blob.Name,
			}
			return []any{tmp}, nil
		})
	})
}

func metadata() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "depot", "metadata", jsonStr)

		return promisify(func() ([]any, error) {
			blob := &common.Blob{}
			r := strings.NewReader(jsonStr)
			if err := json.NewDecoder(r).Decode(blob); err != nil {
				return nil, err
			}

			if err := _sdk.Metadata(blob); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}

func upload() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		jsonStr := args[0].String()
		log.D("wasm", "depot", "upload", jsonStr)

		return promisify(func() ([]any, error) {
			blob := &common.Blob{}
			r := strings.NewReader(jsonStr)
			if err := json.NewDecoder(r).Decode(blob); err != nil {
				return nil, err
			}

			if err := _sdk.Upload(blob); err != nil {
				return nil, err
			}

			return nil, nil
		})
	})
}
