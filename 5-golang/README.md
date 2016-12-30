网关
```
网关是一个统称，它用于连接起一个或多个网络。其中的中继器在物理层面上进行操作，它将信息从一个子网复制到另一个子网上。桥接在数据连接层面上进行操作，它在网络之间复制帧。路由器在网络层面上进行操作，它不仅在网络之间复制信息，还决定了信息的传输路线。
```

连接模型(TODO)
```
面向连接模型：即为会话建立单个连接，沿着连接进行双向通信。当会话结束后，该连接就会断开。这类似于电话交谈。例子就是TCP。

无连接模型：在无连接系统中，消息的发送彼此独立。这类似于普通的邮件。无连接模型的消息可能不按顺序抵达。例子就是IP协议。面向连接的传输可通过无连接模型——基于IP的TCP协议建立。无连接传输可通过面向连接模型——基于IP的HTTP协议建立。
```

通信模型(TODO)
```
消息传递
远程过程调用(RPC)

```

协议(TODO)
```
IP和TCP/UDP协议基本上就等价于网络协议栈。例如, 蓝牙定义了物理层和协议层，但最重要的是IP协议栈，可以在许多蓝牙设备使相同的互联网编程技术。同样, 开发4G无线手机技术，如LTE（Long Term Evolution）也将使用IP协议栈。
```

IPv6 地址
```
128位地址
由':'分隔的4位16进制组成。如：2002:c0e8:82e7:0:0:0:c0e8:82e7
例如："localhost"地址是：0:0:0:0:0:0:0:1，可以缩短到::1(可以省略一些零和重复的数字)
```

net
```
IP 类型
IP类型在net包中被定义位一个字节数组
type IP []byte

IP 掩码
type IPMask []byte
// 用一个4字节的IPv4地址来创建一个掩码
func IPv4Mask(a, b, c, d byte) IPMask
// 一个IP的方法返回默认的子网掩码,一个掩码的字符串形式是一个十六进制数，如掩码255.255.0.0为ffff0000
func (ip IP) DefaultMask() IPMask
// 一个掩码可以使用一个IP地址的方法，找到该IP地址的网络(网络地址)
func (ip IP) Mask(mask IPMask) IP

IPAddr 类型
type IPAddr{
    IP IP
}
// 通过IP主机名(域名)，执行NDS 查找，返回IP地址
// 参数：net: "ip","ip4" 或者 "ip6"
func ResolveIPAddr(net,addr string)(*IPAddr, os.Error)
// 通过域名查找多种IP地址，IPv6,IPv4 ,返回数组
func LookupHost(host string) (addrs []string, err error)
// 通过传输协议(TCP/UDP/SCTP) 和 服务类型(ftp/ssh) 获得其对应的端口号
func LookupPort(network, service string) (port int, err os.Error)

TCPAddr 类型
type TCPAddr struct{
    IP IP
    Port int
}
//获取TCPAddr(ip+port) net:"ip","ip4" 或者 "ip6" addr: 主机名/IP地址:port(端口)
func ResolveTCPAddr(net,addr string)(*TCPAddr, os.Error)
```

TCP 套接字
```
net.TCPConn 是允许在客户端和服务器之间的全双工通信的GO类型
// TCPConn 被客户端和服务器用来读写消息
func (c *TCPConn) Write(b []byte)(n int, err os.Error)
func (c *TCPConn) Read(b []byte)(n int, err os.Error)
//客户端建立TCP连接，返回用于通信的TCPConn
//laddr:本地地址，通常设置为nil; raddr:是一个服务的远程地址;net:字符串,"ip","ip4","ip6"
func DialTCP(net string, laddr, raddr *TCPAddr)(c *TCPConn, err os.Error)

在一个服务器上注册并监听一个端口，然后它阻塞在一个 "accept" 操作。并等待客户端连接。
当一个客户端连接，accept 调用返回一个连接(connection)对象。
// 如果想监听所有网络，IP地址应设置为 0，只是想监听一个简单网络接口，IP地址可以设置为该网络的地址。如果端口设置为0，O/S会选择一个端口。
func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
func (l *TCPListener) Accept() (c Conn, err os.Error)

//超时(在套接字读写前)
func (c *TCPConn) SetTimeout(nsec int64) os.Error
//存活状态
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

UDP
```
func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
```










