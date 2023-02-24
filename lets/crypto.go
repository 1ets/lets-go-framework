package lets

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"lets-go-framework/lets/types/crypt"
)

// Crypto structure
type Crypto struct {
	Rsa       *RsaKeys
	Payload   []byte
	Signature []byte
	Error     error
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
func (c *Crypto) GenerateSignature() (err error) {
	h := sha256.New()
	h.Write(c.Payload)
	d := h.Sum(nil)

	c.Signature, err = rsa.SignPKCS1v15(rand.Reader, c.Rsa.PrivateKey, crypto.SHA256, d)
	if err != nil {
		return
	}

	return
}

// Get signature as []byte.
func (c *Crypto) GetSignature() []byte {
	if c.Error != nil {
		return []byte{}
	}

	return c.Signature
}

// Get signature as base64.
func (c *Crypto) GetSignatureBase64() (signature string) {
	signature = base64.StdEncoding.EncodeToString(c.Signature)

	return
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

	return rsa.VerifyPKCS1v15(c.Rsa.PublicKey, crypto.SHA256, d, c.Signature)
}

const (
	RSA512 = iota
	RSA1024
	RSA2048
	RSA4096
)

// Generate a pair of RSA keys.
func (c *Crypto) Generate(keyType int, paths crypt.KeyPath) (rsaKeys *RsaKeys, err error) {
	rsaKeys = &RsaKeys{}
	rsaKeys.PrivateKey, err = rsa.GenerateKey(rand.Reader, getkeylength(keyType))
	if err != nil {
		return
	}

	rsaKeys.PublicKey = &rsaKeys.PrivateKey.PublicKey
	return
}

// func (c *Crypto) GenerateKey(keyType int) {
// 	var err error

// 	c.PrivateKey, err = rsa.GenerateKey(rand.Reader, getkeylength(keyType))
// 	if err != nil {
// 		LogE("GenerateKey: %w", err)
// 		return
// 	}
// 	c.PublicKey = &c.PrivateKey.PublicKey
// }

// func (c *Crypto) GetPrivateKey() string {
// 	var err error
// 	buf := new(bytes.Buffer)

// 	var pemKey = &pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(c.PrivateKey),
// 	}

// 	err = pem.Encode(buf, pemKey)

// 	if err != nil {
// 		LogE("GenerateKey: Save PEM: %w", err)
// 		return ""
// 	}

// 	return buf.String()
// }

// func (c *Crypto) SavePrivateKey(filename string) {
// 	// Save PEM Private file
// 	pemFile, err := os.Create(filename)

// 	if err != nil {
// 		LogE("GenerateKey: Create PEM: %w", err)
// 		return
// 	}

// 	var pemKey = &pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(c.PrivateKey),
// 	}

// 	err = pem.Encode(pemFile, pemKey)

// 	if err != nil {
// 		LogE("GenerateKey: Save PEM: %w", err)
// 		return
// 	}

// 	pemFile.Close()
// }

// func (c *Crypto) GetPublicKey() string {
// 	var err error
// 	buf := new(bytes.Buffer)

// 	var pemKey = &pem.Block{
// 		Type:  "PUBLIC KEY",
// 		Bytes: x509.MarshalPKCS1PublicKey(c.PublicKey),
// 	}

// 	err = pem.Encode(buf, pemKey)
// 	if err != nil {
// 		LogE("GenerateKey: Save PEM: %w", err)
// 		return ""
// 	}

// 	return buf.String()
// }

// func (c *Crypto) SavePublicKey(filename string) {
// 	// Save PEM Public file
// 	pemFile, err := os.Create(filename)
// 	if err != nil {
// 		LogE("GenerateKey: Create PEM: %w", err)
// 		return
// 	}

// 	var pemKey = &pem.Block{
// 		Type:  "PUBLIC KEY",
// 		Bytes: x509.MarshalPKCS1PublicKey(c.PublicKey),
// 	}

// 	err = pem.Encode(pemFile, pemKey)

// 	if err != nil {
// 		LogE("GenerateKey: Save PEM: %w", err)
// 		return
// 	}

// 	pemFile.Close()
// }

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

func (c *Crypto) EncryptOAEP(message string) (b []byte, b64 string) {
	b, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, c.Rsa.PublicKey, []byte(message), nil)
	if err != nil {
		LogE("EncryptOAEP: EncryptOAEP: %w", err)
		return
	}

	b64 = base64.StdEncoding.EncodeToString(b)

	return
}

func (c *Crypto) DecryptOAEP(b []byte) string {
	opts := &rsa.OAEPOptions{Hash: crypto.SHA256}

	decryptedBytes, err := c.Rsa.PrivateKey.Decrypt(nil, b, opts)
	if err != nil {
		LogE("DecryptOAEP: Decrypt: %w", err)
		return ""
	}

	return string(decryptedBytes)
}

func (c *Crypto) DecryptB64OAEP(b64 string) string {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		LogE("DecryptOAEP: Decrypt: %w", err)
		return ""
	}

	return c.DecryptOAEP(b)
}
