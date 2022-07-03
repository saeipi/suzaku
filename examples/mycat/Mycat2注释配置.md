https://www.w3cschool.cn/mycat2/mycat2-xok13kry.html

# Mycat2 注释配置

注释使用SQL注释方式表达,可以用于动态更新Mycat配置并且把配置持久化,它的设计目标是为了动态的更新mycat的配置。但是由于配置的属性复杂,它不会自动的更改真实数据库的schema。

通过注解配置不会自动创建物理库物理表(与直接使用自动化建表语句不同,它会自动建物理库物理表),所以要保证物理库物理表的在真实数据库上是与配置对应的。一般来说,原型库(prototype)上必须存在与逻辑库逻辑表完全一致得物理库物理表,以便mycat读取表和字段信息。

如果搞不懂配置,可以尝试使用自动化建表语句创建测试的物理库物理表,它会自动生成配置文件,然后通过查看本地的配置文件,观察它的属性,就可以知道什么回事。因为自动化建表语句过于简单,可能不适合公司的业务,此时需要更改配置文件的属性来调整。这种自己更改调整的属性值不属于mycat的开发测试范畴之内,也不能受mycat为自动化建表的测试保证。

### 重置配置

`/*+ mycat:resetConfig{} */`

### 创建用户

`/*+ mycat:createUser{
  "username":"user",
  "password":"",
  "ip":"127.0.0.1",
  "transactionType":"xa"
} */`

### 删除用户

`/*+ mycat:dropUser{
  "username":"user"} */`

### 显示用户

`/*+ mycat:showUsers */`

### 修改表序列号为雪花算法

`/*+ mycat:setSequence{"name":"db1_travelrecord","time":true} */;`

### 创建数据源

`/*+ mycat:createDataSource{
  "dbType":"mysql",
  "idleTimeout":60000,
  "initSqls":[],
  "initSqlsGetConnection":true,
  "instanceType":"READ_WRITE",
  "maxCon":1000,
  "maxConnectTimeout":3000,
  "maxRetryCount":5,
  "minCon":1,
  "name":"dr0",
  "password":"123456",
  "type":"JDBC",
  "url":"jdbc:mysql://127.0.0.1:3306?useUnicode=true&serverTimezone=UTC&characterEncoding=UTF-8",
  "user":"root",
  "weight":0
} */;`

### 删除数据源

`/*+ mycat:dropDataSource{
  "dbType":"mysql",
  "idleTimeout":60000,
  "initSqls":[],
  "initSqlsGetConnection":true,
  "instanceType":"READ_WRITE",
  "maxCon":1000,
  "maxConnectTimeout":3000,
  "maxRetryCount":5,
  "minCon":1,
  "name":"newDs",
  "type":"JDBC",
  "weight":0
} */;`

### 显示数据源

`/*+ mycat:showDataSources{} */`

### 创建集群

`/*! mycat:createCluster{
  "clusterType":"MASTER_SLAVE",
  "heartbeat":{
    "heartbeatTimeout":1000,
    "maxRetry":3,
    "minSwitchTimeInterval":300,
    "slaveThreshold":0
  },
  "masters":[
    "dw0" //主节点
  ],
  "maxCon":2000,
  "name":"c0",
  "readBalanceType":"BALANCE_ALL",
  "replicas":[
    "dr0" //从节点
  ],
  "switchType":"SWITCH"
} */;`

### 删除集群

`/*! mycat:dropCluster{
  "name":"testAddCluster"
} */;`

### 显示集群

`/*+ mycat:showClusters{} */`

### 创建Schema

确保原型库上,存在`test_add_Schema`物理库,以下注解才能正常运行.

`/*+ mycat:createSchema{
  "customTables":{},
  "globalTables":{},
  "normalTables":{},
  "schemaName":"test_add_Schema",
  "shardingTables":{},
  "targetName":"prototype"
} */;` 

建表注释可以[参考库(schema)配置](https://www.w3cschool.cn/mycat2/mycat2-adpm3ks3.html) 以下注解相当于把配置推送到mycat中进行更新

### 创建单表(用于读写分离,映射物理表)

`/*+ mycat:createTable{
  "normalTable":{
    "createTableSQL":"create table normal(id int)",
    "dataNode":{
      "schemaName":"testSchema", //物理库
      "tableName":"normal", //物理表
      "targetName":"prototype" //目标
    }
  },
  "schemaName":"testSchema",//逻辑库
  "tableName":"normal" //逻辑表
} */;`

### 1.18后

`/*+ mycat:createTable{
  "normalTable":{
    "createTableSQL":"create table normal(id int)",
    "locality":{
      "schemaName":"testSchema", //物理库
      "tableName":"normal", //物理表
      "targetName":"prototype" //目标
    }
  },
  "schemaName":"testSchema",//逻辑库
  "tableName":"normal" //逻辑表
} */;`

当目标是集群的时候,自动进行读写分离,根据集群配置把查询`sql`根据事务状态发送到从或主数据源,如果目标是数据源,就直接发送`sql`到这个数据源.在`Mycat2`中,是否使用`Mycat`的集群配置应该是整体的架构选项,只能选其一.当全体目标都是数据源,要么全体目标都是集群。前者一般在数据库集群前再部署一个`SLB`服务,Mycat访问这个`SLB`服务,这个`SLB`服务实现**读写分离**和**高可用**。后者则是Mycat直接访问数据库,`Mycat`负责**读写分离**和**集群高可用**.当配置中出现集群和数据源的情况,尽量配置成他们的表的存储节点在一个物理库的实例中没有交集,这样可以**避免**因为多使用连接导致事务一致性和隔离级别破坏产生的问题.

### 创建全局表

`/*+ mycat:createTable{
  "globalTable":{
    "createTableSQL":"create table global(id int)",
    "dataNodes":[
      {
        "targetName":"prototype"
      }
    ]
  },
  "schemaName":"testSchema",
  "tableName":"global"
} */;`

### 1.18后

`/*+ mycat:createTable{
  "globalTable":{
    "createTableSQL":"create table global(id int)",
    "broadcast":[
      {
        "targetName":"prototype"
      }
    ]
  },
  "schemaName":"testSchema",
  "tableName":"global"
} */;`

### 创建范围分表

`/*+ mycat:createTable{
  "schemaName":"testSchema",
  "shadingTable":{
    "createTableSQL":"create table sharding(id int)",
    "dataNode":{
      "schemaNames":"testSchema",
      "tableNames":"sharding",
      "targetNames":"prototype"
    },
    "function":{
      "clazz":"io.mycat.router.mycat1xfunction.PartitionConstant",
      "properties":{
        "defaultNode":"0",
        "columnName":"id"
      }
    }
  },
  "tableName":"sharding"
} */;`

### 1.18后

`/*+ mycat:createTable{
  "schemaName":"testSchema",
  "shadingTable":{
    "createTableSQL":"create table sharding(id int)",
    "partition":{
      "schemaNames":"testSchema",
      "tableNames":"sharding",
      "targetNames":"prototype"
    },
    "function":{
      "clazz":"io.mycat.router.mycat1xfunction.PartitionConstant",
      "properties":{
        "defaultNode":"0",
        "columnName":"id"
      }
    }
  },
  "tableName":"sharding"
} */;`

### 显示session引用的IO缓冲块计数

`/*+ mycat:showBufferUsage{}*/`

### 显示用户

`/*+ mycat:showUsers{}*/`

### 显示schema

`/*+ mycat:showSchemas{}*/`

### 显示调度器

`/*+ mycat:showSchedules{}*/`

### 显示心跳配置

`/*+ mycat:showHeartbeats{}*/`

### 显示心跳状态

`/*+ mycat:showHeartbeatStatus{}*/`

### 显示实例状态

`/*+ mycat:showInstances{}*/`

### 显示Reactor状态

`/*+ mycat:showReactors{}*/`

### 显示线程池状态

`/*+ mycat:showThreadPools{}*/`

### 显示表

`/*+ mycat:showTables{"schemaName":"mysql"}*/`

### 显示mycat连接

`/*+ mycat:showConnections{}*/`

### 显示存储节点

`/*+ mycat:showDataNodes{//1.18前
  "schemaName":"db1",
  "tableName":"normal"
} */;




/*+ mycat:showTopology{//1.18后
  "schemaName":"db1",
  "tableName":"normal"
} */;`