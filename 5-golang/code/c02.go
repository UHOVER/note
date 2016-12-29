/**
* 通过 输入一个 域名获取IP地址
* ./c02 www.google.com
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]

	// 根据域名查找IP地址
	// addr, err := net.ResolveIPAddr("ip", name)
	// if err != nil {
	// 	fmt.Println("ResolveIPAddr error", err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println("Resolved address is ", addr.String())

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}
