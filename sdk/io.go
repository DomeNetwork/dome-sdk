package sdk

import (
	"io"

	"github.com/domenetwork/dome-lib/pkg/log"
)

// Read using the SDK IO from the given key into the provided writer.
func (sdk *SDK) Read(key string, w io.Writer) (err error) {
	log.D("sdk", "io", "read", key)
	err = sdk.io.Read(key, w)
	return
}

// Write the contents of the reader using the SDK IO to the given key.
func (sdk *SDK) Write(key string, r io.Reader) (err error) {
	log.D("sdk", "io", "write", key)
	err = sdk.io.Write(key, r)
	return
}
