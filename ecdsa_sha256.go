package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {

	// P256, P384 or P512
	ellipticCurve := elliptic.P256()

	//privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(ellipticCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	//var publicKey ecdsa.PublicKey
	publicKey := privateKey.PublicKey

	fmt.Printf("private key: %x\n", privateKey)
	fmt.Printf("public key : %x\n", publicKey)

	//var hash hash.Hash
	hash := sha256.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	data := "This is data."

	hash.Write(([]byte)(data))
	hashed := hash.Sum(nil)

	r, s, serr := ecdsa.Sign(rand.Reader, privateKey, hashed)
	if serr != nil {
		fmt.Println(serr)
		return
	}

	fmt.Printf("r: %x\n", r)
	fmt.Printf("s: %x\n", s)

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("signature: %x\n", signature)

	// SHA256: key size=32, curve bits=256
	// SHA384: key size=48, curve bits=384
	// SHA512: key size=66, curve bits=521 (note: not typo)
	keySize := 32
	parsedR := big.NewInt(0).SetBytes(signature[:keySize])
	parsedS := big.NewInt(0).SetBytes(signature[keySize:])

	fmt.Printf("parsed r: %x\n", parsedR)
	fmt.Printf("parsed s: %x\n", parsedS)

	var c elliptic.Curve
	c = elliptic.P256()
	var x, y *big.Int
	x = publicKey.X
	y = publicKey.Y
	generatedPublicKey := ecdsa.PublicKey{Curve: c, X: x, Y: y}

	verifystatus := ecdsa.Verify(&generatedPublicKey, hashed, parsedR, parsedS)
	fmt.Printf("verify: %t\n", verifystatus)
}
