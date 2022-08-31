package sdk

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"

	"github.com/domenetwork/dome-lib/pkg/log"
)

// Base64 is a convenience method that takes in the provided string and returns
// its base64 encoded equivalent.
func (sdk *SDK) Base64(s string) (v string, err error) {
	log.D("sdk", "codec", "base64", s)
	buf := new(bytes.Buffer)
	if _, err = base64.NewEncoder(base64.StdEncoding, buf).Write([]byte(s)); err != nil {
		return
	}

	v = buf.String()
	return
}

// Hex is a convenience method that encode and return the provided string in
// hexadecimal format.
func (sdk *SDK) Hex(s string) (v string, err error) {
	log.D("sdk", "codec", "hex", s)
	buf := new(bytes.Buffer)
	if _, err = hex.NewEncoder(buf).Write([]byte(s)); err != nil {
		return
	}

	v = buf.String()
	return
}
