#### MySQL

# 安装
```shell
# 查看是否安装 mysql
rpm -qa | grep mysql
# 查看 yum 上提供的 mysql 可下载的版本
yum list | grep mysql
# 安装 MySQL
> CentOS 7 没有 mysql-server 文件，需要去官网下载
wget http://dev.mysql.com/get/mysql-community-release-el7-5.noarch.rpm
rpm -ivh mysql-community-release-el7-5.noarch.rpm
> 安装
yum install mysql-community-server

# 修改配置
sudo vim /etc/my.cnf
```

# 启动 与 退出
```shell
# 验证是否启动
sudo netstat -tap | grep mysql
# 启动 MySQL 服务
sudo service mysqld start
# 使用 root 用户登录
mysql -u root
# 退出
quit/exit
```

# 使用
```shell

# SQL 语句部分大小写
# 创建数据库 CREATE DATABASE <数据库名>；  前面的可以小写
CREATE DATABASE mysql_test;
# 删除数据库 DROP DATABASE <数据库名>
drop database mysql_test;
# 查看数据库(必须要有 ; )
show databases；
# 连接数据库， 语句格式 use <数据库名>
use information_schema
# 查看表
show tables;
# 创建数据表
CREATE TABLE 表的名字
(
    列名 数据类型(数据长度)，
    ...
);
# 重命名 一张表
rename table 原表名 to 新表名;
alter table 原表名 rename 新表名;
alter table 原表名 rename to 新表名;
# 删除一张表
drop table 表名;
# 增加一列
alter table 表名 add count 列名 数据类型 约束 alter 已有的某列名; # count 可以省略 后面的 'alter 已有的某列名' 表示新加的某列放在 '已有的某列名' 这一列后面，省略表示加到最后一列，'alter 已有的某列名' 改成 'first' 表示放在第一列
alter table tab_a add test int(10) default 11 first;
# 删除一列
alter table 表名 drop count 列名; # count 可以省略
# 重命名一列
alter table 表名 change 原列名 新列名 数据类型 约束; # 数据类型不能省略，否则重命名失败，原列名和新列名相同，指向新的数据类型或约束则是修改数据类型或约束
# 修改数据类型
alter table 表名 modify 列名 新数据类型;


# 查看一张表的各列的数据类型
select column_name,data_type from information_schema.columns where table_name='表名';

# 查找
select * from 表名;
# 插入 (全部的列名都要插入值时，可省略)
insert into 表名(列名a，列名b ...) vales(值1，值2 ...);
# mysql> 下执行某个sql 文件
source /home/xxx/xxx/xxx.sql
```

# 数据类型
```
CHAR 和 VARCHAR 的区别: 
CHAR 的长度是固定的，而 VARCHAR 的长度是可以变化的，比如，存储字符串 “abc"，对于 CHAR(10)，表示存储的字符将占 10 个字节(包括 7 个空字符)，而同样的 VARCHAR(12) 则只占用3个字节的长度，12 只是最大值，当你存储的字符小于 12 时，按实际长度存储。
CHAR 和 VARCHAR 的区别: 
CHAR 的长度是固定的，而 VARCHAR 的长度是可以变化的，比如，存储字符串 “abc"，对于 CHAR(10)，表示存储的字符将占 10 个字节(包括 7 个空字符)，而同样的 VARCHAR(12) 则只占用3个字节的长度，12 只是最大值，当你存储的字符小于 12 时，按实际长度存储。
```

# SQL 的约束
```
----------------------------------------------------------
约束类型 | 主键         | 默认值   | 唯一          | 非空
关键字   | PRIMARY KEY | DEFAULT | FOREIGN KEY  | NOT NULL
-----------------------------------------------------------
# 约束在创建表时定义
# 主键约束[PRIMARY KEY]
id INT(10) PRIMARY KEY, # 主键
CONSTRAINT id_pk PRIMARY KEY(id) # 这是在创建表末尾加的语句，id_pk 为自定义的主键名称，id 为上面定义的列名，把它作为主键
CONSTRAINT id_pk PRIMARY KEY(id,name) # 复合主键，id和name两列

# 默认值约束[DEFAULT],当插入为空是，会有默认值
xxx int(10) DEFAULT '10',

# 唯一约束[UNIQUE]，不能重复
UNIQUE(列名),

# 外键约束[FOREIGN KEY]，一个表可以有多个外键，每个外键必须 REFERENCES(参考)另一个表的主键，被外键约束的列，取值必须在它参考列中有对应的值
CONSTRAINT emp_fk FOREIGN KEY(in_dpt) REFERENCES table_xxx(dpt_name), # emp_fk自定义的外键名，in_dpt 作为外键的列，table_xxx(dpt_name) 参考的 表 table_xxx 的 dpt_name 列，如果 插入是 dpt_name 的值不存在 table_xxx，则插入失败

# 非空约束[NOT NULL],插入该值，不能为空
```

select 语句(查)
```
# 1、格式
select 要查询的列名 from 表名 where 限制条件; 

# 以下都是 where + 限制

# 2、数学符号条件 [=、<、>、>=、<=]
select * from tab_xxx where age>25; # 年龄大于25 的

# 3、"AND" 与 "OR" [OR(或) AND(且)]
select * from tab_xxx where age<25 or age>30; # 年龄小于25 或 大于30
select * from tab_xxx where age between 25 and 30; # 年龄在 25 和 30 之间

# 4、IN 和 NOT IN [在 和 不在 某个范围内的结果]
select * from tab_xxx where xxx in ('aaa','bbb'); # 列xxx 的值 属于 'aaa' 或 'bbb' 
... not in ... # 不属于 ...

# 5、通配符 [_代表一个未指定的字符，%代表不定个未指定字符]
select * from tab_xxx where phone like '186%' # 电话号码是186 开头的

# 6、排序 [order by 排序关键字，asc 升序，desc 降序，默认升序，从小到大]
select * from tab_xxx order by salary asc; # 按薪水的多少升序排序

# 7、内置函数和计算
函数名 | count | sum | avg    | max  | min
作用   | 计数  | 求和 | 求平均值 |最大值 | 最小值
> count 函数可用于任何数据类型，其他4个函数只能对数字类型进行计算
select max(salary) as max_salary, min(salary) as min_salary from tab_xxx; # as 关键字可给值重命名

# 8、子查询
> 可以进行 select 嵌套

# 9、连接查询
> 处理多个表时，需要显示两个或多个表的数据时，使用 连接(join .. on) 操作，连接思想就是把两个或多个表当成一个新的表来处理
select id, name, people_num from tab_a, tab_b where tab_a.xxxa = tab_b.xxxb;
select id, name, people_num from tab_a join tab_b on tab_a.xxxa = tab_b.xxxb; # 这个 sql 等价于上面的 sql
```

updata(改)
```
update 表名 set 列1=值1，列2=值2 where 条件;
update tab_b set age=13 where id=1;
```

delete(删)
```
delete from 表名 where 条件;
delete from tab_b where name='xxx';
```

索引
```shell
# 加快查询速度, 在查询时，where里面的条件，会自动判断有没有可用的索引 v)

# 方式一
> alter table 表名 add index 索引名(列名);
alter table tab_b add index idx_id(id); # 在表上 id 列上建立名为 idx_id 的索引

# 方式二
> create index 索引名 on 表名(列名);
create index idx_name on tab_b(name);

# 查看索引
show index from 表名;
```

视图
```shell
# 虚拟存在的表，提供专门的数据。
# 数据库只存放视图的动员，没有存放视图中的数据，查询视图会从数据库原来表中取出对应数据
# 视图数据依赖原来表的数据，一旦原表数据发生改变，视图查询出的数据也是改变后的数据
> create view 视图名(列a,列b,列c) as select 列1，列2，列3 from 表名;
create view v_tab(v_name,v_age) as select name, age from tab_b; 
```

导入 与 导出
```shell
# 导入
> load data infile '文件路径' into table 表名; # txt 文件
load data infile '/tmp/data/in.txt' into table tab_b(列1...); # 默认可以不注明列名

# 导出
# 选择导出会报错: ERROR 1290 (HY000): The MySQL server is running with the --secure-file-priv option so it cannot execute this statement
# 报错原因：secure-file-priv 设置指定了目录，需要在指定目录下进行数据导出
mysql> show variables like '%secure%'; # 可查询 secure-file-priv 指定的目录。例如：'/var/lib/mysql-files/'

# secure_file_priv参数说明
# 这个参数用来限制数据导入和导出操作的效果，例如执行LOAD DATA、SELECT ... INTO OUTFILE语句和LOAD_FILE()函数。这些操作需要用户具有FILE权限。
# 如果这个参数为空，这个变量没有效果；
# 如果这个参数设为一个目录名，MySQL服务只允许在这个目录中执行文件的导入和导出操作。这个目录必须存在，MySQL服务不会创建它；
# 如果这个参数为NULL，MySQL服务会禁止导入和导出操作。这个参数在MySQL 5.7.6版本引入。

> select 列1，列2 from 表名 into outfile '导出文件(目录/var/lib/mysql-files/下)';
select name, age from tab_b into outfile '/var/lib/mysql-files/1.txt';
```

备份 与 恢复
```shell
# 备份是把数据库的结构，包括数据、约束、索引、视图等全部另存为一个文件
> mysqldump -u root 数据库名>备份文件名 # 备份整个数据库
> mysqldump -u root 数据库名 表名>备份文件名 # 备份表
mysqldump -u root xxx_db>bak.sql

# 恢复数据库 (加入恢复到 test 数据库)
mysql -u root test < bak.sql
```

mysql 的一些操作
```shell
# 查看 mysql 版本
select version();
# 查看当前时间和日期
select now();
# 充当计算器
select sin(pi()/4);

# 如果输入错误，不想执行当前语句，则在后面加 \c
mysql> xxxx \c

# 查询当前数据库
select database();

# 查看表结构
> select column_name,data_type from information_schema.columns where table_name='表名'; 只是查看表列对应的数据类型
describe tab_name; # 查看整张表的结构
```

auto_increment
```shell
# 自增
create table xxx(
    id mediumint not null auto_increment,
    primary key (id)
); # id 自增
> auto_increment 语句生成的起始值不是1，通过create table 或 alter table 来设置
alter table tab_b auto_increment = 100;
```

mysql 数据类型
```
[NATIONAL] CHAR(M) [BINARY| ASCII | UNICODE] 
固定长度字符串，当保存数据时在右侧自动填充空格以达到指定的长度。M表示列长度。M的范围是0到255个字符。 注：当检索CHAR类型值时尾部空格将被删除。 如果想要将某个CHAR的长度设为大于255，执行CREATE TABLE或ALTER TABLE语句时将失败并提示错误，尝试输入：
 mysql> CREATE TABLE c1 (col1 INT, col2 CHAR(500));
* mysql> SHOW CREATE TABLE c1;
*   CHAR是CHARACTER的简写。NATIONAL CHAR(或其等效短形式NCHAR)是标准的定义CHAR列应使用默认字符集的SQL方法。BINARY属性是指定列字符集的二元校对规则的简写。排序和比较基于数值字符值。列类型CHAR BYTE是CHAR BINARY的一个别名，这是为了保证兼容性。通过指定latin1字符集，可以为CHAR指定ASCII属性。通过指定ucs2字符集可以为CHAR指定UNICODE属性。 MySQL允许创建类型为CHAR(0)的列。这主要用于与必须有一个列但实际上不使用它的值的旧版本中应用程序相兼容。 
* CHAR 这是CHAR(1)的同义词。 

[NATIONAL] VARCHAR(M) [BINARY] 
变长字符串。M表示最大列长度。M的范围是0到65,535。(VARCHAR的最大实际长度由最长的行的大小和使用的字符集确定。最大有效长度是65,532字节）。 注：MySQL 5.1遵从标准SQL规范，并且不自动移除VARCHAR值的尾部空格。 VARCHAR是字符VARYING的简写。BINARY属性是指定列的字符集的二元校对规则的简写。排序和比较基于数值字符值。VARCHAR保存时用一个字节或两个字节长的前缀+数据。如果VARCHAR列声明的长度大于255，长度前缀是两个字节。 

BINARY(M) 
BINARY类型类似于CHAR类型，但存储二进制字节字符串而不是非二进制字符串。 

VARBINARY(M) 
VARBINARY类型类似于VARCHAR类型，但存储二进制字节字符串而不是非二进制字符串。 

TINYBLOB
 最大长度为255字节的BLOB列。 

TINYTEXT 
最大长度为255字符的TEXT列。 
BLOB[(M)]
 最大长度为65,535字节的BLOB列。 可以给出该类型的可选长度M。如果给出，则MySQL将列创建为最小的但足以容纳M字节长的值的BLOB类型。 

TEXT[(M)] 
最大长度为65,535字符的TEXT列。 可以给出可选长度M。则MySQL将列创建为最小的但足以容纳M字符长的值的TEXT类型。 

MEDIUMBLOB 
最大长度为16,777,215字节的BLOB列。 

MEDIUMTEXT
 最大长度为16,777,215字符的TEXT列。 

LONGBLOB
 最大长度为4,294,967,295或4GB字节的BLOB列。LONGBLOB列的最大有效(允许的)长度取决于客户端或服务器协议中配置最大包大小和可用的内存。 

LONGTEXT
 最大长度为4,294,967,295或4GB字符的TEXT列。LONGTEXT列的最大有效(允许的)长度取决于客户端或服务器协议中配置最大包大小和可用的内存。 

ENUM('value1','value2',...) 枚举类型。一个字符串只能由一个值，从值列列表'value1'，'value2'，...，NULL中或特殊''错误值中选出。ENUM列最多可以有65,535个截然不同的值。ENUM值在内部用整数表示。 

SET('value1','value2',...) 
一个SET类型列可以64个不同值。字符串对象可以有零个或多个值，每个值必须来自列表值'value1'，'value2'，...，SET对应的值在内部用整数表示。 
```

比较函数和操作符
```
比较运算产生的结果：1(true)、0(false)、null

<=>
空值(null)安全的等号，

<>\!=
不等于

is true/false、is not true/false

coalesce(value...)
返回参数列表中第一个非null值，没有非null，则返回null

greatest(value,value,...)
返回最大值

interval(N,N1,N2...)
取连续的最后一个大于 N 的参数，N表示第0位，例:
select interval(3,0,1,2,3,4);# 结果为4,即索引4(参数是3)

```

触发器
```
# 触发器(触发程序)是与表有关的固定的数据对象，当表上出现特定事件，将激活该对象。一般用于检查给表插入新的值或者进行表内的数值计算之类的更新。

#语法
> create trigger trigger_name trigger_time trigger_event on tbl_name for each row trigger_stmt

# tbl_name表，不能将触发程序与临时表或视图关联起来
# trigger_time: 触发时间，可以是 before 或 after,在激活触发器的语句之前或之后触发 。
# trigger_event:激活触发器的类型
  -> insert: 插入数据时激活触发器，通过 insert、load data、replace 语句
  -> update: 更新数据时激活触发器,通过 update 语句
  -> delete: 删除数据时激活触发器，通过delete、replace 语句
  某个表触发时间和时间相同的触发器只能有一个，例如 before update
# trigger_stmt: 当触发器激活时执行的语句。如果执行多个语句，可使用 begin ... end
例：test.sql
...
delimiter |
    create trigger testref before insert on test1
    for each row begin
        insert into test2 set a2 = new.a1;
        ...
    end
|
delimiter ;
...
使用别名 old 和 new 能够引用触发器相关的表中的列。在INSERT触发程序中，仅能使用NEW.col_name，没有OLD.col_name。在DELETE触发程序中，仅能使用OLD.col_name，没有NEW.col_name。在UPDATE触发程序中，可以使用OLD.col_name来引用更新前的某一行的列，也能使用NEW.col_name来引用更新后的行中的列。

> drop trigger [schema_name.]trigger_name
舍弃触发器。方案名称(schema_name) 可选。

```

视图
```
# 创建视图
create [or replace] [algorithm = {undefined | merge | temptable}]
  view view_name [(column_list)] as select_statement
  [with [cascaded | local] check option]
# or replace: 改语句替换已有的视图，可选
# select_statement: 给出视图的定义，从基表或其他视图进行选择
# 视图是属于数据库的，db_name.view_name

# 修改视图
alter [algorithm = {undefined | merge | temptable}]
  view view_name [(column_list)] as select_statement
  [with [cascaded | local] check option]

# 删除视图
drop view [if exists] view_name [,view_name] ... [restrict | cascade]

```

# 存储过程和函数
```shell
# TODO
```

# mysql 权限管理
```shell
# 账户权限信息被存储在mysql数据库中的 user、db、host、bles_priv、columns_priv、procs_priv 表中。

# TODO

#设置账户密码
# 用 mysqladmin 命令设置
mysqladmin -u user_name -h host_name password 'xxx'


```


END