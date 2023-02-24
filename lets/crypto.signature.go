package lets

import (
	"crypto"
	"encoding/base64"
)

// A place for create or load signature ;
type RsaSign struct {
	Hash      crypto.Hash
	Signature []byte
	Payload   []byte
}

// Get signature to base64.
func (r *RsaSign) ToBase64() string {
	return base64.StdEncoding.EncodeToString(r.Signature)
}

// Get signature as base64.
func (r *RsaSign) FromBase64(signature string) (err error) {
	r.Signature, err = base64.StdEncoding.DecodeString(signature)
	return
}
