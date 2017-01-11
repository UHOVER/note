/**
* TLSEchoServer
* 当前互联网上最流行的加密消息传输方案是TLS（Transport Layer Security安全传输层），其前身为SSL（Secure Sockets Layer安全套接字层）。
* 在TLS中，客户端和服务器之间使用X.509证书进行身份验证。身份验证完成后，两者之间会生成一个密钥，所有的加密和解密过程都使用这个密钥。虽然客户端和服务端协商的过程相对较慢，但一旦完成就会使用一个较快的私钥机制。
**/

package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("data/wahhung.suen.pem", "data/private.pem")
	checkError(err)
	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	service := "0.0.0.0:8090"
	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)
	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

}
