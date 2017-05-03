# NodeJS

#### 安装

> 利用 Homebrew 安装

```
    $brew install nodejs
```

#### 工具的安装

Sublime Text

> 安装同步插件 sftp, 可本地编辑代码，同步到虚拟机的 Linux

```
Package Control->Install package->sftp
安装完成后，新建代码目录，在Sublime Text中选中该目录,右键，然后选择 SFTP/FTP->Map to Remote。然后会在该目录下出现 sftp-config.json 文件，修改其中配置
    
{
    "upload_on_save": false, // 默认是 false, 改为 true 则可以自动同步
    "host": "example.com",   // 域名或IP，如 192.168.100.100
    "user": "username",      // 登录的用户名，如 root
    "password": "password",  // 密码
    "port": "22",            // 端口
    
    "remote_path": "/example/path/", // 同步到 linux 对应的目录，如 /home/code/
}

这样在 sublime Text 中关联的目录下编写文件，便可传输到Linux 服务器上

```

WebStorm

```
Create New Project -> Node.js Express App
    Location:项目的位置
    Node interpreter: Node.js 可执行文件存放位置 /usr/local/bin/node

```


NPM

> NPM 是随同 NodeJS 一起安装的包管理工具，能解决 NodeJS代码部署上的问题。
> 新版的 NodeJS 已经集成了 npm, 安装NodeJS 时一并安装了

```
// 查看版本
$npm -v
// 升级 -g 表示全局安装
$sudo npm install npm -g 
// 卸载 
$npm uninstall <name> -g <--save-dev>
//更新插件
npm update <name> [-g] <--save-dev>
// 查看已安装的插件
$npm list
```

> 用于 npm 安装插件从国外服务器，可能被墙，所以可以选择 cnpm。
> cnpm 是 国内服务器，淘宝团队copy 的

```
// 安装 cnpm, cnpm 和 npm 用法一致 --registry 指定注册机构
$ npm install cnpm -g --registry=https://registry.npm.taobao.org
```

Express

> Express 是基于 Node.js 平台的 web 开发框架
> 例如现在工作目录是 project

```
// 切换到该目录下 安装 express
$cd project
$cnpm install express
```

> 在项目下产生一个 新的目录, node_modules 是 express 的存放路径

Nodemon

> 用来监控 Node.js 源代码的任何变化和自动重启服务器。

```
// 安装
$cnpm install -g nodemon
```

PM2

> 是一个负载均衡功能的 Node 应用进程管理器

```
// 安装
$cnpm install -g pm2
```

node-inspector

> 比 Node.js 的内置调试器好，跟Chrome的 Javascript调试器相似

```
// 安装
$cnpm install -g node-inspector
// 在项目下运行(含 app.js)，即可启动web 服务器进行调试
$node-inspector
```

Chrome 开发人员工具

> Network: 浏览器和服务器的各种通讯，其中包括静态资源或者向服务器发起的请求，重点关注 XHR
> Sources: 查看资源文件， comman+p 则是快速搜索

END