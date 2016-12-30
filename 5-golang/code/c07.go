/**
* 回射客户端
* 与 c6.go 的回射服务器对应
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println("服务器IP地址:", tcpAddr.IP.String())

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HELLOWORLD"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("CLIENT:", string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
