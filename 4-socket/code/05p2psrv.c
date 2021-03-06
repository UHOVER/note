/*
* 点对点服务器
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
	int listenfd;
	// PF_INET:IPv4, SOCK_STREAM: 流式套接字, 0:表示内核自己去选择协议
	// (listenfd = socket(PF_INET, SOCK_STREAM, 0)) < 0 
	if ((listenfd = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
		ERR_EXIT("socket");

	struct sockaddr_in servaddr;
	memset(&servaddr, 0, sizeof(servaddr)); 
	servaddr.sin_family = AF_INET;
	servaddr.sin_port = htons(5188); // 端口
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY); // INADDR_ANY 本机的任意地址
	// servaddr.sin_addr.s_addr = inet_addr("127.0.0.1");
	// inet_aton("127.0.0.1", &servaddr.sin_addr);

	// 开启地址重复使用，不用等待 TCP 的 TIME_WAIT 状态消失
	int on = 1; // on=1 表示开启
	// SO_REUSEADDR: 不同的选项会有不同的结构
	if (setsockopt(listenfd, SOL_SOCKET, SO_REUSEADDR, &on, sizeof(on)) < 0)
		ERR_EXIT("setsockopt");

	if (bind(listenfd, (struct sockaddr*)&servaddr, sizeof(servaddr)) < 0)
		ERR_EXIT("bind");

	// SOMAXCONN 队列的最大值
	if (listen(listenfd, SOMAXCONN) < 0)
		ERR_EXIT("listen");

	struct sockaddr_in peeraddr;
	socklen_t peerlen = sizeof(peeraddr);
	int conn; // 已连接套接字，主动
	pid_t pid;
	pid = fork(); // 创建一个子进程，父进程用来发送数据，子进程用来接收数据
	if (pid == -1)
		// 失败
		ERR_EXIT("fork");
	
	if (pid == 0)
	{
		char sendbuf[1024] = {0};
		while (fgets(sendbuf, sizeof(sendbuf), stdin) != NULL) 
		{
			write(conn, sendbuf, strlen(sendbuf));
			memset(sendbuf, 0, sizeof(sendbuf));
		}
		printf("child close\n");
		exit(EXIT_SUCCESS);

	}else{
		char recvbuf[1024];
		while(1)
		{
			memset(recvbuf, 0, sizeof(recvbuf));
			int ret = read(conn,recvbuf,sizeof(recvbuf));
			if (ret == -1)
				ERR_EXIT("read");
			if (ret == 0)
			{
				// 客户端关闭了,返回
				printf("peer_close\n");
				break;
			}

			fputs(recvbuf,stdout);
		}	
		printf("parent close\n");
		exit(EXIT_SUCCESS);
	}	
	
	return 0;
}