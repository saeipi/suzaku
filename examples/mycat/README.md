## centos容器 安装mysql8

## yum 安装

### 下载mysql 
```
从官网下载安装包（在Centos7上下载 RH Linux 7 的安装包）
获取官方下载地址 https://dev.mysql.com/downloads/repo/yum/
下载mysql安装包 ：wget https://repo.mysql.com//mysql80-community-release-el8-1.noarch.rpm
[root@97982c781d21 local]# cd opt/
[root@97982c781d21 opt]# wget https://repo.mysql.com//mysql80-community-release-el8-4.noarch.rpm
yum localinstall mysql80-community-release-el8-4.noarch.rpm
```
### 查看是否挂载成功
```
yum repolist enabled | grep “mysql.-community.”
```
### 禁用centos8.0自带的mysql模块
```
yum module disable mysql
```
### 安装mysql
```
yum install mysql-community-server
```


### 安装java
```
yum install java

[root@574ab73efc24 /]# java -version
openjdk version "1.8.0_332"
OpenJDK Runtime Environment (build 1.8.0_332-b09)
OpenJDK 64-Bit Server VM (build 25.332-b09, mixed mode)
```
### 安装wget
```
yum -y install wget
```

### 安装vim
```
yum -y install vim
```
### 更新证书
```
yum update ca-certificates -y
```
### 查看是否有mysql依赖 如果有则卸载
```
rpm -qa | grep mysql

//普通删除模式
rpm -e xxx(mysql_libs)
//强力删除模式,如果上述命令删除时，提示有依赖其他文件，则可以用该命令对其进行强力删除
rpm -e --nodeps xxx(mysql_libs)
```
### 检查是否有mariadb 如果有则卸载
```
rpm -qa | grep mariadb

rpm -e --nodeps mariadb-libs
rpm -e --nodeps mariadb-devel-5.5.65-1.el7.x86_64
```

### 安装mysql依赖包
```
yum install libaio
```

### 下载mysql
```
wget https://mirrors.tuna.tsinghua.edu.cn/mysql/downloads/MySQL-8.0/mysql-8.0.27-el7-x86_64.tar.gz
```
### 解压 更名 移动
```
tar -zxvf mysql-8.0.27-el7-x86_64.tar.gz
mv mysql-8.0.27-el7-x86_64 mysql
mv mysql /usr/local/
```
### 创建数据库文件存放的文件夹。这个文件夹将来存放每个数据库的库文件
```
cd /usr/local/mysql
mkdir mysqldb
```

### mysql安装目录赋予权限
```
chmod -R 777 /usr/local/mysql
```

### 修改mysql配置文件
```
vi /etc/my.cnf

[mysqld]
# 设置3306端口
port=3306
# 设置mysql的安装目录
basedir=/usr/local/mysql
# 设置mysql数据库的数据的存放目录
datadir=/usr/local/mysql/mysqldb
# 允许最大连接数
max_connections=10000
# 允许连接失败的次数。这是为了防止有人从该主机试图攻击数据库系统
max_connect_errors=10
# 服务端使用的字符集默认为UTF8
character-set-server=utf8
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
# 默认使用“mysql_native_password”插件认证
default_authentication_plugin=mysql_native_password
[mysql]
# 设置mysql客户端默认字符集
default-character-set=utf8
[client]
# 设置mysql客户端连接服务端时默认使用的端口
port=3306
default-character-set=utf8
```

### 安装mysql
```
cd /usr/local/mysql/bin/
./mysqld --initialize --console
```

### 参考
```
https://www.jianshu.com/p/60a64203ca13
https://blog.csdn.net/weixin_42326851/article/details/123984601
https://www.cnblogs.com/chenjiangbin/p/16049726.html
https://blog.csdn.net/m0_67401835/article/details/123868011
```

### 文章目录

- [前言](https://blog.csdn.net/weixin_42326851/article/details/123984601#_4)
- [一、卸载MariaDB](https://blog.csdn.net/weixin_42326851/article/details/123984601#MariaDB_16)
- - [1.1 查看版本：](https://blog.csdn.net/weixin_42326851/article/details/123984601#11__22)
    - [1.2 卸载](https://blog.csdn.net/weixin_42326851/article/details/123984601#12__26)
    - [1.3 检查是否卸载干净：](https://blog.csdn.net/weixin_42326851/article/details/123984601#13__30)
- [二、安装MySQL](https://blog.csdn.net/weixin_42326851/article/details/123984601#MySQL_35)
- - [2.1 下载资源包](https://blog.csdn.net/weixin_42326851/article/details/123984601#21__36)
    - - [2.1.1 官网下载](https://blog.csdn.net/weixin_42326851/article/details/123984601#211__37)
        - [2.1.2 wget下载](https://blog.csdn.net/weixin_42326851/article/details/123984601#212_wget_44)
    - [2.2 解压](https://blog.csdn.net/weixin_42326851/article/details/123984601#22__49)
    - [2.3 重命名](https://blog.csdn.net/weixin_42326851/article/details/123984601#23__54)
    - [2.4 添加PATH变量](https://blog.csdn.net/weixin_42326851/article/details/123984601#24_PATH_62)
- [三、用户和用户组](https://blog.csdn.net/weixin_42326851/article/details/123984601#_71)
- - [3.1 创建用户组和用户](https://blog.csdn.net/weixin_42326851/article/details/123984601#31__72)
    - [3.2 数据目录](https://blog.csdn.net/weixin_42326851/article/details/123984601#32__83)
- [四、初始化MySQL](https://blog.csdn.net/weixin_42326851/article/details/123984601#MySQL_95)
- - [4.1 配置参数](https://blog.csdn.net/weixin_42326851/article/details/123984601#41__96)
    - [4.2 初始化](https://blog.csdn.net/weixin_42326851/article/details/123984601#42__174)
- [五、启动MySQL](https://blog.csdn.net/weixin_42326851/article/details/123984601#MySQL_184)
- - [5.1 启动服务](https://blog.csdn.net/weixin_42326851/article/details/123984601#51__187)
    - [5.2 登录](https://blog.csdn.net/weixin_42326851/article/details/123984601#52__199)
    - [5.3 修改密码](https://blog.csdn.net/weixin_42326851/article/details/123984601#53__208)
    - [5.4 设置允许远程登录](https://blog.csdn.net/weixin_42326851/article/details/123984601#54__216)
    - [5.5 在Navicat上测试连接](https://blog.csdn.net/weixin_42326851/article/details/123984601#55_Navicat_224)
- [总结](https://blog.csdn.net/weixin_42326851/article/details/123984601#_228)

* * *

# 前言

- **[MySQL查看表占用空间大小](https://blog.csdn.net/weixin_42326851/article/details/124213228)**
- **[CentOS7 环境下MySQL常用命令](https://blog.csdn.net/weixin_42326851/article/details/124209898)**
- **[MySQL: 范围查询优化](https://blog.csdn.net/weixin_42326851/article/details/124993822?spm=1001.2014.3001.5501)**

**环境介绍 :**  
**服务器：** 阿里云轻量应用服务器  
**系统版本：** CentOS 7  
**MySQL版本：** 8.0

* * *

# 一、卸载MariaDB

> **在CentOS中默认安装有MariaDB，是MySQL的一个分支，主要由开源社区维护。  
> CentOS 7及以上版本已经不再使用MySQL数据库，而是使用MariaDB数据库。  
> 如果直接安装MySQL，会和MariaDB的文件冲突。  
> 因此，需要`先卸载自带的MariaDB，再安装MySQL。`**

## 1.1 查看版本：

```shell
rpm -qa|grep mariadb
```

## 1.2 卸载

```shell
rpm -e --nodeps 文件名
```

## 1.3 检查是否卸载干净：

```shell
rpm -qa|grep mariadb
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/d86f3bf8ebbf4ed5826f49f00460290b.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAbGZ3aA==,size_20,color_FFFFFF,t_70,g_se,x_16)

# 二、安装MySQL

## 2.1 下载资源包

### 2.1.1 官网下载

**MySQL官网下载地址 :**

```html
https://dev.mysql.com/downloads/mysql/
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/19b0ecbd2ffe4a33b8838a0c26ecd172.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAbGZ3aA==,size_20,color_FFFFFF,t_70,g_se,x_16)

### 2.1.2 wget下载

```shell
wget https://cdn.mysql.com//Downloads/MySQL-8.0/mysql-8.0.29-linux-glibc2.12-x86_64.tar.xz

wget https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.20-linux-glibc2.12-x86_64.tar.xz
```

## 2.2 解压

```shell
.tar.gz后缀：tar -zxvf 文件名
.tar.xz后缀：tar -Jxvf 文件名
```

## 2.3 重命名

将解压后的文件夹重命名（或者为文件夹创建软链接）

```shell
# 重命名
mv 原文件夹名 mysql8
# 软链接
ln -s 文件夹名 mysql8
```

## 2.4 添加PATH变量

添加PATH变量后，可在全局使用MySQL。

- **有两种添加方式：export命令临时生效、修改配置文件用久生效；**

```shell
#临时环境变量，关闭shell后失效，通常用于测试环境
export PATH=$PATH:/data/software/mysql8/bin
```

# 三、用户和用户组

## 3.1 创建用户组和用户

```shell
# 创建一个用户组：mysql
groupadd mysql
# 创建一个系统用户：mysql，指定用户组为mysql
useradd -r -g mysql mysql
```

- **创建用户组：**`groupadd`
- **创建用户：**`useradd`  
    `-r`**：创建系统用户**  
    `-g`**：指定用户组**

## 3.2 数据目录

**1、创建目录**

```shell
mkdir -p /data/software/mysql8/datas
```

**2、赋予权限**

```shell
# 更改属主和数组
chown -R mysql:mysql /data/software/mysql8/datas
# 更改模式
chmod -R 750 /data/software/mysql8/datas
```

# 四、初始化MySQL

## 4.1 配置参数

> 在`/data/software/mysql8/`下，创建`my.cnf`配置文件，用于**初始化MySQL数据库**

```shell
[mysql]
# 默认字符集
default-character-set=utf8mb4
[client]
port       = 3306
socket     = /tmp/mysql.sock

[mysqld]
port       = 3306
server-id  = 3306
user       = mysql
socket     = /tmp/mysql.sock
# 安装目录
basedir    = /data/software/mysql8
# 数据存放目录
datadir    = /data/software/mysql8/datas/mysql
log-bin    = /data/software/mysql8/datas/mysql/mysql-bin
innodb_data_home_dir      =/data/software/mysql8/datas/mysql
innodb_log_group_home_dir =/data/software/mysql8/datas/mysql
#日志及进程数据的存放目录
log-error =/data/software/mysql8/datas/mysql/mysql.log
pid-file  =/data/software/mysql8/datas/mysql/mysql.pid
# 服务端使用的字符集默认为8比特编码
character-set-server=utf8mb4
lower_case_table_names=1
autocommit =1
 
 ##################以上要修改的########################
skip-external-locking
key_buffer_size = 256M
max_allowed_packet = 1M
table_open_cache = 1024
sort_buffer_size = 4M
net_buffer_length = 8K
read_buffer_size = 4M
read_rnd_buffer_size = 512K
myisam_sort_buffer_size = 64M
thread_cache_size = 128
  
#query_cache_size = 128M
tmp_table_size = 128M
explicit_defaults_for_timestamp = true
max_connections = 500
max_connect_errors = 100
open_files_limit = 65535
   
binlog_format=mixed
    
binlog_expire_logs_seconds =864000
    
# 创建新表时将使用的默认存储引擎
default_storage_engine = InnoDB
innodb_data_file_path = ibdata1:10M:autoextend
innodb_buffer_pool_size = 1024M
innodb_log_file_size = 256M
innodb_log_buffer_size = 8M
innodb_flush_log_at_trx_commit = 1
innodb_lock_wait_timeout = 50
transaction-isolation=READ-COMMITTED
      
[mysqldump]
quick
max_allowed_packet = 16M
       
[myisamchk]
key_buffer_size = 256M
sort_buffer_size = 4M
read_buffer = 2M
write_buffer = 2M
        
[mysqlhotcopy]
interactive-timeout
```

## 4.2 初始化

```shell
mysqld --defaults-file=/data/software/mysql8/my.cnf --basedir=/data/software/mysql8/ --datadir=/data/software/mysql8/datas/mysql --user=mysql --initialize-insecure
```

**参数（重要）**

- `defaults-file`：指定配置文件（要放在–initialize 前面）
- `user`： 指定用户
- `basedir`：指定安装目录
- `datadir`：指定初始化数据目录
- `intialize-insecure`：初始化无密码

# 五、启动MySQL

查看 MySQL的 bin路径下，是否包含`mysqld_safe`，用于后台安全启动MySQL。  
![在这里插入图片描述](https://img-blog.csdnimg.cn/edd866433afb462087d105a10fad2811.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAbGZ3aA==,size_20,color_FFFFFF,t_70,g_se,x_16)

## 5.1 启动服务

```shell
# 完整命令
/data/software/mysql8/bin/mysqld_safe --defaults-file=/data/software/mysql8/my.cnf &
# 添加PATH变量后的命令（省略bin目录的路径）
mysqld_safe --defaults-file=/data/software/mysql/my.cnf &
```

**查看是否启动**

```shell
ps -ef|grep mysql
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/b652504a24f1487abd0b82348988995a.png)

## 5.2 登录

```shell
# 无密码登录方式
/data/software/mysql8/bin/mysql -u root --skip-password
# 有密码登录方式（初始的随机密码在/data/mysql8_data/mysql/mysql.log下）
mysql -u root -p
password:随机密码
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/0834429f9640415dbfc963036e51fd0e.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAbGZ3aA==,size_20,color_FFFFFF,t_70,g_se,x_16)

## 5.3 修改密码

```shll
# 修改密码
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
# 刷新权限
FLUSH PRIVILEGES;
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/3a979a1e03e64111aa2722b3df4e297c.png)

## 5.4 设置允许远程登录

**登录到mysql里执行**

```shell
mysql> use mysql
mysql> update user set user.Host='%'where user.User='root';
mysql> flush privileges;
mysql> quit
```

## 5.5 在Navicat上测试连接

![连接成功](https://img-blog.csdnimg.cn/893a021fa2b348a8a66d6c36f3042c25.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAbGZ3aA==,size_20,color_FFFFFF,t_70,g_se,x_16)

* * *

# 总结

> 如果此篇文章有帮助到您, 希望打大佬们能`关注、点赞、收藏、评论`支持一波，非常感谢大家！  
> 如果有不对的地方请指正!!!