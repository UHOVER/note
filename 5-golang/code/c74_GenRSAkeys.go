/**
* GenRSAKeys
* gen Gob & Pem
*
* 公钥加密
* 公钥加密和解密需要两个key：一个用来加密，另一个用来解密。加密key通常是公开的，这样任何人都可以给你发送加密数据。解密key必须保密，否则任何人都可以解密数据。公钥系统是非对称的，不同的key有不同的用途。
* Go支持很多公钥加密系统，RSA就是一个典型的例子。
* 下面是一个生成RSA公钥和私钥的程
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
