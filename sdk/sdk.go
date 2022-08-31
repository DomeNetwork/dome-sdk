package sdk

import (
	"crypto/ecdsa"

	"github.com/domenetwork/dome-lib/pkg/codec"
	"github.com/domenetwork/dome-lib/pkg/fetch"
	"github.com/domenetwork/dome-lib/pkg/io"
	"github.com/domenetwork/dome-lib/pkg/wallet"
)

// TODO: add check to make sure wallet is ready before usage.

// SDK is the interface, facade, to the entire DOME platform.
// The SDK is setup in a way that allows it be compiled into WASM
// or used directly as a package in Go.  It is also allows for
// easy integration into other programming languages using the
// C FFI.
type SDK struct {
	codec      codec.Codec
	fetch      *fetch.Client
	io         io.IO
	privateKey *ecdsa.PrivateKey
	wallet     *wallet.Wallet
}

// New will return a SDK instance with the provided Codec and IO.
// The client and wallet or instantiated as well but the wallet should
// be loaded before usage.
func New(codec codec.Codec, io io.IO) (sdk *SDK) {
	sdk = &SDK{
		codec:  codec,
		fetch:  fetch.NewClient(),
		io:     io,
		wallet: wallet.New(),
	}
	return
}

// A convenience method that allows for easy derivation of the root
// identity key that is used in Nym to authorize signed requests by
// authenticated users of the platform.
func (sdk *SDK) deriveRootIdentity() (err error) {
	var path wallet.Path
	if path, err = wallet.ParsePath("m/0'/0'/0'/0/0"); err != nil {
		return
	}

	var acct *wallet.Account
	if acct, err = sdk.wallet.Derive(path); err != nil {
		return
	}

	sdk.privateKey, err = acct.PrivateKey()
	return
}
