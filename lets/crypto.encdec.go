package lets

import (
	"crypto"
	"encoding/base64"
)

type EncryptWith int

const (
	PKCS1v15 EncryptWith = iota
	OAEP
)

// Encryption / decryption setting
type RsaEncrypt struct {
	Hash       crypto.Hash
	Encryption EncryptWith
	Message    []byte
	Encrypted  []byte
}

// Get encrypted to base64.
func (r *RsaEncrypt) ToBase64() string {
	return base64.StdEncoding.EncodeToString(r.Encrypted)
}

// Get encrypted from base64.
func (r *RsaEncrypt) FromBase64(encrypted string) (err error) {
	r.Encrypted, err = base64.StdEncoding.DecodeString(encrypted)
	return
}
