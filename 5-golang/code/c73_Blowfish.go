/**
* Blowfish
*
* 数据加密有两种机制。第一种方式是在加密和解密时都使用同一个key。加密方和解密方都需要知道这个key。此处如何在这两者之间传输这个key。
* 目前有很多使用hash算法的加密算法。其中很多都有弱点，而且随着时间的推移，计算机越来越快，通用hash算法变得越来越弱。Go已经支持好几个对称加密算法，如Blowfish和DES。
* 这些算法都是block算法。因为它们必须基于数据块(block)。如果你的数据不匹配block的大小，那就必须在最后使用空格来填充多余的空间。
* 每个算法都被表示为一个Cipher对象。可通过在相应的包中使用对称key作为参数调用NewCipher方法来创建该对象。
* 创建cipher对象后，就能通过它加密和解密数据块。Blowfish需要8位的block，详见以下程序
**/

package main

import (
	"bytes"
	"code.google.com/p/go.crypto/blowfish"
	"fmt"
)

func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	src := []byte("hello\n\n\n")
	var enc [512]byte

	cipher.Encrypt(enc[0:], src)

	var decrypt [8]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	fmt.Println(string(result.Bytes()))
}
