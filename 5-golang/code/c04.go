/**
*	获取 http 信息
* ./c04 www.google.com:80
 */

package main

import (
	"fmt"
	// "io/ioutil"
	"bytes"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// 第一种写法
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	// checkError(err)

	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// checkError(err)

	// _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	// checkError(err)

	// result, err := ioutil.ReadAll(conn)
	// checkError(err)

	// 第二种写法
	conn, err := net.Dial("tcp", service)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
