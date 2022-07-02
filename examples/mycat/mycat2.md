安装mycat2
### 安装java
yum install java

mysql自行安装 注：官方推荐mysql版本8.0.14以上 

### 新建目录soft
mkdir /soft
 
### 进入/soft
cd /soft
 
### 下载zip包
wget http://dl.mycat.org.cn/2.0/install-template/mycat2-install-template-1.20.zip
 
### 解压
unzip mycat2-install-template-1.20.zip

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
 
### 编辑mysql数据库配置，修改当前mysql配置信息
vim conf/datasources/prototypeDs.datasource.json
```
{
        "dbType":"mysql",
        "idleTimeout":60000,
        "initSqls":[],
        "initSqlsGetConnection":true,
        "instanceType":"READ_WRITE",
        "maxCon":1000,
        "maxConnectTimeout":3000,
        "maxRetryCount":5,
        "minCon":1,
        "name":"prototypeDs",
        "password":"123456",
        "type":"JDBC",
        "url":"jdbc:mysql://10.10.10.10:3306/mysql?useUnicode=true&serverTimezone=Asia/Shanghai&characterEncoding=UTF-8",
        "user":"root",
        "weight":0
}
```
### 保存
:wq
 
### 进入到bin目录
cd bin/
 
### 给mycat文件执行权限
chmod +x mycat
 
### 启动mycat
./mycat start
　　

提示 ./wrapper-linux-x86-64 (Found but not executable.)、./wrapper-linux-x86-32 (Found but not executable.) 这两个文件不可执行

### 给权限
chmod 755 ./wrapper-linux-x86-64 ./wrapper-linux-x86-32
 
### 再次启动
./mycat start
　　
### 登陆成功
mysql -uroot -p123456 -P8066 -hlocalhost


### 配置集群
#### 创建逻辑库
```
mysql> create database szk;
```

#### 指定数据源
```
vim /usr/local/mycat2/conf/schemas/skz.schema.json
{
  "schemaName": "szk",
  "targetName": "prototype"
}
```

#### 增加数据源
```
/*+ mycat:createDataSource{ "name":"rwSepw2", "url":"jdbc:mysql://10.0.115.108:13307/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepr2","url":"jdbc:mysql://10.0.115.108:13309/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepw1", "url":"jdbc:mysql://10.0.115.108:13306/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;


/*+ mycat:createDataSource{ "name":"rwSepr1","url":"jdbc:mysql://10.0.115.108:13308/suzaku?useSSL=false&characterEncodin g=UTF-8&useJDBCCompliantTimezoneShift=true", "user":"root", "password":"123456" } */;

```

#### 修改集群配置
```
vim /usr/local/mycat2/conf/clusters/prototype.cluster.json

{
        "clusterType":"MASTER_SLAVE",
        "heartbeat":{
                "heartbeatTimeout":1000,
                "maxRetry":3,
                "minSwitchTimeInterval":300,
                "slaveThreshold":0
        },
        "masters":[
                "rwSepw1","rwSepw2"
        ],
        "replicas":["rwSepw2","rwSepr1","rwSepr2" ],
        "timer":{ "initialDelay": 30, "period":5,"timeUnit":"SECONDS" },
        "maxCon":200,
        "name":"prototype",
        "readBalanceType":"BALANCE_ALL",
        "switchType":"SWITCH"
}
```

```
vim suzaku.schema.json

{
  "schemaName": "suzaku",
  "targetName": "prototype"
}
```
#### 重启mycat
```
[root@26a7b7740b39 bin]# cd /usr/local/mycat2/bin
[root@26a7b7740b39 bin]# ./mycat restart
```