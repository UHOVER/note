# NodeJS

### Node

- 回调函数

> Node.js 异步编程的直接体现就是回调
> 回调函数在完成任务后就会被调用，Node所有API都支持回调函数。
> 可以一边读取文件，一边执行其他命令，在文件读取完后，将文件内容作为回调函数的参数返回。在执行代码时就没有阻塞或等待 I/O 操作，可以处理并发请求。

```js
    var fs = require("fs");
    
    // 阻塞处理
    var data = fs.readFileSync('input.txt');

    // 非阻塞处理
    var data = fs.readFile('input.txt',function(err,data){
        if (err) return console.error(err);
        });

```

- 事件循环

> Node.js 是单线程应用程序，但通过事件和回调支持并发
> Node.js 每个API都是异步的，并作为一个独立线程运行，使用异步函数调用，并处理并发。
> Node.js 基本上所有的事件机制都是用设计模式中观察者模式实现。
> Node.js 单线程类似进入一个 while(true) 的事件循环，直到没有事件观察者后退出，每个异步事件都生成一个事件观察者，如果有事件发生就调用该回调函数。

事件驱动程序

> Node.js 使用事件驱动模型
> 当 web server 接收到请求，就把这个请求关闭进行处理，然后去服务下一个web请求；
> 当 这个请求的处理完成后，它被放回处理队列，当到达队列开头，这个结果被返回给用户。
> webserver 一直接受请求而不等待任何读写操作。也被称为非阻塞式IO或事件驱动IO

```
    // 引入 events 模块
    var events = require('events');
    // 创建 eventEmitter 对象
    var eventEmitter = new events.EventEmitter();

    //绑定事件及事件的处理程序
    eventEmitter.on('eventName', eventHandler);

    // 触发事件
    eventEmitter.emit('eventName');
```

EventEmitter 类

> events 模块只提供一个对象: events.EventEmitter
> EventEmitter 的核心就是事件触发与事件监听器功能的封装
> 
> 不会直接使用 EventEmitter, 而是在对象中继承它。

- Buffer(缓冲区)

> JS 只有字符串,没有二进制数据类型
> 处理TCP流或文件流时，必须使用二进制。
> 在Node.js 中，定义了一个Buffer类，存放二进制数据的缓冲区,核心库。
> 默认是 utf-8 格式

```js
// 创建
var buf = new Buffer("www.baidu.com", "utf-8");
// 写入缓冲区 offset[开始索引，默认0] length[写入字节数，默认buffer.lenth] encoding[使用编码，默认 utf-8]
len = buf.write(string[, offset[, length]][, encoding])

// 从缓冲区读取数据 start[开始读取索引，默认0] end[结束位置，默认为缓冲区末尾]
str = buf.toString([encoding[, start[, end]]]);

// Buffer 转换为 JSON 对象
var json = buf.toJSON();

// 缓冲区合并 list[Buffer 对象的数组] totalLength[合并后的总长度]
var buf = Buffer.concat(list[, totalLength]);

```












### Http 基础

```
http://www.baidu.com:80/course?id=21302   [URL]
协议: [http] 约定好的通信内容格式
主机名称:[www.baidu.com] 请求的服务器(DNS 域名解析)
端口:[80] Http 默认使用80端口，可以省略
主机+端口:确定了唯一的通信通道，可以在此通道上完成通信过程

资源路径:[/course] 指明请求当前 WEB 服务器上的什么资源，服务端按此部分内容决定处理行为
查询参数:[?id=21302]
```

通信过程
> 用户浏览器-> 请求(Request, req)[Head和Body两部分；Head描述请求的基本信息；Body:要发送给服务器的数据]
> -> 服务端应用概念:路由(Route)[根据请求资源路径决定处理方法，生成响应结果，并将结果返回给浏览器]
> -> 响应(Request, res)[Head和Body两部分；Head描述返回信息的属性；Body:返回的数据内容]
> 
> 请求: 浏览器发给服务器的数据包
> 路由: 服务器匹配处理模块的过程
> 响应: 服务器返回给浏览器的数据包

END