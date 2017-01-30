### Mongodb
```
# 介于关系数据库和非关系数据库之间。支持的数据结构非常松散，类似json 的bson格式，可以存储比较复杂的数据类型。
# Mongodb 一个数据有多个集合(集合等价于mysql的表)，一个集合有多个文档(文档等价于mysql的表每一行数据)
```

安装 与 使用
```shell
# Mongodb 提供 centos yum 安装方式(例：安装 mongodb 3.4)
vi /etc/yum.repos.d/mongodb-org-3.4.repo # 新建mongodb-org-3.4.repo，输入以下内容
[mongodb-org-3.4]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/3.4/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-3.4.asc

# 安装 mongdb
yum install -y mongodb-org

# 配置文件
/etc/mongod.conf

# 启动
service mongod start
# 停止
service mongod stop
mongod --shutdown
# 重启
service mongod restart
# 增加开机启动
chkconfig mongod on

# 登录
mongo
# 退出
exit
# 查看数据库
show dbs
# 连接指定的数据库
use test_db
# 查看当前连接的数据库
db
# 创建数据库(使用一个不存在的数据库，则是创建数据库)
use mydb # 这时$show dbs 时，因为mydb为空，不显示mydb 或者显示 mydb(empty)
# 销毁数据库
use local # 先连接数据库
db.dropDatabase()

```

元数据
```
# 数据库的信息存储在集合中，他们统一使用系统的命名空间：DBNAME.system.* DBNAME可用db或数据库名替代

DBNAME.system.namespaces ：列出所有名字空间
DBNAME.system.indexs ：列出所有索引
DBNAME.system.profile ：列出数据库概要信息
DBNAME.system.users ：列出访问数据库的用户
DBNAME.system.sources ：列出服务器信息
```

操作
```shell
use mydb
# 创建集合
db.createCollection(name,options)#options 是一个作为初始化的文档(可选)
# 查看集合
show collections
# 删除集合(删除 users 集合)
db.users.drop()

# 插入数据 
# 使用 insert()
> 插入数据时，如果集合不存在，则会自动创建
db.users.insert([
        {name : "aa", email : "aa@github.com"},
        {name : "bb", email : "bb@github.com"},
        ...
    ])
# 使用 save()
> 插入数据时，如果集合不存在，则会自动创建
db.users.save([
        {name : "aa", email : "aa@github.com"},
        {name : "bb", email : "bb@github.com"},
        ...
    ])

# 查询数据
# find()
> db.collection_name.find()
db.users.find()
# pretty()
> 使输出更美观,json 格式
db.user.find().pretty()

# and
> 当 find() 中传入多个键值对时，MongoDB 就会将其作为 AND 查询处理
db.db_name.find({key1:value1,key2:value2}).pretty()
 
# or
> or 查询语句以 $or 作为关键词，满足多个键值对其中一个即可被查出
db.post.find({
    $or:[{key1:value1},{key1:value1}]
    }).pretty()

#可同时使用and 和 or
db.post.find({
    "likes":{$gt:10}
    $or:[...]
    }).pretty()
# {$gt:10} 表示大于10，
# $lt 表示小于
# $lte 表示小于等于
# $gte 表示大于等于
# $ne 表示不等于

# {$type:n}
# 双精度型-1/字符串-2/对象-3/数组-4/二进制数据-5/对象ID-7/布尔类型-8/数据-9
# 空-10/正则表达式-11/JS代码-13/符号-14/有作用域的JS代码-15/32位整型数-16
# 时间戳-17/64位整型数-18/Min key-255/Max key-127

# 读取指定数量的数据记录 limit(n)
db.collection_name.find().limit(n) # 读取n条记录
# 读取时跳过指定数量的数据记录 skip(n)
db.collection_name.find().skip(n) # 前n 条不读取
# 排序，sort({key:1|-1}),升序用1，降序用-1
db.collection_name.find().sort({key:1})

# 更新
> 例如将 user 集合 中的 user_id = 1 的 文档 的 name 更新为 abc
db.user.update({"user_id" : 1},{$set:{"name" : "abc"}})
#              查找条件，可多个     更新内容，可多个
# 发现上面的只能更新第一条条件符合的数据，只对一个文档更新，如果想作用所有文档，加入 multi:true
db.user.update({"user_id" : 1},{$set:{"name" : "abc"}},{multi : true})
# 更换已存在的文档，没有 update 好用
db.user.save({...})

# 删除
> 删除所有条件符合的文档
db.user.remove({"user_id" : 2})

```

索引
```shell
# 语法
db.collection_name.ensureIndex({key:1|-1})
# 1 代表升序，-1 代表降序

ensureIndex() 可选参数
# 参数                类型            描述
background          boolean         建立的索引要不要阻塞其他数据库操作，默认false
unique              boolean         建立的索引是否唯一，默认false
name                string          索引的名称，未指定系统自定生成
dropDups            boolean         建立唯一索引时，是否删除重复记录，默认false
sparse              boolean         对文档不存在的字段数据不启用索引，默认false
expireAfterSeconds  integer         设置集合的生存时间，单位为秒
v                   index version   索引的版本号
weights             document        索引权重值，范围为1到99999
default-language    string          默认为英语
language_override   string          默认值为 language
# 例
db.user.ensureIndex({"user_id" : 1},{background:1})
```

聚合
```shell
# 语法
db.collection_name.aggregate([
    {$match:{key:value}},
    {$limit:num},
    {$group:{new_k:new_v}}
    ])

操作
# $macth: 查询，和 find() 一样
# $limit: 限制数量
# $skip: 跳过数量
# $sort: 排序
# $group: 按照给定表达式组合结果,将集合中的文档分组，可用于统计结果
# $project:修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。
# $unwind: 将文档中的某一个数组类型字段拆分成多条，每条包含数组中的一个值。
# $geoNear：输出接近某一地理位置的有序文档

例:
db.user.aggregate([{$group:{_id:"$name",user:{$sum:"user_id"}}}])
# $name 是取key为name的值

聚合表达式
# 名称        描述
 $sum        计算总和
 $avg        计算平均值
 $min和$max   计算最小和最大值
 $push       在结果文档中插入值到一个数组
 $addToSet   在结果文档中插入值到一个数组，但不创建副本
 $first      根据资源文档的排序获取第一个文档数据
 $last       根据资源文档的排序获取最后一个文档数据

```

END