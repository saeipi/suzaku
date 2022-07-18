#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-teamgram2022}
MYSQL_DB="teamgram"
SCRIPT_PATH=$(cd $(dirname $0);pwd)
SQL_FILE_INIT_DB="/init/teamgram2.sql"
SQL_FILE_MIGRATES="/migrates"

CREATE_DATABASE="create database ${MYSQL_DB} default character set utf8mb4 collate utf8mb4_unicode_ci;"
mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "$CREATE_DATABASE"

mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} < "$SCRIPT_PATH$SQL_FILE_INIT_DB"

folder=$SCRIPT_PATH$SQL_FILE_MIGRATES
for file in ${folder}/*
do
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} < ${file}
done

CHANGE_AUTHENTICATION="update mysql.user set authentication_string='' where user='root';"
CHANGE_PASSWORD="ALTER user 'root'@'localhost' IDENTIFIED BY '';"
FLUSH_PRIVILEGES_SQL="FLUSH PRIVILEGES;"
mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "$CHANGE_AUTHENTICATION $CHANGE_PASSWORD $FLUSH_PRIVILEGES_SQL"