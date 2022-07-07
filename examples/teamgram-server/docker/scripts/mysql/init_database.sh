#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-teamgram2022}
MYSQL_DB="teamgram"
SCRIPT_PATH=$(cd $(dirname $0);pwd)
SQL_FILE="/sqls"

folder=$SCRIPT_PATH$SQL_FILE
for file in ${folder}/*
do
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} < ${file}
done

CHANGE_AUTHENTICATION="update mysql.user set authentication_string='' where user='root';"
CHANGE_PASSWORD="ALTER user 'root'@'localhost' IDENTIFIED BY '';"
FLUSH_PRIVILEGES_SQL="FLUSH PRIVILEGES;"
mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "$CHANGE_AUTHENTICATION $CHANGE_PASSWORD $FLUSH_PRIVILEGES_SQL"