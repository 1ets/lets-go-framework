package lets

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// Handle load and save keys to memory/storage.
type RsaKeys struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey

	PrivateKeyFile string
	PublicKeyFile  string
}

////////////
// LOADER //
////////////

// Load Keys file.
func (r *RsaKeys) Load() (err error) {
	err = r.LoadPrivateKey()
	if err != nil {
		return
	}

	err = r.LoadPublicKey()
	if err != nil {
		return
	}

	return
}

// Load file and setup private key.
func (r *RsaKeys) LoadPrivateKey() (err error) {
	privateKey, err := os.ReadFile(r.PrivateKeyFile)
	if err != nil {
		return
	}

	return r.SetPrivateKey(privateKey)
}

// Load file and setup public key.
func (r *RsaKeys) LoadPublicKey() (err error) {
	publicKey, err := os.ReadFile(r.PublicKeyFile)
	if err != nil {
		return
	}

	return r.SetPublicKey(publicKey)
}

///////////
// SAVER //
///////////

// Save all keys into storage.
func (r *RsaKeys) Save() (err error) {
	LogI("SAVED: %s", r.PrivateKeyFile)
	err = r.SavePrivateKey()
	if err != nil {
		return
	}

	err = r.SavePublicKey()
	if err != nil {
		return
	}

	return
}

// Save PrivateKey to storage.
func (r *RsaKeys) SavePrivateKey() (err error) {
	pemFile, err := os.Create(r.PrivateKeyFile)
	if err != nil {
		return
	}

	pemKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(r.PrivateKey),
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}
	LogI("SAVED: %s", r.PrivateKeyFile)
	pemFile.Close()
	return
}

// Save PublicKey to storage.
func (r *RsaKeys) SavePublicKey() (err error) {
	pemFile, err := os.Create(r.PublicKeyFile)
	if err != nil {
		return
	}

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(r.PublicKey),
	}

	err = pem.Encode(pemFile, pemKey)
	if err != nil {
		return
	}

	pemFile.Close()
	return
}

////////////
// SETUPS //
////////////

// Parses a PEM encoded private key.
func (r *RsaKeys) SetPrivateKey(privateKey []byte) (err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key not found")
		return
	}

	r.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	return
}

// Set the private key string.
func (r *RsaKeys) SetPrivateKeyString(privateKey string) (err error) {
	return r.SetPrivateKey([]byte(privateKey))
}

// Parses a PEM encoded public key.
func (r *RsaKeys) SetPublicKey(publicKey []byte) (err error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key key not found")
		return
	}

	var key interface{}
	// key, err = x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	// 	return
	// }

	key, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}

	switch keyType := key.(type) {
	case *rsa.PublicKey:
		r.PublicKey = keyType
	default:
		err = errors.New("invalid public key type")

	}

	return
}

// Set the public key string.
func (r *RsaKeys) SetPublicKeyString(publicKey string) (err error) {
	return r.SetPublicKey([]byte(publicKey))
}
