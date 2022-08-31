package sdk

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/crypto/aes"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-lib/pkg/wallet"
)

// Check that the wallet is currently saved using the configured
// IO instance which can JS localStorage, disk, etc.
func (sdk *SDK) Check() (err error) {
	log.D("sdk", "wallet", "check")
	if !sdk.io.Check(cfg.Str("wallet.path")) {
		err = fmt.Errorf("wallet save location not found")
	}
	return
}

// Load the SDK wallet with the provided mnemonic words.  This should
// be accomplished first before using the wallet methods.
func (sdk *SDK) Load(words string) (err error) {
	log.D("sdk", "wallet", "load", words)
	var seed []byte
	if seed, err = wallet.SeedFromMnemonic(words); err != nil {
		return
	}

	if err = sdk.wallet.Load(seed); err != nil {
		return
	}

	if err = sdk.deriveRootIdentity(); err != nil {
		return
	}

	sdk.fetch.Signer(sdk.Signer)
	return
}

// Mnemonic will generate a new mnemonic word list and return it or
// an error explaining why it failed.
func (sdk *SDK) Mnemonic() (words string, err error) {
	log.D("sdk", "wallet", "mnemonic")
	words, err = wallet.NewMnemonic()
	return
}

// Open will decrypt the saved wallet seed using the provided secret
// attempt to load the wallet with the plain text from decryption.
func (sdk *SDK) Open(secret string) (err error) {
	log.D("sdk", "wallet", "open", secret)
	buf := new(bytes.Buffer)
	path := cfg.Str("wallet.path")
	log.D("sdk", "wallet", "open from", path)
	if err = sdk.io.Read(path, buf); err != nil {
		return
	}

	enc := make([]byte, hex.DecodedLen(buf.Len()))
	if _, err = hex.Decode(enc, buf.Bytes()); err != nil {
		return
	}

	var seed []byte
	log.D("sdk", "wallet", "decrypt cipher", enc)
	if seed, err = aes.Decrypt(secret, enc); err != nil {
		return
	}

	log.D("sdk", "wallet", "mnemonic from seed", string(seed))
	var words string
	if words, err = wallet.MnemonicFromSeed(seed); err != nil {
		return
	}

	log.D("sdk", "wallet", "load with mnemonic", words)
	err = sdk.Load(words)
	return
}

// Save the wallet by first encrypting the seed with the provided secret
// and then writing the cipher to the configured IO.
func (sdk *SDK) Save(secret string) (err error) {
	log.D("sdk", "wallet", "save", secret)
	seed := sdk.wallet.Seed()

	var enc []byte
	log.D("sdk", "wallet", "save", "encrypt seed", seed)
	if enc, err = aes.Encrypt(secret, seed); err != nil {
		return
	}

	buf := make([]byte, hex.EncodedLen(len(enc)))
	log.D("sdk", "wallet", "save", "encode cipher", sdk.codec, enc)
	hex.Encode(buf, enc)

	path := cfg.Str("wallet.path")
	log.D("sdk", "wallet", "save", "write to", path)

	r := bytes.NewReader(buf)
	if err = sdk.io.Write(path, r); err != nil {
		return
	}
	return
}

// Signer is a convenience method that is used by the fetching mechanism
// of the SDK to make requests to the backend API services.
func (sdk *SDK) Signer(hash []byte) (sig []byte, err error) {
	log.D("sdk", "wallet", "signer", hash)
	sig, err = ecdsa.SignASN1(rand.Reader, sdk.privateKey, hash)
	return
}
