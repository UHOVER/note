/**
* MD5Hash
*
* Go支持几个hash算法,包括MD4, MD5, RIPEMD-160, SHA1, SHA224, SHA256, SHA384 and SHA512。它们都尽可能按照Go程序员关注的，遵循相同的模式：在适当的包中定义New或类似的方法，返回一个hash包中的Hash 对象。
* 一个Hash结构体拥有一个io.Writer接口，你可以通过writer方法写入被hash的数据.你可以通过Size方法获取hash值的长度，Sum方法返回hash值。
* MD5算法是个典型的例子。使用md5包，hash值是一个16位的数组。通常以ASCII形式输出四个由4字节组成的十六进制数。程序如下
**/

package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
