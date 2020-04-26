package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func main() {

	// $ openssl genrsa 2048 > private.key
	privateKeyString := `
MIIEpAIBAAKCAQEAvwtbWLSV3U8Rpyta6ul+0CAi4sok9+hB9GSsgTh3Q4FYUZXs
JsZuacL5T42QCSOlHBQOkfYVL73+LgdPDPgViz+mf6/6L11rtm9Kj1VmVXFiXvQm
AFpXx2pljERfj47G3EQcminNGK1C06Nloq8M/o1/vyn4acxwMI3wepe5vnv+NZLm
pfugXgiTt6bl2z5XQUCkKxt997g12WpVvpxkdjcpUa8dzZZWjvdFBrrJbWm8sjzB
UrTxGrs0qqGvW8SYg7xx1/gO1nizCvhZq9J4nWbP2Ltz1nxH4GApLEa7fS+Kj1An
c+rPChOlHveKwvytJvzc45ncxENQYcRl2he9uwIDAQABAoIBAGD/qCqazfim28Sv
+6KIWU3c8zmI/0orz8kBkKCvhcZtluUdpOBvIcJrL2BX9Qje40clW9x6QHmUEslW
BqoEEBQ6hhQQyotf+H+RdB7gcmvxfMvPVLgbJrRmbhbQ5GAkUw2lO4x+qtbbqbGB
Jep6zLM2LuVlru4w0cmlV7M29CUlOPOgbNG++22BaS+3wzNiPj/Gl2YbTRsoIhJx
emDKiXjzWGxd4Fk3v6t8671E/dHh4nyAhUAQg00GVZ0ClNH/6jhne5feOi8qgK6x
NNbLYraGddFVZ53bsLJG9tKGahLgc82O6OAz5J9cgfzi8utDyPq3/rkpGXLOj/aa
xQxo8ukCgYEA3Ud/y49PWs+Q4EV5zJlqT3mlQ+dODunAbcibEzw6uDQnszqZP85J
o1YqmAHFbiLuIpzcvZUspGilnPCKfFTHyznufbc91InNbnP86/stNbff4AP2hWBa
+/L0AMGFe9EUBMhvYhZzVojTnqJ0/mclxafgpDtf8b3zHFjbvGL9pd8CgYEA3QVc
HvzizzBW3CS4ATi9JSkxKpry2mnPhRT5qOwwvYyYmCqvoYSqQOMyjwNt83IPbc/0
9kO9MHjNNyIBNnRAiomQAYELFev452HIXt1drWvE3dhsBL1ciij85uVNX3OmGYWl
uGSflfYRD87Muur+eEqUV/nc6i/Rx0FxRRBsy6UCgYEAvO4TOyZ9Rrf6psIrIHnM
v1bJuJSBnVIPrqydW2sNZ8GANBNQTZ5AWWl0rJy2iTbhxEPSZTw9BZMj9D+cvlNU
0zv/WO9fp1yRPkFiLcoj6723NHmtvmtqw7vIgey5n+IACaVpFIK+r5/br5Jd+ejv
4zdXImJfpPPd4tIrq0mJ8FkCgYBCPesANpbbtgcyb6beZtz5mEDuHgaPQ4s4vbKd
2Dw7czoA0TpWVGaaj/2FM2fuwM6zANLQRDkdn/cRgRWP9oOpgdUxPjXOWiz9XCcr
l3kOEvCr9MNbIE3t8p7prOvlocm0eIPUogPadCdk73wYwXmHIAMZ4v89CRv8dja2
llelKQKBgQCK+196EYTc3Gx/19XcPH6bF6WMACMTXQ6Wv0PCa1i0rYbWP5ARMFO6
NURds5IZ8/WeGFF4Guyqoolq15pybKc4QNYuuiAc9f66PkgHRk0BzncYklHfciBM
Q935O+Rh2dqYK2FH9EZQzBUPGoK/qse9YVGE8FJohkGinBuQZA7vzQ==
`

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		fmt.Println("failed to decode privateKeyString")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(decodedPrivateKey)
	if err != nil {
		fmt.Println("failed to generate private key")
		return
	}

	data := "This is data."

	hash := crypto.Hash.New(crypto.SHA256)
	hash.Write(([]byte)(data))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Println("failed to generate signature")
		return
	}

	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("sig: " + encodedSignature)

	// $ openssl rsa -pubout < private.key > public.key
	publicKeyString := `
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvwtbWLSV3U8Rpyta6ul+
0CAi4sok9+hB9GSsgTh3Q4FYUZXsJsZuacL5T42QCSOlHBQOkfYVL73+LgdPDPgV
iz+mf6/6L11rtm9Kj1VmVXFiXvQmAFpXx2pljERfj47G3EQcminNGK1C06Nloq8M
/o1/vyn4acxwMI3wepe5vnv+NZLmpfugXgiTt6bl2z5XQUCkKxt997g12WpVvpxk
djcpUa8dzZZWjvdFBrrJbWm8sjzBUrTxGrs0qqGvW8SYg7xx1/gO1nizCvhZq9J4
nWbP2Ltz1nxH4GApLEa7fS+Kj1Anc+rPChOlHveKwvytJvzc45ncxENQYcRl2he9
uwIDAQAB
`

	decodePublicKey, err := base64.StdEncoding.DecodeString(publicKeyString)
	if err != nil {
		fmt.Println("failed to decode publicKeyString")
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(decodePublicKey)
	if err != nil {
		fmt.Println("failed to generate public key")
		return
	}

	decodedSignature, err := base64.StdEncoding.DecodeString(encodedSignature)
	if err != nil {
		fmt.Println("failed to decode signature")
		return
	}

	hash = crypto.Hash.New(crypto.SHA256)
	hash.Write([]byte(data))
	hashed = hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed, decodedSignature)
	if err != nil {
		fmt.Println("failed to verify signature")
		return
	}

	fmt.Println("success to verify")
}
