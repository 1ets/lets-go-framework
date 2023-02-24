package controllers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types/crypt"
)

// Create new RSA key pair and save to storage.
func GenerateKeys() (response data.ResponseExample, err error) {
	var myCrypto lets.Crypto
	keys, err := myCrypto.Generate(lets.RSA4096, crypt.KeyPath{
		PublicKey:  "keys/public.pem",  // Save path
		PrivateKey: "keys/private.pem", // Save path
	})

	if err != nil {
		lets.LogE("GenerateKeys: %w", err)
		return
	}

	keys.Save()
	return
}

func CreateSignatureExample() (signature string, err error) {
	payload := "Lets Go Framework"

	rsaKeys := &lets.RsaKeys{
		PrivateKeyFile: "keys/private.pem",
	}

	err = rsaKeys.LoadPrivateKey()
	if err != nil {
		return
	}

	crypto := lets.Crypto{
		Rsa: rsaKeys,
	}
	crypto.SetPayloadString(payload)
	crypto.CreateSignatureSHA256WithRSA()

	return crypto.GetSignatureBase64()
}

func VerifySignatureExample(signature string) error {
	payload := "Lets Go Framework"

	rsaKeys := &lets.RsaKeys{
		PublicKeyFile: "keys/public.pem",
	}
	err := rsaKeys.Load()

	crypto := lets.Crypto{
		Rsa: rsaKeys,
	}

	crypto.SetPayloadString(payload)
	crypto.SetSignatureBase64(signature)

	return crypto.VerifySignatureSHA256WithRSA()
}
