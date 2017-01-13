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

Conn,PacketConn 和 Listener 类型
```
//net可以是"tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only)和"ip6" (IPv6-only)任何一种。它将返回一个实现了Conn接口的类型。注意此函数接受一个字符串而不是raddr地址参数，因此，使用此程序可避免的地址类型。
func Dial(net,laddr,raddr string)(c Conn,err os.Error)
func Listen(net, laddr string) (l Listener, err os.Error)
func (l Listener) Accept() (c Conn, err os.Error)

```


### 数据序列化
```
IP，TCP或者UDP网络包并不知道编程语言的数据类型的含义，它们只是字节序列的载体。因此，写入网络包的时候，应用需要将要传输的(有类型的)数据 序列化 成字节流，反之，读取网络包的时候，应用需要将字节流反序列化成合适的数据结构，这两个操作被分别称为编组和解组。
func Marshal(var interface{})([]byte, os.Error)
func Unmarshal(var interface{}, b []byte)(rest []byte, err os.Error)

Go内部实际是使用reflect包来编、解组结构，因此reflect包必须能访问所有的字段。

json/gob/base64
NewEncoder/NewDecoder

gob 包
只能编码 Go 的数据类型
为了使用gob编组一个数据值，首先你得创建Encoder。它使用Writer作为参数，编组操作会将最终结果写入此流中。encoder有个Encode方法，它执行将值编组成流的操作。此方法可以在多份数据上被调用多次。但是对于每一种数据类型，类型信息却只会被写入一次。
你将使用Decoder来执行解组序列化后的数据流的操作。它持有一个Reader参数，每次读取都将返回一个解组后的数据值。
```

### 应用协议
```
客户端和服务器需要通过消息来进行交互。TCP和UDP是信息交互的两种传输机制。在这两种传输机制之上就需要有协议来约定传输内容的含义。协议清楚说明分布式应用的两个模块之间交互消息的消息体、消息的数据类型、编码格式等。
```

### 安全
```
hash 算法
crypto/md5
func NewMD5(key []byte) hash.Hash

key对称加密
Blowfish和DES

当前互联网上最流行的加密消息传输方案是TLS（Transport Layer Security安全传输层），其前身为SSL（Secure Sockets Layer安全套接字层）。
在TLS中，客户端和服务器之间使用X.509证书进行身份验证。身份验证完成后，两者之间会生成一个密钥，所有的加密和解密过程都使用这个密钥。虽然客户端和服务端协商的过程相对较慢，但一旦完成就会使用一个较快的私钥机制。
```

### http
```
URL:指定资源的位置(资源通常为 HTML文档、图片、声音文件等静态文件和动态生成的对象，如根据数据信息生成)
请求资源返回并不是资源本身，而是资源的代表，如静态文件的副本。
HTTP协议是无状态，面向连接和可靠的。每次请求都包括一个独立的TCP连接。

// 最简单的请求是由用户代理发起 "HEAD" 命令(http.Head func), 响应状态对应response 中的 Status 属性(response.Status)，Header 属性对应 HTTP 响应的 header 域(response.Header)
func Head(url string) (r *Response, err os.Error)
// GET 请求收到的是一个资源的内容(http.Get func),响应内容为response的Body属性(response.body),是一个 io.ReadCloser 类型
func Get(url string) (r *Response, finalURL string, err os.Error)

简单代理
向代理服务器发送一个"GET"请求，但是请求URL必须是完整的目标地址。此外，设置代理的HTTP头应当包含 "Host" 字段。只要代理服务器设置为运行这样的请求通过。
Go 把这看成 HTTP 传输层的一部分。可使用 Transport 类进行管理。假设有一个代理服务器地址字符串URL，相应的创建Transport对象并交给Client对象的代码：
proxyURL, err := url.Parse(URL)
transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
client := &http.Client{Transport: transport}

有些代理服务器要求通过用户名和密码进行身份验证才能传递请求。一般的方法是“基本身份验证”：将用户名和密码串联成一个字符串“user:password”，然后进行Base64编码，然后添加到HTTP请求头的“Proxy-Authorization”中，再发送到代理服务器
basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
...
request.Header.Add("Proxy-Authorization", basic)

HTTP = HTTP+TLS
服务器必须在客户端接受从其数据前返回有效的X.509证书。如果证书有效，Go会在内部处理好所有的事情，而客户端会在使用HTTPS地址

Go提供了一个multi-plexer，即一个读取和解释请求的对象。它把请求交给运行在自己线程中的handlers。这样，许多读取HTTP请求，解码并转移到合适功能上的工作都可以在各自的线程中进行。
对于文件服务器，Go提供了一个FileServer对象，它知道如何发布本地文件系统中的文件。它需要一个“root”目录，该目录是在本地系统中文件树的顶端；还有一个针对URL的匹配模式。最简单的模式是“/”，这是所有URL的顶部，可以匹配所有的URL。

```

### 模板
```
大多数服务器端语言的机制主要是在静态页面插入一个动态生成的组件，如清单列表项目。GO的template包中采取了相对简单的脚本化语言。
源文件被称作 template ，包括文本传输方式不变，以嵌入命令可以作用于和更改文本。命令规定如 {{ ... }} ，类似于JSP命令 <%= ... =%> 和PHP命令 <?php ... ?>。
import "html/template"
type Job struct {
    Employer string
    Role     string
}
-> 模板
{{with .Jobs}}
    {{range .}}
        An employer is {{.Employer}}
        and the role is {{.Role}}
    {{end}}
{{end}}


```





















