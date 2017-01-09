### 一步一步尝试，最终写一个 IM

client: MAC
server: Linux

##### 第一步检查并开启端口
```
netstat 检查端口占用 netstat -anlp | grep 8899
lsof 检查端口占用 lsof -i TCP | fgrep LISTEN (lsof -i tcp:port)
临时开放端口 $ sudo iptables -A INPUT -p tcp -m multiport --dport 8899  -j ACCEPT
开发端口 
    $cd /etc/sysconfig
    $sudo vim iptables
    然后参照 22 端口复制一行，改成想要开发的端口(yy复制，p粘贴)
    $service iptables restart // 重启 iptables 服务
```