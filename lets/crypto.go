package lets

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

// Crypto structure
type Crypto struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	Payload    []byte
	Signature  []byte
	Error      error
}

// Set the private key file.
func (c *Crypto) SetPrivateKeyFile(path string) {
	privateKey, err := os.ReadFile(path)
	if err != nil {
		LogE("SetPrivateKeyFile: %s", err.Error())
		c.Error = err

		return
	}

	c.ParsePrivateKey(privateKey)
}

// Set the private key string.
func (c *Crypto) SetPrivateKeyString(privateKey string) {
	c.ParsePrivateKey([]byte(privateKey))
}

// Parses a PEM encoded private key.
func (c *Crypto) ParsePrivateKey(pemBytes []byte) {
	if c.Error != nil {
		return
	}

	var err error
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		LogE("ParsePrivateKey: %s", "Private key not found.")
		c.Error = errors.New("private key not found")

		return
	}

	c.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		LogE("ParsePrivateKey: %s", err.Error())
		c.Error = err

		return
	}
}

// Set the public key file.
func (c *Crypto) SetPublicKeyFile(path string) {
	publicKey, err := os.ReadFile(path)
	if err != nil {
		LogE("SetPublicKeyFile: %s", err.Error())
		c.Error = err

		return
	}

	c.ParsePublicKey(publicKey)
}

// Set the public key string.
func (c *Crypto) SetPublicKeyString(publicKey string) {
	c.ParsePublicKey([]byte(publicKey))
}

// Parses a PEM encoded private key.
func (c *Crypto) ParsePublicKey(pemBytes []byte) {
	if c.Error != nil {
		return
	}

	var err error
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		LogE("ParsePublicKey: %s", "PublicKey: not found.")
		c.Error = err

		return
	}

	var key interface{}
	key, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		LogE("ParsePublicKey: %s", err.Error())
		c.Error = err

		return
	}

	switch keyType := key.(type) {
	case *rsa.PublicKey:
		c.PublicKey = keyType
	default:
		LogE("ParsePublicKey: %s", "Invalid type key")
		c.Error = err

	}
}

// Setup payload from []byte.
func (c *Crypto) SetPayload(payload []byte) {
	c.Payload = payload
}

// Setup payload from string.
func (c *Crypto) SetPayloadString(payload string) {
	c.SetPayload([]byte(payload))
}

// Process create signature.
func (c *Crypto) CreateSignatureSHA256WithRSA() {
	if c.Error != nil {
		return
	}

	h := sha256.New()
	h.Write(c.Payload)
	d := h.Sum(nil)

	var err error
	c.Signature, err = rsa.SignPKCS1v15(rand.Reader, c.PrivateKey, crypto.SHA256, d)
	if err != nil {
		LogE("CreateSignature: %s", "Failed to create signature")
		c.Error = err

		return
	}
}

// Get signature as []byte.
func (c *Crypto) GetSignature() []byte {
	if c.Error != nil {
		return []byte{}
	}

	return c.Signature
}

// Get signature as base64.
func (c *Crypto) GetSignatureBase64() string {
	if c.Error != nil {
		return ""
	}

	signature := base64.StdEncoding.EncodeToString(c.Signature)
	return signature
}

// Set []byte signature.
func (c *Crypto) SetSignature(signature []byte) {
	c.Signature = signature
}

// Set base64 signature.
func (c *Crypto) SetSignatureBase64(signature string) {
	if c.Error != nil {
		return
	}

	var err error
	c.Signature, err = base64.StdEncoding.DecodeString(signature)
	if err != nil {
		LogE("SetSignatureBase64: %s", "Failed to decode signature Base64")
		c.Error = err

		return
	}
}

// Verify SHA 256 with RSA signature and payload.
func (c *Crypto) VerifySignatureSHA256WithRSA() error {
	if c.Error != nil {
		return c.Error
	}

	h := sha256.New()
	h.Write(c.Payload)
	d := h.Sum(nil)

	return rsa.VerifyPKCS1v15(c.PublicKey, crypto.SHA256, d, c.Signature)
}
