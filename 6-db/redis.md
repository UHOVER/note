NoSQL (不存在表结构)
```shell
# 特点
# 1、处理超大量的数据
# 2、运行在便宜的PC服务器集群上
# 3、击碎了性能瓶颈

# 使用场景
# 1、对数据高并发读写
# 2、对海量数据的高效存储和访问
# 3、对数据的高效可扩展性和高可用性(没有固定的表结构)
```


### Redis
```shell
# 数据结构服务器

#Redis的优点
> 性能极高 – Redis能支持超过 100K+ 每秒的读写频率。
> 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。
> 原子 – Redis的所有操作都是原子性的，同时Redis还支持对几个操作全并后的原子性执行。
> 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。
> 为了保证效率，数据都是缓存在内存中

# 适用场景
# 1、应用程序先去访问 Redis,获取数据失败再访问 MySQL
# 2、灵活的数据结构和数据操作，海量存储
> 取最新的 N 个数据的操作
> 排行榜应用，取 TOP N 操作
> 需要精确设定过期时间的应用(对键设计过期)
> 计数器应用
> Uniq 操作，获取某段时间所有数据排重值
> 实时系统，反垃圾系统
> Pub/Sub 构建实时消息系统(redis独有发布/订阅)
> 构建队列系统(栈和队列)
> 缓存
```

安装
```shell
# CentOS7 安装 Redis
# 进入 root，下载 Redis 安装包(https://redis.io/download 下的Stable寻找路径)
wget http://download.redis.io/releases/redis-3.2.6.tar.gz
# 解压
tar xvfz redis-3.2.6.tar.gz
# 解压后，进行编译
cd redis-3.2.6
make
make install
# 运行测试，确定 redis 的功能是否正常
make test
# 发现 All test passed without errors! 则表示成功

# 错误
# 发现安装时，You need tcl 8.5 or newer in order to run the Redis test

# 需安装tcl
wget http://downloads.sourceforge.net/tcl/tcl8.6.1-src.tar.gz  
sudo tar xzvf tcl8.6.1-src.tar.gz  -C /usr/local/  
cd  /usr/local/tcl8.6.1/unix/  
sudo ./configure  
sudo make  
cd src
sudo make install 

# 安装完成，在切到 redis-3.2.6 进行 $make install 然后 $make test，安装成功

# 切换到 utils 目录下，执行 redis 初始化的脚本 install_server.sh
cd utils/
./install_server.sh

#redis初始化后redis配置文件为/etc/redis/6379.conf，日志文件为/var/log/redis_6379.log，数据文件dump.rdb存放到/var/lib/redis/6379目录下，启动脚本为/etc/init.d/redis_6379

# 现在我们要使用 systemd，所以在 /etc/systems/system 下创建一个单位文件名字为 redis.service。
vim /etc/systemd/system/redis.service # 下面为 redis.service 的内容
[Unit]
Description=Redis on port 6379
[Service]
Type=forking
ExecStart=/etc/init.d/redis_6379 start
ExecStop=/etc/init.d/redis_6379 stop
[Install]
WantedBy=multi-user.target

# 启动 Redis-server
sudo service redis start
# 查看 Redis
ps -ef | grep redis
# 通过启动命令检查 Redis 服务器状态
netstat -nlt | grep 6379
# 如果启动失败，可以查看失败原因
systemctl status redis.service
# 启动 Redis-client
redis-cli


```

## 数据类型

strings
```shell
# 一个key对应一个value。二进制安全，一个键只需要有一个

set
# 设置key对应的值为string类型的value
6379> set name xiaoming # set key value
setnx(set not exit)
# 设置key对应的值为string类型的value。如果key已经存在，返回0。不覆盖
6379> setnx name xiaoming
setex
# 设置key对应的值为string类型的value,并指定此键值对应的有效期
6379> setex name 10 xiaoming # setex key time value
# 设置 name的值位xiaoming,有效期是10s,10s后获取name的值为null
setrange
# 设置指定key的替换value 的某些字符
假如 name=xiaoming
6379> setrange name 2 haha # setrange key index substr
# index 是从value的下标为index开始，替换 substr，substr 字符个数即替换的个数
# get name -> xihahang
mset
# 批量设置多个key的值，成功返回ok表示所有值都被设置，失败返回0表示没有任何值被设置
6379> mset key1 value1 key2 values2 ... 
msetnx
# 相当于批量设置多个 setnx

get
# 获取key的值
6379> get name # 获取字符串
getset
# 设置key的值，返回key的旧值
6379> getset name hello
getrange
# 获取子字符串
6379> get range 0 5 # 获取 0 到 5 之间的字符
mget
# 批量获取

incr 
# 递增1 设置不存在的key，认为不存在的key默认为0
6379> incr num
incrby
# 指定递增多少，设置不存在的key，认为不存在的key默认为0
6379> incrby num 5  
decr
# 减1
decrby
# 减 n

append
# 拼接字符串，返回字符串的长度
6379> append name .com
strlen
# 获取指定key的值的长度
6379> strlen name

```

hashes 类型
```shell
Redis hash 是一个string 类型的field 和 value 的映射表，适合存储对象

hset 
# 设置 hash filed 为指定值，若key不存在，则先创建
6379> hset myhash field1 value
6379> hset user:001 name hell0
hsetnx
# key不存在，先创建，key存在则设置不成功，返回0
hmset
# 批量设置
6379> hset myhash field1 value1 field2 value2 ..
6379> hset user:001 name hello age 20 ...

hget
# 获取
6379> hget user:001 name
hmget
# 批量获取
6379> hmget user:001 filed1 filed2 filed3

hincrby
# 给 hash field 加上给定的值
6379> hincrby myhash filed1 num

hexists
# hash field 是否存在
6379> hexists myhash field1

hlen
# 返回 hash field 的数量
6379> hlen myhash

hdel
# 删除指定 hash 的 field
6379> hdel myhash field

hkeys
# 返回 hash 所有 field
6379> hkeys myhash
hvals
# 返回hash所有value
6379> hvals myhash

hgetall
# 获取 hash 中所有 field 及 value

```

list 类型
```shell
# list 是一个链表，每一个子元素都是 string 类型的双向链表

lpush
# 在key对应的list的头部添加字符串元素  # 先进后出
6379> lpush mylist value

rpush
# 在key对应的list的尾部部添加字符串元素 # 先进先出
6379> rpush mylist value

linsert
# 在key对应的list的特定位置前或后添加字符串(添加到已经存在的值的前或后面)
6379> linsert mylist before/after exists_value new_value 

lset
# 指定下标的元素替换新值(替换不存在的下标，报错)
6379> lset mylist index new_value

lrem
# 从list中，删除n个和 value相同的元素，n<0从尾部删除，n=0全部删除(一个不剩)
6379> lrem mylist n value # value不存在返回0

ltrim
# 保留 list 下标为 n 到 m 的元素，其他删除
6379> ltrim mylist n m

lpop
# 在 list 的头部，弹出一个元素
6379> lpop mylist
rpop
# 在 list 的尾部，弹出一个元素
6379> rpop mylist
rpoplpush
# 从第一个list尾部移除一个元素，并添加到第二个list的头部
6379> rpoplpush list1 list2

lrange
# 从 头的 第n个元素到尾部第m个元素，依次取出
6379> lrange mylist 0 -1 # 从头部的第一个到尾部最后一个

lindex
# 返回list的索引index下的值
6379> lindex mylist index

llen
# 返回list的长度
6379> llen mylist

```

sets类型
```shell
# set是集合，是string类型的无序集合。通过hast table 实现的

sadd
# 向集合添加元素
6379> sadd myset value # add相同的元素，会失败，返回0

srem
# 向集合删除元素
6379> srem myset value

spop
# 随机删除并返回一个元素
6379> spop myset 
srandmember
# 随机返回集合的一个元素，但不删除元素
6379> srandmember myset

smember
# 查看集合的元素
6379> smember myset

sdiff
# 差集，返回其他几个集合和第一个集合的差集(存在于第一个集合，却不存在与其他集合的元素)
6379> sdiff set1 set2 ...
sdiffstore
# 取差集并存储(将第二个集合，和后面N个集合的差集，存储在第一个集合)
6379> sdiffstore set_new set1 set2 ...

sinter
# 交集(取N个集合的交集)
6379> sinter set1 set2 ...
sinterstore
# 从第二个集合开始取N个集合的交集，存到第一个集合中
6379> sinterstore set_new set1 set2...

sunion
# 并集
6379> sunion set1 set2...
sunionstore
# 从第二个集合开始取N个集合的并集，存到第一个集合中
6379> sunionstore set_new set1 set2...

smove
# 从第一个集合中移除元素并添加到第二个集合
6379> smove set1 set2 set1_value # 如set1_value 这个元素在set2 中存在，那么仅仅是在set1中删除

scard
# 返回集合元素的个数
6379> scard myset 

sismember
# 判断某元素是否为集合的元素
6379> sismember myset value

srandmember
# 随机返回集合的一个元素，但不删除元素
6379> srandmember myset
```

sorted set 类型(zset)
```shell
# 是set的一个升级版本，在set基础上增加顺序属性，zst是有序集合

zadd
# 向 zset 中添加指定顺序的元素，如果元素相同，则会更新顺序
6379> zadd myzset index value
zrem
# 指定删除集合元素
6379> zrem myzset value

zrange
# 取出所有元素(顺序号的升序排序)
6379> zrange myzset 0 -1 withscores # 输出元素和顺序号
zrevrange
# 取出所有元素(顺序号的降序排序)
6379> zrevrange myzset 0 -1 withscores # 输出元素和顺序号

zincrby
# 如果集合中存在元素value，则该元素的顺序会增加num，如果不存在，则向集合中添加元素为value，且顺序为num
6379> zincrby myzset num value

zrank
# 获取集合某个元素的下标(下标是集合的顺序从小到大排序的下标顺序，下标不等于顺序号，顺序号是下标排序的依据)
6379> zrank myzset value

zrevrank
# 返回集合中元素的顺序号从大到小排名的下标(即翻转zrank)

zrangebysore
# 取顺序号 n 到 n 的 元素(给定区间)
6379> zrangebysore myzset 2 3 withscores

zcount
# 返回顺序号 n 到 m 的元素个数(给定区间)
6379> zcount myzset 2 3
zcard
# 返回集合元素的个数

zremrangebyrank
# 删除集合下标为 n 到 m 的元素
6379> zremrangebyrank myzset n m
zremrangebyscore
# 删除集合中顺序 n 到 m 的元素

```

常用命令

服务器相关命令
```shell
select
# 选择数据库，redis数据库 16个(编号0 ~ 15)数据库,登录时默认进入0数据库
6379> select 0 # 选择0数据库

quit/exit
# 退出数据库

dbsize
# 返回当前数据库中key的数目

info 
# 相关信息

config get
# 实时存储收到的请求
6379>  config get * 

flushdb
# 删除当前数据库中的所有key(清空数据库)

flushall
# 删除所有数据库的所有key

```

键值相关命令
```shell
# 键值相关命令(key相当于mysql所有表)

keys
# 返回满足给定pattern的所有key
6379> keys *

exists
# 确定一个key是否存在
6379> exists key_name

del
# 删除一个key
6379> del key_name

expire
# 设置一个key的过期时间
6379> expire key_name time
6379> ttl key_name # 获取这个key的有效时长，-1表示此值已过期

move
# 将当前数据库的key转移到其他数据库中
6379> move age 1 # 将age这个key转移到下标为1数据库(0~15)

persist
# 移除给定key的过期时间
6379> persist key_name 

randomkey
# 随机返回数据库的key

rename 
# 重命名

type
# 返回key的数据类型

ping
# 测试连接是否存活，如果正常连接，返回PONG，否则显示连接失败

echo
# 打印

```


Redis 高级应用
```shell
## 安全性
# 设置客户端连接后进行任何其他指定前需要使用的密码
redis 配置 /etc/redis/6379.conf
修改选项：(默认的密码选项，默认注释)
# requirepass foobared # 可以去掉注释或重新设置新的密码
设置完成，重新启动redis,再次登录，可以登录，但是访问时，报错，没有权限。要授权
6379> auth password_text
也可以在登录时，使用密码 6379> redis-cli -a password_text

## 主从复制
# 通过主从复制，可以允许多个slave server 拥有和 master server 相同的数据库副本
特点：
1、mater 可以拥有多个slave
2、多个slave可以连接同一个master外，还可以连接其它slave
3、主从复制不会阻塞master，同步数据时，master可以继续处理client请求
4、提高系统的伸缩性

过程
1、slave与master建立连接，发送sync同步命令。
2、master会启动一个后台进程，将数据库快照保存到文件中，同时master主进程会开始收集新的写命令并缓存。
3、后台完成保存后，就将此文件发送给slave。
4、slave将此文件保存到硬盘上。

配置主从服务器：
在配置文件中，redis 配置 /etc/redis/6379.conf，找到:
# slaveof <masterip> <masterport> # 指定master的ip和端口
找到：
# masterauth <master-password> # 主机的密码
可以通过info 命令查看本机的role，可以判断主从角色

## 事物处理
Redis只能保证一个client发起的事物中的命令可以连续的执行，而中间不会插入其他client的命令。
当一个client在一个连续中发出 multi 命令时，这个连接会进入一个事务上下文，该连接后续的命令不会立即执行，而是先放到一个队列中，当执行 exec 命令时，redis会顺序的执行队列中的所有命令。
6379> multi # 开启事务，所有操作都将放到 QUEUE 中
6379> exec # 结束事务，顺序执行 QUEUE 所有操作。所有操作有一个不成功，则事务不会回滚(需要改进)
6379> discard # 取消事务，当 执行 multi 时，操作会放到 QUEUE 中，执行取消事务，则放弃 QUEUE 的所有操作。达到事务回滚

乐观锁
# 基于数据版本的记录机制实现的。通过数据库表添加一个version字段，更新时读取version，此版本+1.将提交数据的版本号与数据库表对应记录的当前版本号进行对比，如果提交版本大于数据库版本，则更新，否则认为是过期数据。

Redis乐观锁实例：
watch 命令会监视给定的key,当 exec 时，如果监视的 key 从调用 watch 后发生变化，则整个事务失败。(事务回滚)
可以调用 watch 多次监视多个 key ，这样可以对指定的 key 加乐观锁了。watch 的 key 是对整个连接有效的，事物也是。
断开连接，监听和事物都会被自动清除，exec,discard,unwatch 命令都会清除连接中的所有监视。
6379> watch key_name

## 持久化机制
支持持久化的内存数据库。两种持久化方式
1、snapshotting(快照)也是默认方式
# 将内存中数据以快照的方式写入到二进制文件中，默认的文件名 dump.rdp。可以通过配置设置自动做快照持久化的方式，配置redis在n秒内如果超过m个key被修改就自动做快照
在 /etc/redis/6379.conf 中
save 900 1 # 900秒内如果超过1个key被修改，则发起快照保存
# The filename where to dump the DB(备份位置及文件)
dbfilename dump.rdb

2、Append-only file(缩写 aof)方式
# redis会将每个收到的写命令通过 write 函数追加到文件中，当 redis 重启是会通过重新执行文件中保存的写命令来在内存中重建整个数据库的内容。
# 不会立即写到磁盘中，可以通过配置文件，告诉redis想通过 fsync 函数强制 os 写入到磁盘的时机
#/etc/redis/6379.conf
appendonly yes # 将 no 改成 yes 表示启用 aof 持久化方式

三种写入磁盘的时机
# appendfsync always  # 收到命令就立即写入磁盘，最慢，但是保证完全的持久化
appendfsync everysec # 每秒钟写入磁盘一次，在性能和持久化方面做了很好的折中
# appendfsync no # 完全依赖 os ，性能最好，持久化没保证

## 发布订阅消息(pub/sub)
是一种消息通信模式，目的是解除消息发布者和消息订阅者之间的耦合
订阅这可以通过 subscribe 和 psubscribe 命令向 redis server 订阅自己消息类型，redis 将信息类型称为 通道(channel)。当通过 publish 命令向 redis server 发送特定类型的信息时，订阅该类型的全部 client 都会收到此消息。
# session1
6379> subscribe c1, c2 # 订阅，监听频道
# session2 
6379> publish c1 value # 广播频道c1

## 虚拟内存的使用
和操作系统的虚拟内存不是一个，目的和思路相同。就是暂时把不经常访问的数据从内存交换到磁盘中，从而腾出内存用于其他需要访问的数据。还可以提高数据库容量
#vm配置
vm-enabled yes                  # 开启vm功能
vm-swap-file /tmp/redis.swap    #交换出来的value保存的文件路径
vm-max-memory 1000000           #redis使用的最大内存上限
vm-page-size 32                 #每个页面的大小32字节
vm-pages 134217728              #最多使用多少页面
vm-max-thrcads 4                #用于执行value对象换入的工作线程数量

```



























END