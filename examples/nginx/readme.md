### 将容器内的配置文件拷贝到当前目录: docker container cp nginx:/etc/nginx .
### 重命名: mv nginx conf
### vim /etc/nginx/conf.d/default.conf
```
upstream suzaku {
    server 10.0.115.108:9166;
    server 10.0.115.108:9167;
    server 10.0.115.108:9168;
}

server {
    listen       80;
    server_name  suzaku.com;

    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        proxy_pass http://suzaku;
    }
}
```

### Golang redis分布式锁
```
Golang redis分布式锁
https://redis.io/docs/reference/patterns/distributed-locks/
https://github.com/go-redsync/redsync
```