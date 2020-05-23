package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// P256, P384 or P521
	ellipticCurve := elliptic.P521()

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
	hash := crypto.SHA512.New()
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

	keySize := privateKey.Curve.Params().BitSize / 8
	if privateKey.Curve.Params().BitSize%8 > 0 {
		keySize += 1
	}

	// note: r and s need zero left padding if bytes size is lack. this event may occur when you use P521
	var signature []byte
	rPad := keySize - len(r.Bytes())
	if rPad > 0 {
		zeroPad := make([]byte, rPad)
		signature = append(zeroPad, r.Bytes()...)
	} else {
		signature = append(signature, r.Bytes()...)
	}

	sPad := keySize - len(s.Bytes())
	if sPad > 0 {
		zeroPad := make([]byte, sPad)
		signature = append(signature, zeroPad...)
		signature = append(signature, s.Bytes()...)
	} else {
		signature = append(signature, s.Bytes()...)
	}

	fmt.Printf("signature: %x\n", signature)

	// SHA256: key size=32, curve bits=256
	// SHA384: key size=48, curve bits=384
	// SHA512: key size=66, curve bits=521 (note: not typo)
	keySize = 66
	parsedR := big.NewInt(0).SetBytes(signature[:keySize])
	parsedS := big.NewInt(0).SetBytes(signature[keySize:])

	fmt.Printf("parsed r: %x\n", parsedR)
	fmt.Printf("parsed s: %x\n", parsedS)

	var c elliptic.Curve
	c = elliptic.P521()
	var x, y *big.Int
	x = publicKey.X
	y = publicKey.Y
	generatedPublicKey := ecdsa.PublicKey{Curve: c, X: x, Y: y}

	verifystatus := ecdsa.Verify(&generatedPublicKey, hashed, parsedR, parsedS)
	fmt.Printf("verify: %t\n", verifystatus)
}
