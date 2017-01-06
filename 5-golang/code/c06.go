/**
* 回射服务器 + 多线程 = 多线程回射服务器
* EchoServer + Thread = ThreadedEchoServer
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":8090"
	// 获取 ip+port
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	// checkError(err)
	// fmt.Println("服务器IP地址:", tcpAddr.IP.String())

	// listener, err := net.ListenTCP("tcp", tcpAddr)
	// checkError(err)

	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("SERVER:", string(buf[0:]))
		// _, err2 := conn.Write(buf[0:n])
		// if err2 != nil {
		// 	return
		// }
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
