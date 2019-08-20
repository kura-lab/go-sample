package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/big"
)

func main() {
	// generate rsa private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("failed to generate private key")
		return
	}

	publicKey := privateKey.PublicKey
	modulus := base64.RawURLEncoding.EncodeToString(publicKey.N.Bytes())
	exponent := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(publicKey.E)).Bytes())

	fmt.Println("modulus: " + modulus)
	fmt.Println("exponent: " + exponent)

	data := "This is data."

	hash := crypto.Hash.New(crypto.SHA256)
	hash.Write(([]byte)(data))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Println("failed to generate signature")
		return
	}

	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)

	fmt.Println("sig: " + encodedSignature)

	// generate public key from n(modulus) and e(exponent)
	decodedModulus, err := base64.RawURLEncoding.DecodeString(modulus)
	if err != nil {
		fmt.Println("failed to decode modulus")
		return
	}
	n := big.NewInt(0)
	n.SetBytes(decodedModulus)

	decodedExponent, err := base64.StdEncoding.DecodeString(exponent)
	if err != nil {
		fmt.Println("failed to decode exponent")
		return
	}
	var exponentBytes []byte
	if len(decodedExponent) < 8 {
		exponentBytes = make([]byte, 8-len(decodedExponent), 8)
		exponentBytes = append(exponentBytes, decodedExponent...)
	} else {
		exponentBytes = decodedExponent
	}
	reader := bytes.NewReader(exponentBytes)
	var e uint64
	err = binary.Read(reader, binary.BigEndian, &e)
	if err != nil {
		fmt.Println("failed to read binary exponent")
		return
	}
	generatedPublicKey := rsa.PublicKey{N: n, E: int(e)}

	// verifiy signature
	decodedSignature, err := base64.RawURLEncoding.DecodeString(encodedSignature)
	if err != nil {
		fmt.Println("failed to decode signature")
		return
	}

	hash = crypto.Hash.New(crypto.SHA256)
	hash.Write([]byte(data))
	hashed = hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(&generatedPublicKey, crypto.SHA256, hashed, decodedSignature)
	if err != nil {
		fmt.Println("failed to verify signature")
		return
	}

	fmt.Println("success to verify")
}
