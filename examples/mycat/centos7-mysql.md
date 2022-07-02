## 准备工作
### 安装java---容器
```
yum install java

[root@574ab73efc24 /]# java -version
openjdk version "1.8.0_332"
OpenJDK Runtime Environment (build 1.8.0_332-b09)
OpenJDK 64-Bit Server VM (build 25.332-b09, mixed mode)
```

### 安装wget---容器
```
yum -y install wget
```

### 安装vim---容器
```
yum -y install vim
```
```

# 卸载MariaDB
# 查看版本：
[root@localhost opt]# rpm -qa|grep mariadb
mariadb-libs-5.5.68-1.el7.x86_64
# 卸载
[root@localhost opt]# rpm -e --nodeps mariadb-libs-5.5.68-1.el7.x86_64
# 检查是否卸载干净
rpm -qa|grep mariadb
```

```
rpm -qa | grep mysql

[root@localhost opt]# rpm -qa | grep mysql
mysql80-community-release-el8-4.noarch
[root@localhost opt]# yum remove mysql80-community-release-el8-4.noarch
```

```
[root@localhost opt]# find / -name mysql
/etc/selinux/targeted/active/modules/100/mysql
/usr/lib64/mysql
/usr/share/mysql
[root@localhost opt]# rm -rf /etc/selinux/targeted/active/modules/100/mysql
[root@localhost opt]# rm -rf /usr/lib64/mysql
[root@localhost opt]# rm -rf /usr/share/mysql
[root@localhost opt]# rpm -pa | grep mariadb
```

## 源码安装
### 下载mysql
```
加速站
yum update ca-certificates -y
https://mirrors.tuna.tsinghua.edu.cn/mysql/downloads/MySQL-8.0/mysql-8.0.29-el7-x86_64.tar

wget https://cdn.mysql.com//Downloads/MySQL-8.0/mysql-8.0.29-linux-glibc2.12-x86_64.tar.xz

.tar.gz后缀：tar -zxvf 文件名
.tar.xz后缀：tar -Jxvf 文件名
mv mysql-8.0.29-linux-glibc2.12-x86_64 mysql

mv mysql /usr/local/

cd /usr/local/mysql
mkdir mysqldb

mysql安装目录赋予权限
chmod -R 777 /usr/local/mysql

创建组
[root@localhost mysql]# pwd
/usr/local/mysql

groupadd mysql

创建用户(-s /bin/false参数指定mysql用户仅拥有所有权，而没有登录权限)
useradd -r -g mysql -s /bin/false mysql

将用户添加到组中
chown -R mysql:mysql ./
```
### 修改mysql配置文件 vi /etc/my.cnf
```
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

密码:VkbjIv8dXp(j
2022-07-02T06:42:47.925459Z 6 [Note] [MY-010454] [Server] A temporary password is generated for root@localhost: VkbjIv8dXp(j
```

## rpm安装会各种错误
```
rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022
yum update ca-certificates -y
```

### 启动mysql服务
```
cd /usr/local/mysql/support-files
./mysql.server start

如果第一次启动，当初始化执行会有报错
[root@localhost bin]# cd /usr/local/mysql/support-files
[root@localhost support-files]# ./mysql.server start
Starting MySQL.Logging to '/usr/local/mysql/mysqldb/localhost.localdomain.err'.
 ERROR! The server quit without updating PID file (/usr/local/mysql/mysqldb/localhost.localdomain.pid).

此时不要担心，重新给mysql安装目录赋予一下权限后，再次执行
chmod -R 777 /usr/local/mysql
./mysql.server start
```

### 将mysql添加到系统进程中
```
cp /usr/local/mysql/support-files/mysql.server /etc/init.d/mysqld
此时我们就可以使用服务进程操作mysql了
```
### 设置mysql自启动
```
chmod +x /etc/init.d/mysqld
systemctl enable mysqld

[root@localhost support-files]# chmod +x /etc/init.d/mysqld
[root@localhost support-files]# systemctl enable mysqld
mysqld.service is not a native service, redirecting to /sbin/chkconfig.
Executing /sbin/chkconfig mysqld on

```

### 修改root用户登录密码
```
登录mysql
cd /usr/local/mysql/bin/
./mysql -u root -p

mysql> alter user 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '1234';

```
### 设置允许远程登录
```
mysql> use mysql
mysql> update user set user.Host='%'where user.User='root';
mysql> flush privileges;
mysql> quit

```
### 重启服务且测试
```
systemctl restart mysql	
systemctl status mysql

```

### 查看防火墙开放端口
```
firewall-cmd --list-all

```

### 在防火墙中将3306端口开放
```
firewall-cmd --zone=public --add-port=3306/tcp --permanent
firewall-cmd --reload
//--permanent为永久生效，没有此参数 服务器重启后配置失效

# 关闭防火墙
systemctl stop firewalld.service
# 查看防火墙状态
firewall-cmd --state

```
## yum 安装 会出现各种问题
```
[root@localhost opt]# wget https://repo.mysql.com//mysql80-community-release-el8-4.noarch.rpm
yum localinstall mysql80-community-release-el8-4.noarch.rpm
```

```
yum repolist enabled | grep "mysql.*-community.*"

[root@97982c781d21 opt]# yum repolist enabled | grep "mysql.*-community.*"
Failed to set locale, defaulting to C
mysql-connectors-community/x86_64       MySQL Connectors Community          128
mysql-tools-community/x86_64            MySQL Tools Community                51
mysql80-community/x86_64                MySQL 8.0 Community Server          142

# Failed to set locale, defaulting to C解决方案
方案一：设置系统环境变量
echo "export LC_ALL=en_US.UTF-8"  >>  /etc/profile
source /etc/profile

方案二：设置个人环境变量
echo "export LC_ALL=en_US.UTF-8"  >>  ~/.bashrc
source ~/.bashrc
```

```
yum install mysql-community-server
您可以尝试添加 --skip-broken 选项来解决该问题
您可以尝试执行：rpm -Va --nofiles --nodigest
(尝试添加 '--skip-broken' 来跳过无法安装的软件包 或 '--nobest' 来不只使用最佳选择的软件包) 

yum -y update --skip-broken
```