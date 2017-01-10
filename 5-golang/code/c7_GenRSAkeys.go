/**
* GenRSAKeys
* gen Gob & Pem
**/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes ", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent ", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modules", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("data/private.key", key)
	saveGobKey("data/public.key", publicKey)

	savePEMkey("data/private.pem", key)

}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func savePEMkey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
