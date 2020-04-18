package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("plane_text")
	fmt.Printf("        plaintext: %v\n", plaintext)

	plaintext = pad(plaintext)
	fmt.Printf("padding plaintext: %v\n", plaintext)

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	encrypted := make([]byte, aes.BlockSize+len(plaintext))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	fmt.Printf("iv: %x\n", iv)

	encryptMode := cipher.NewCBCEncrypter(block, iv)
	encryptMode.CryptBlocks(encrypted[aes.BlockSize:], plaintext)

	fmt.Printf("encrypted: %x\n", encrypted)

	decryptMode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(plaintext))
	decryptMode.CryptBlocks(decrypted, encrypted[aes.BlockSize:])
	fmt.Printf("          decrypted: %v\n", decrypted)

	decrypted = unpad(decrypted)

	fmt.Printf("unpadding decrypted: %v\n", decrypted)
}

// PKCS7#Padding: 1-255byte block size padding
func pad(b []byte) []byte {
	padSize := aes.BlockSize - (len(b) % aes.BlockSize)
	fmt.Printf("aes.BlockSize: %d\n", aes.BlockSize) // 16
	fmt.Printf("padSize      : %d\n", padSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(b, pad...)
}

func unpad(b []byte) []byte {
	padSize := int(b[len(b)-1])
	return b[:len(b)-padSize]
}
