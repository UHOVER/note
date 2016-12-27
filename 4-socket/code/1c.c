/**
* 字节序的查看
**/
#include <stdio.h>
#include <arpa/inet.h>

int main(void){
	// 主机字节序
	unsigned int x = 0x12345678;
	unsigned char *p = (unsigned char*)&x;
	printf("%0x %0x %0x %0x \n", p[0],p[1],p[2],p[3] );
	

	// 转换网络字节序
	unsigned int y = htonl(x);
	p = (unsigned char*)&y;
	printf("%0x %0x %0x %0x \n", p[0],p[1],p[2],p[3] );
	return 0;
}