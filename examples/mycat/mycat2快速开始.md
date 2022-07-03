### 安装java
yum install java
mysql自行安装 注：官方推荐mysql版本8.0.14以上 

### 下载zip包
mkdir /soft
cd /soft
wget http://dl.mycat.org.cn/2.0/install-template/mycat2-install-template-1.21.zip

### 解压
unzip mycat2-install-template-1.21.zip

### 进入到解压后的mycat
cd mycat/lib
 
### 下载最新的jar包
wget http://dl.mycat.org.cn/2.0/1.21-release/mycat2-1.21-release-jar-with-dependencies.jar

### 返回mycat
cd ../../
 
### 移动mycat到/usr/local/mycat2
mv mycat /usr/local/mycat2
 
### 进入 /usr/local/mycat2
cd /usr/local/mycat2
chmod -R 777 bin

### 进入到bin目录
cd bin/

### 验证mysql是否可以链接成功
```
mysql -uroot -p123456 -h 10.0.115.108 -P 13306
mysql -uroot -p123456 -h 10.0.115.108 -P 13307
mysql -uroot -p123456 -h 10.0.115.108 -P 13308
mysql -uroot -p123456 -h 10.0.115.108 -P 13309
```

### 登陆mycat
```
mysql -uroot -p123456 -P8066 -h127.0.0.1
mysql -uroot -p123456 -P8066 -hlocalhost
```

#### 创建逻辑库
```
mysql> show databases;
mysql> create database suzaku;
mysql> drop database suzaku;

# 先添加数据源 和 集群
/*+ mycat:createSchema{
    // 物理库
    "schemaName": "suzaku",
    // 指向集群，或者数据源
    "targetName": "suzaku",
    // 这里可以配置数据表相关的信息，在物理表已存在或需要启动时自动创建物理表时配置此项
    "normalTables": {}
} */;

/*+ mycat:showSchemas{} */;

# 通过注解设置为雪花算法
/*+ mycat:setSequence{"name":"suzaku","time":true} */;
# 通过注解设置为数据库方式全局序列号
/*+ mycat:setSequence
{"name":"suzaku","clazz":"io.mycat.plug.sequence.SequenceMySQLGenerator"} */;
```

### 配置集群
#### 增加数据源
```
/*+ mycat:createDataSource{ "name":"rwSepw2", "url":"jdbc:mysql://10.0.115.108:13307/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepr2","url":"jdbc:mysql://10.0.115.108:13309/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepw1", "url":"jdbc:mysql://10.0.115.108:13306/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepr1","url":"jdbc:mysql://10.0.115.108:13308/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;

/*+ mycat:showDataSources{} */;

# 删除数据源
/*+ mycat:dropDataSource{
  "dbType":"mysql",
  "idleTimeout":60000,
  "initSqls":[],
  "initSqlsGetConnection":true,
  "instanceType":"READ_WRITE",
  "maxCon":1000,
  "maxConnectTimeout":3000,
  "maxRetryCount":5,
  "minCon":1,
  "name":"rwSepw2",
  "type":"JDBC",
  "weight":0
} */;

```

#### 增加集群
```
# 集群名称以c开头 例如c0
/*! mycat:createCluster{"name":"c0","masters":["rwSepw1","rwSepw2"],"replicas":["rwSepw2","rwSepr1","rwSepr2"]} */;
/*+ mycat:showClusters{} */;

# 删除集群
/*! mycat:dropCluster{
  "name":"c0"
} */;

/*! mycat:createCluster{"name":"c0","masters":["rwSepw1"],"replicas":["rwSepr1"]} */; 
/*! mycat:createCluster{"name":"c1","masters":["rwSepw2"],"replicas":["rwSepr2"]} */;
```

#### 创建逻辑库
```
mysql> show databases;
mysql> create database suzaku;
mysql> drop database suzaku;

# 重置配置
/*+ mycat:resetConfig{} */;
# 创建用户


/*+ mycat:createSchema{
    // 物理库
    "schemaName": "suzaku",
    // 指向集群，或者数据源
    "targetName": "suzaku",
    // 这里可以配置数据表相关的信息，在物理表已存在或需要启动时自动创建物理表时配置此项
    "normalTables": {}
} */;

/*+ mycat:showSchemas{} */;

/*+ mycat:showUsers */;
```

#### 重启mycat
```
[root@26a7b7740b39 bin]# cd /usr/local/mycat2/bin
[root@26a7b7740b39 bin]# ./mycat restart
```

### #在建表语句中加上关键字 BROADCAST（广播，即为全局表）
```
CREATE TABLE `registers` (
  `user_id` varchar(40) NOT NULL DEFAULT '' COMMENT '用户ID 系统生成',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  PRIMARY KEY (`user_id`),
  KEY `idx_userId_password` (`user_id`,`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 BROADCAST;

INSERT INTO `registers` (`user_id`, `password`, `ex`, `created_ts`, `updated_ts`) VALUES ('123456', '123123', 'test', 1656832442, 1656832442);
INSERT INTO `registers` (`user_id`, `password`, `ex`, `created_ts`, `updated_ts`) VALUES ('456789', '123123', 'test', 1656832442, 1656832442);
INSERT INTO `registers` (`user_id`, `password`, `ex`, `created_ts`, `updated_ts`) VALUES ('010101', '123123', 'test', 1656832442, 1656832442);
INSERT INTO `registers` (`user_id`, `password`, `ex`, `created_ts`, `updated_ts`) VALUES ('101010', '123123', 'test', 1656832442, 1656832442);

drop table registers; 
```

### 创建分片表(分库分表)
```
CREATE TABLE friend_groups (
  `id` bigint NOT NULL AUTO_INCREMENT,
   `user_id` int DEFAULT '0',
  `group_name` varchar(60) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `id` ( `id` ) 
) ENGINE = INNODB DEFAULT CHARSET = utf8 dbpartition BY mod_hash (user_id) tbpartition BY mod_hash (user_id) tbpartitions 1 dbpartitions 2;

INSERT INTO `friend_groups` (`group_name`, `user_id`, `desc`) VALUES ('组1',123456, '第1组');

CREATE TABLE orders (
	id BIGINT NOT NULL AUTO_INCREMENT,
	order_type INT,
	customer_id INT,
	amount DECIMAL ( 10, 2 ),
	PRIMARY KEY ( id ),
	KEY `id` ( `id` ) 
) ENGINE = INNODB DEFAULT CHARSET = utf8 dbpartition BY mod_hash ( customer_id ) tbpartition BY mod_hash ( customer_id ) tbpartitions 1 dbpartitions 2;

drop table orders; 

INSERT INTO orders(id,order_type,customer_id,amount) VALUES(1,101,100,100100); 
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(2,101,100,100300); 
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(3,101,101,120000); 
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(4,101,101,103000); 
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(5,102,101,100400); 
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(6,102,102,100020);
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(7,102,103,100020);
INSERT INTO orders(id,order_type,customer_id,amount) VALUES(8,102,104,100020);
SELECT * FROM orders;
```

### Mycat2 注释配置
```
https://www.w3cschool.cn/mycat2/mycat2-xok13kry.html
# 重置配置
/*+ mycat:resetConfig{} */;

# 创建用户
/*+ mycat:createUser{
  "username":"mycat",
  "password":"123456",
  "ip":"127.0.0.1",
  "transactionType":"xa"
} */;

# 删除用户
/*+ mycat:dropUser{
  "username":"user"} */;

# 显示用户
/*+ mycat:showUsers */;

# 修改表序列号为雪花算法
/*+ mycat:setSequence{"name":"suzaku","time":true} */;
```