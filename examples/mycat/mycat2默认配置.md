```
[root@localhost conf]# cd users/
[root@localhost users]# ls
root.user.json
[root@localhost users]# cat root.user.json 
{
	"dialect":"mysql",
	"ip":null,
	"password":"123456",
	"transactionType":"proxy",
	"username":"root"
}


[root@localhost conf]# cd clusters/
[root@localhost clusters]# ls
prototype.cluster.json
[root@localhost clusters]# cat prototype.cluster.json
{
	"clusterType":"MASTER_SLAVE",
	"heartbeat":{
		"heartbeatTimeout":1000,
		"maxRetry":3,
		"minSwitchTimeInterval":300,
		"slaveThreshold":0
	},
	"masters":[
		"prototypeDs"
	],
	"maxCon":200,
	"name":"prototype",
	"readBalanceType":"BALANCE_ALL",
	"switchType":"SWITCH"
}

[root@localhost conf]# cd datasources/
[root@localhost datasources]# ls
prototypeDs.datasource.json
[root@localhost datasources]# cat prototypeDs.datasource.json
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
	"url":"jdbc:mysql://localhost:3306/mysql?useUnicode=true&serverTimezone=Asia/Shanghai&characterEncoding=UTF-8",
	"user":"root",
	"weight":0
}[root@localhost datasources]# 


[root@localhost conf]# ls
clusters     dbseq.sql    mycat.lock  sequences    simplelogger.properties  sqlcaches   users        wrapper.conf
datasources  logback.xml  schemas     server.json  sql                      state.json  version.txt
[root@localhost conf]# cat state.json
{
	"replica":{
		"prototype":["prototypeDs"]
	}
}


[root@localhost conf]# ls
clusters     dbseq.sql    mycat.lock  sequences    simplelogger.properties  sqlcaches   users        wrapper.conf
datasources  logback.xml  schemas     server.json  sql                      state.json  version.txt
[root@localhost conf]# cat server.json
{
  "loadBalance":{
    "defaultLoadBalance":"BalanceRandom",
    "loadBalances":[]
  },
  "mode":"local",
  "properties":{},
  "server":{
    "bufferPool":{

    },
    "idleTimer":{
      "initialDelay":3,
      "period":60000,
      "timeUnit":"SECONDS"
    },
    "ip":"0.0.0.0",
    "mycatId":1,
    "port":8066,
    "reactorNumber":8,
    "tempDirectory":null,
    "timeWorkerPool":{
      "corePoolSize":0,
      "keepAliveTime":1,
      "maxPendingLimit":65535,
      "maxPoolSize":2,
      "taskTimeout":5,
      "timeUnit":"MINUTES"
    },
    "workerPool":{
      "corePoolSize":1,
      "keepAliveTime":1,
      "maxPendingLimit":65535,
      "maxPoolSize":1024,
      "taskTimeout":5,
      "timeUnit":"MINUTES"
    }
  }
}

```