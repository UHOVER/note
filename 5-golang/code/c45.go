/**
* ASN1 DaytimeClient
* go run c14.go ip:port
**/

package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	daytime := time.Now()
	fmt.Println("客户端正在向服务器发送当前的时间:", daytime.String())

	mdata, _ := asn1.Marshal(daytime)
	conn.Write(mdata)

	// result, err := readFully(conn)
	// checkError(err)

	// var newtime time.Time
	// _, err = asn1.Unmarshal(result, &newtime)
	// checkError(err)

	// fmt.Println("客户端接收服务器发送的时间:", newtime.String())

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

////////////////////////////////////////////////////////////////////////////
/**
* ASN1 DaytimeServer
* client:c15.go
**/

package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		result, err := readFully(conn)
		checkError(err)

		var newtime time.Time
		_, err = asn1.Unmarshal(result, &newtime)
		fmt.Println("服务器接收到客户端的时间:", newtime.String())
		checkError(err)

		// daytime := time.Now()
		// fmt.Println("服务器接收到客户端时间后，发送自己的时间:", daytime.String())
		// mdata, _ := asn1.Marshal(daytime)
		// conn.Write(mdata)

		conn.Close()
	}
}

func readFully(conn net.Conn) ([]byte, error) {
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

