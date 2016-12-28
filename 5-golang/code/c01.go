/**
* 将
* ./c01 127.0.0.1
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// os.Args 获取命令行参数 os.Args[0] 为执行的程序 os.Args[1] 第一个参数
	// 例如 $./c01 127.0.0.1 os.Args[0]=./c01 os.Args[1]=127.0.0.1
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	dotAdder := os.Args[1]

	// ParseIP 解析IP，IPv4和IPv6，无效IP文本返回 nil
	addr := net.ParseIP(dotAdder)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	// 返回默认的掩码
	mask := addr.DefaultMask()
	// 一个掩码可以使用一个IPhone地址的方法，找到该IP地址的网络
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		" Leading ones count is ", ones,
		" Mask is (hex) ", mask.String(),
		" Network is ", network.String())

	os.Exit(0)
}
