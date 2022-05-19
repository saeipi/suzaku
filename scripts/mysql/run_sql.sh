#!/usr/bin/env bash
MYSQL_USERNAME="root"
MYSQL_PASSWORD=""
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=3306
MYSQL_DB="sdb"

folder="sqls"
for file in ${folder}/*
do
  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -D${MYSQL_DB} < ${file}
done

<<xxxx

mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}

mysql -h127.0.0.1 -P3306 -uroot -pxxx -Dtest
#参数
-h:host主机
-P:port端口
-u:user用户名
-p:password密码
-D:database数据库

xxxx
