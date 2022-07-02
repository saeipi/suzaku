数据库中间件
```
https://github.com/XiaoMi/Gaea
https://github.com/MyCATApache/Mycat2
https://github.com/apache/shardingsphere
```

当Mysql容器首次启动时，会在 /docker-entrypoint-initdb.d目录下扫描 .sh，.sql，.sql.gz类型的文件。如果这些类型的文件存在，将执行它们来初始化一个数据库。这些文件会按照字母的顺序执行。默认情况下它们会初始化在启动容器时声明的 MYSQL_DATABASE变量定义的数据库中,例如下面的命令会初始化一个suzaku数据库：

```
docker run --name some-mysql -e MYSQL_DATABASE=suzaku -d mysql:8.0.29
```

如果你的启动命令没有指定数据库那么就必须在数据库DDL脚本中声明并指定使用该数据库。否则就会实现下面的异常：
```
ERROR 1046 ( 3D000) at line 7: No database selected
```

需要提前同步设置脚本的权限，否则容器会出现错误。/bin/bash: bad interpreter: Permission denied
```
# chmod -R 777 scripts
```