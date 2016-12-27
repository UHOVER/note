/**
* 地址转换函数
**/

#include <stdio.h>
#include <netinet/in.h>
#include <arpa/inet.h>

int main(void){
	// 将 IP 地址转换为 网络字节序 32位的整数
	unsigned long addr = inet_addr("192.168.0.100");
	printf("addr=%u\n", ntohl(addr));

	// 将地址结构转换为点分十进制的IP地址
	struct in_addr ipaddr;
	ipaddr.s_addr = addr;
	printf("%s\n", inet_ntoa(ipaddr));

	return 0;
}