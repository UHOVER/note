/*
* 点对点客户端
*/

#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>

#define ERR_EXIT(m) 				\
		do 							\
		{ 							\
			perror(m); 				\
			exit(EXIT_FAILURE); 	\
		} while(0) 


int main(void)
{
	int sock;
	// PF_INET:IPv4, SOCK_STREAM: 流式套接字, 0:表示内核自己去选择协议
	// (sock = socket(PF_INET, SOCK_STREAM, 0)) < 0 
	if ((sock = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
		ERR_EXIT("socket");

	struct sockaddr_in servaddr;
	memset(&servaddr, 0, sizeof(servaddr)); 
	servaddr.sin_family = AF_INET;
	servaddr.sin_port = htons(5188); // 端口
	servaddr.sin_addr.s_addr = inet_addr("127.0.0.1"); // 服务端地址
	
	if (connect(sock,(struct sockaddr*)&servaddr, sizeof(servaddr)) < 0)
		ERR_EXIT("connect");

	pid_t pid;
	pid = fork();
	if (pid == -1)
		ERR_EXIT("fork");

	if (pid == 0)
	{
		char recvbuf[1024];
		while(1)
		{
			memset(recvbuf, 0, sizeof(recvbuf));
			int ret = read(sock, recvbuf, sizeof(recvbuf));
			if (ret == -1)
				ERR_EXIT("read");
			else if (ret == 0)
				printf("peer close\n");

			fputs(recvbuf, stdout);
		}
		close(sock);
	}
	else
	{
		char sendbuf[1024] = {0};
		while(fgets(sendbuf,sizeof(sendbuf), stdin) != NULL)
		{
			write(sock, sendbuf, strlen(sendbuf));
			memset(sendbuf, 0, sizeof(sendbuf));
		}
		close(sock);

	}

	return 0;
}