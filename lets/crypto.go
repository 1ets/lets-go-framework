package lets

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"hash"
)

const (
	RSA512 = iota
	RSA1024
	RSA2048
	RSA4096
)

func getkeylength(keyType int) int {
	switch keyType {
	case RSA512:
		return 512
	case RSA1024:
		return 1024
	case RSA2048:
		return 2048
	case RSA4096:
		return 4096
	default:
		return 2048
	}
}

// Crypto structure
type Crypto struct {
	Rsa *RsaKeys
}

// Sign payload and get signature.
func (c *Crypto) Sign(sign *RsaSign) (err error) {
	hash := sha256.New()
	if _, err = hash.Write(sign.Payload); err != nil {
		return
	}

	sum := hash.Sum(nil)
	sign.Signature, err = rsa.SignPKCS1v15(rand.Reader, c.Rsa.PrivateKey, sign.Hash, sum)
	if err != nil {
		return
	}

	return
}

// Verify signature and match with payload.
func (c *Crypto) VerifySign(sign *RsaSign) (err error) {
	hash := sha256.New()
	if _, err = hash.Write(sign.Payload); err != nil {
		return
	}

	sum := hash.Sum(nil)
	return rsa.VerifyPKCS1v15(c.Rsa.PublicKey, sign.Hash, sum, sign.Signature)
}

// Generate a pair of RSA keys.
func (c *Crypto) Generate(keyType int, rsaKeys *RsaKeys) (err error) {
	rsaKeys.PrivateKey, err = rsa.GenerateKey(rand.Reader, getkeylength(keyType))
	if err != nil {
		return
	}

	rsaKeys.PublicKey = &rsaKeys.PrivateKey.PublicKey
	return
}

func (c *Crypto) Encrypt(encryption *RsaEncrypt) (err error) {
	if encryption.Encryption == OAEP {

		// Generate hash
		var hash hash.Hash
		if encryption.Hash == crypto.SHA256 {
			hash = sha256.New()
		}

		encryption.Encrypted, err = rsa.EncryptOAEP(hash, rand.Reader, c.Rsa.PublicKey, encryption.Message, nil)
		return
	}

	encryption.Encrypted, err = rsa.EncryptPKCS1v15(rand.Reader, c.Rsa.PublicKey, encryption.Message)
	return
}

// Decrypt encrypted message to plain text.
func (c *Crypto) Decrypt(encryption *RsaEncrypt) (err error) {
	if encryption.Encryption == OAEP {
		opts := &rsa.OAEPOptions{Hash: encryption.Hash}

		encryption.Message, err = c.Rsa.PrivateKey.Decrypt(nil, encryption.Encrypted, opts)
		return
	}

	opts := &rsa.PKCS1v15DecryptOptions{}
	encryption.Message, err = c.Rsa.PrivateKey.Decrypt(nil, encryption.Encrypted, opts)

	return
}
