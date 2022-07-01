#!/bin/bash

MYSQL_USERNAME=${MYSQL_USERNAME:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-123456}
MYSQL_PORT=${MYSQL_PORT:-3306}

SLAVE_SYNC_USER="${SLAVE_SYNC_USER:-sync_admin}"
SLAVE_SYNC_PASSWORD="${SLAVE_SYNC_PASSWORD:-123456}"

MYSQL_HOST_MASTER_01=${MYSQL_MASTER_01:-10.10.10.10}
MYSQL_HOST_MASTER_02=${MYSQL_MASTER_01:-10.10.10.100}

sleep 10
# 连接master数据库，查询二进制数据，并解析出logfile和pos，这里同步用户要开启 REPLICATION CLIENT权限，才能使用SHOW MASTER STATUS;
RESULT=`mysql -h"$MYSQL_HOST_MASTER_02" -P"$MYSQL_PORT" -u"$SLAVE_SYNC_USER" -p"$SLAVE_SYNC_PASSWORD" -e "SHOW MASTER STATUS;" | grep -v grep |tail -n +2| awk '{print $1,$2}'`
# 解析出logfile
LOG_FILE_NAME=`echo $RESULT | grep -v grep | awk '{print $1}'`
# 解析出pos
LOG_FILE_POS=`echo $RESULT | grep -v grep | awk '{print $2}'`

# 设置连接master的同步相关信息
SYNC_SQL="change master to master_host='$MYSQL_HOST_MASTER_02',master_user='$SLAVE_SYNC_USER',master_password='$SLAVE_SYNC_PASSWORD',master_log_file='$LOG_FILE_NAME',master_log_pos=$LOG_FILE_POS;"
# 开启同步
START_SYNC_SQL="start slave;"
# 查看同步状态
STATUS_SQL="show slave status\G;"

mysql -h"$MYSQL_HOST_MASTER_01" -P"$MYSQL_PORT" -u"$MYSQL_USERNAME" -p"$MYSQL_PASSWORD" -e "$SYNC_SQL $START_SYNC_SQL $STATUS_SQL"


sleep 10
# 连接master数据库，查询二进制数据，并解析出logfile和pos，这里同步用户要开启 REPLICATION CLIENT权限，才能使用SHOW MASTER STATUS;
RESULT=`mysql -h"$MYSQL_HOST_MASTER_01" -P"$MYSQL_PORT" -u"$SLAVE_SYNC_USER" -p"$SLAVE_SYNC_PASSWORD" -e "SHOW MASTER STATUS;" | grep -v grep |tail -n +2| awk '{print $1,$2}'`
# 解析出logfile
LOG_FILE_NAME=`echo $RESULT | grep -v grep | awk '{print $1}'`
# 解析出pos
LOG_FILE_POS=`echo $RESULT | grep -v grep | awk '{print $2}'`

# 设置连接master的同步相关信息
SYNC_SQL="change master to master_host='$MYSQL_HOST_MASTER_01',master_user='$SLAVE_SYNC_USER',master_password='$SLAVE_SYNC_PASSWORD',master_log_file='$LOG_FILE_NAME',master_log_pos=$LOG_FILE_POS;"
# 开启同步
START_SYNC_SQL="start slave;"
# 查看同步状态
STATUS_SQL="show slave status\G;"

mysql -h"$MYSQL_HOST_MASTER_02" -P"$MYSQL_PORT" -u"$MYSQL_USERNAME" -p"$MYSQL_PASSWORD" -e "$SYNC_SQL $START_SYNC_SQL $STATUS_SQL"