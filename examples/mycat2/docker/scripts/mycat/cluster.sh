#!/bin/bash

bash /usr/local/mycat/bin/mycat restart
sleep 10

MYSQL_HOST=${MYSQL_HOST:-127.0.0.1}
MYSQL_USERNAME=${MYSQL_USERNAME:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-123456}
MYSQL_PORT=${MYSQL_PORT:-8066}

CREATE_DATASOURCE="
/*+ mycat:createDataSource{ \"name\":\"rwSepw2\", \"url\":\"jdbc:mysql://10.0.115.108:13307/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true\", \"user\":\"root\", \"password\":\"123456\" } */;
/*+ mycat:createDataSource{ \"name\":\"rwSepr2\",\"url\":\"jdbc:mysql://10.0.115.108:13309/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true\", \"user\":\"root\", \"password\":\"123456\" } */;
/*+ mycat:createDataSource{ \"name\":\"rwSepw1\", \"url\":\"jdbc:mysql://10.0.115.108:13306/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true\", \"user\":\"root\", \"password\":\"123456\" } */;
/*+ mycat:createDataSource{ \"name\":\"rwSepr1\",\"url\":\"jdbc:mysql://10.0.115.108:13308/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true\", \"user\":\"root\", \"password\":\"123456\" } */;
/*+ mycat:showDataSources{} */;"

CREATE_CLUSTER="/*! mycat:createCluster{\"name\":\"c0\",\"masters\":[\"rwSepw1\"],\"replicas\":[\"rwSepr1\"]} */;
                /*! mycat:createCluster{\"name\":\"c1\",\"masters\":[\"rwSepw2\"],\"replicas\":[\"rwSepr2\"]} */;
                /*+ mycat:showClusters{} */;"

CREATE_DATABASE="create database suzaku;"

CREATE_SCHEMA="/*+ mycat:createSchema{
                   \"schemaName\": \"suzaku\",
                   \"targetName\": \"suzaku\",
                   \"normalTables\": {}
               } */;
               /*+ mycat:showSchemas{} */;
               /*+ mycat:setSequence{\"name\":\"suzaku\",\"time\":true} */;
               /*+ mycat:setSequence
               {\"name\":\"suzaku\",\"clazz\":\"io.mycat.plug.sequence.SequenceMySQLGenerator\"} */;"

mysql -h"$MYSQL_HOST" -P"$MYSQL_PORT" -u"$MYSQL_USERNAME" -p"$MYSQL_PASSWORD" -e "$CREATE_DATASOURCE $CREATE_CLUSTER $CREATE_DATABASE $CREATE_SCHEMA"

sleep 5
bash /usr/local/mycat/bin/mycat restart