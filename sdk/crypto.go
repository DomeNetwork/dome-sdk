package sdk

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/domenetwork/dome-lib/pkg/crypto/aes"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/ethereum/go-ethereum/crypto"
)

// Hash will take the provided string and return the SHA256 hash of it
// in hexadecimal format.
func (sdk *SDK) Hash(s string) (h string, err error) {
	log.D("sdk", "crypto", "hash", s)
	b := sha256.Sum256([]byte(s))
	h = hex.EncodeToString(b[:])
	return
}

// Encrypt uses the provides secret to encrypt the given plain text using AES256.
// The cipher or an error will be returned.
func (sdk *SDK) Encrypt(secret, plain string) (cipher string, err error) {
	log.D("sdk", "crypto", "encrypt", secret, plain)
	var encrypted []byte
	if encrypted, err = aes.Encrypt(secret, []byte(plain)); err != nil {
		return
	}

	cipher = hex.EncodeToString(encrypted)
	return
}

// Decrypt the provided cipher with the provided secret and return the plain text
// version or an error.
func (sdk *SDK) Decrypt(secret, cipher string) (plain string, err error) {
	log.D("sdk", "crypto", "decrypt", secret, cipher)
	var b []byte
	if b, err = hex.DecodeString(cipher); err != nil {
		return
	}

	var decrypted []byte
	if decrypted, err = aes.Decrypt(secret, b); err != nil {
		return
	}

	plain = string(decrypted)
	return
}

// PublicKey returns the current SDK signer ECDSA public key.  This public key
// is used in the signing of requested and should be the assigned public key of
// the user.  The APIs will use this public key to verify signatures of request bodies.
func (sdk *SDK) PublicKey() (publicKeyHex string, err error) {
	log.D("sdk", "crypto", "public key")
	d := crypto.FromECDSAPub(&sdk.privateKey.PublicKey)
	publicKeyHex = hex.EncodeToString(d)
	return
}
