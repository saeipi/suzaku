# **二，下载安装包**

1，从官网下载安装包（在Centos7上要下载 RH Linux 7 的安装包）

[https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)

mysql-8.0.17-1.el7.x86\_64.rpm-bundle.tar

2，清理环境

2.1 查看系统是否已经安装了mysql数据库

```
rpm -qa | grep mysql
```

2.2 将查询出的文件逐个删除，如

```
yum remove mysql-community-common-5.7.20-1.el6.x86_64
```

2.3 删除mysql的配置文件

```
find / -name mysql
```

2.4 删除配置文件

```
rm -rf /var/lib/mysql
```

2.5删除MariaDB文件

```
rpm -pa | grep mariadb 删除查找出的相关文件和目录，如 yum -y remove mariadb-libs.x86_64
```
2.6 清除yum缓存 && 重新建立缓存
```
yum clean all
yum makecache
```
2.7 更新密钥
```
rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022

```
3，安装

3.1解压

```
tar -xf mysql-8.0.17-1.el7.x86_64.rpm-bundle.tar
```

3.2安装

```
yum install mysql-community-{client,common,devel,embedded,libs,server}-*
```

等待安装成功！

4，配置

4.1 启动mysqld服务，并设为开机自动启动。命令：

```
systemctl start mysqld.service //这是centos7的命令systemctl enable mysqld.service
```

4.2 通过如下命令可以在日志文件中找出密码：

```
grep "password" /var/log/mysqld.log
```

4.3按照日志文件中的密码，进入数据库

```
mysql -uroot -p
```

4.4设置密码（注意Mysql8密码设置规则必须是大小写字母+特殊符号+数字的类型）

```
ALTER USER 'root'@'localhost' IDENTIFIED BY 'new password';
```

4.5开启远程访问

```
use mysql； ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '你的密码'; FLUSH PRIVILEGES;
```

4.6更改语言

打开/etc/my.cnf

添加如下语句

```
[client]default-character-set=utf8 ...... character-set-server=utf8collation-server=utf8_general_ci
```

保存

4.7重启

```
systemctl restart mysqld
```

4.8重新登录mysql，查看status

![](/images/53/99de5b177e5946edb10fae006c1929ed.png)

5，可以在windows上用Navicat远程登录mysql了。