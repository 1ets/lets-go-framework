package controllers

import (
	"crypto"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"
)

// Create new RSA key pair and save to storage.
func GenerateKeys() (response data.ResponseExample, err error) {
	rsaKeys := &lets.RsaKeys{
		PublicKeyFile:  "keys/public.pem",  // Save path
		PrivateKeyFile: "keys/private.pem", // Save path
	}

	var myCrypto lets.Crypto
	err = myCrypto.Generate(lets.RSA4096, rsaKeys)

	if err != nil {
		lets.LogE("GenerateKeys: %w", err)
		return
	}

	rsaKeys.Save()
	return
}

// Demonstrate how to using encryption and decryption.
func EncryptDecrypt() (err error) {
	// Load the private key and/or public key, depending on needs.
	// PrivateKey is for encryption and PublicKey is for decryption.
	// We need to load two keys for demonstrate all process.
	rsaKeys := &lets.RsaKeys{
		PrivateKeyFile: "keys/private.pem",
		PublicKeyFile:  "keys/public.pem",
	}
	if err = rsaKeys.Load(); err != nil {
		return
	}

	// Assign key in to crypto lib.
	myCrypto := lets.Crypto{
		Rsa: rsaKeys,
	}

	////////////////////////// ENCRYPTION //////////////////////////
	// 1. Create a type of encryption with message.
	myEncrypt := &lets.RsaEncrypt{
		Hash:       crypto.SHA256,
		Encryption: lets.PKCS1v15,
		Message:    []byte("This is Lets GO Framework"),
	}

	// 2. Do Encryption.
	if err = myCrypto.Encrypt(myEncrypt); err != nil {
		return
	}

	// 3. Result.
	cipher := myEncrypt.ToBase64()
	lets.LogI("Encrypted: %s", cipher)

	////////////////////////// DECRYPTION //////////////////////////
	// 1. Decryption setup.
	myDecrypt := &lets.RsaEncrypt{
		Hash:       crypto.SHA256,
		Encryption: lets.PKCS1v15,
	}

	// 2. Load cipher text.
	myDecrypt.FromBase64(cipher)

	// 3. Do Decryption
	if err = myCrypto.Decrypt(myDecrypt); err != nil {
		return
	}

	// 4. Result
	lets.LogI("Decrypted: %s", string(myDecrypt.Message))

	return
}

// Attempt to create signature.
func CreateSignature() (signature string, err error) {
	// 1. Load the provate key.
	rsaKeys := &lets.RsaKeys{
		PrivateKeyFile: "keys/private.pem",
	}
	if err = rsaKeys.LoadPrivateKey(); err != nil {
		return
	}

	// 2. Assign key in to crypto lib.
	myCrypto := lets.Crypto{
		Rsa: rsaKeys,
	}

	// 3. Set the payload that want to sign.
	sign := lets.RsaSign{
		Hash:    crypto.SHA256,
		Payload: []byte("This is Lets GO Framework"),
	}

	// 4. Create signature.
	if err = myCrypto.Sign(&sign); err != nil {
		return
	}

	// 5. Get the result in string base64 format.
	signature = sign.ToBase64()

	return
}

// Attempt to create signature.
func VerifySignature(signature string) (result string, err error) {
	// 1. Load the provate key.
	rsaKeys := &lets.RsaKeys{
		PublicKeyFile: "keys/public.pem",
	}
	if err = rsaKeys.LoadPublicKey(); err != nil {
		return
	}

	// 2. Assign key in to crypto lib.
	myCrypto := lets.Crypto{
		Rsa: rsaKeys,
	}

	// 3. Set the payload that want to verify.
	sign := lets.RsaSign{
		Hash:    crypto.SHA256,
		Payload: []byte("This is Lets GO Framework"),
	}

	// 4. Load signature
	if err = sign.FromBase64(signature); err != nil {
		return
	}

	// 4. Verify
	if err = myCrypto.VerifySign(&sign); err != nil {
		result = "Rejected"
		return
	}

	result = "Verified"
	return
}
