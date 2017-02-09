/*
* 回射服务器
*/

#define ERR_EXIT(m) \
		do \
		{ \
			perror(m)； \
			exit(EXIT_FAILURE); \

		} while(0)


int main(void)
{
	int listenfd;
	// PF_INET:IPv4, SOCK_STREAM: 流式套接字, 0:表示内核自己去选择协议
	// (listenfd = socket(PF_INET, SOCK_STREAM, 0)) < 0 
	if ((listenfd = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
		ERR_EXIT("socket")

	struct sockaddr_in servaddr;
	memset(&servaddr, 0, sizeof(servaddr)); 
	servaddr.sin_family = AF_INET;
	servaddr.sin_port = htons(5188); // 端口
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY); // INADDR_ANY 本机的任意地址
	// servaddr.sin_addr.s_addr = inet_addr("127.0.0.1");
	// inet_aton("127.0.0.1", &servaddr.sin_addr);

	if (bind(listenfd, (struct sockaddr*)&servaddr, sizeof(servaddr)) < 0)
		ERR_EXIT("bind")
}