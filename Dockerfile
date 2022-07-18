# docker build -t suzaku:1.0.3 .
# systemctl start docker # 启动
# systemctl daemon-reload # 守护进程重启
# systemctl restart docker / service docker restart # 重启docker服务
# docker service docker stop / docker systemctl stop docker # 关闭
# docker run -it d64194bce4e /bin/bash
# docker rmi -f $(docker images -q -f dangling=true)
# export -p
# docker run -it -p 10000:10000 -p 17778:17778 --network suzaku_szk-network --name szk b1e396e11d21 --privileged=true -v /volumes/suzaku/logs:/suzaku/build/logs
# docker network inspect 896d3748faa3
# docker network disconnect -f suzaku_szk-network minio
# docker stats 状态查看
# 查看端口监听情况, Linux: netstat -lntp, Mac: netstat -AaLlnW
# 删除目录及子目录下的所有.log文件: find . -name "*.log" | xargs rm
#======================== 1 golang ========================#
## 源镜像
FROM golang:1.18 as build

## 设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    AppRunMode=prod

## 作者
MAINTAINER saeipi "saeipi@163.com"
## 在docker的根目录下创立相应的应用目录
RUN mkdir -p /suzaku
## 把宿主机上指定目录下的文件复制到/suzaku目录下
WORKDIR /suzaku
COPY . .
## 编译项目
WORKDIR /suzaku/scripts
RUN chmod +x *.sh
RUN /bin/sh -c ./build.sh

#======================== 2 ubuntu ========================#
FROM ubuntu

RUN rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install apt-transport-https && apt-get install procps\
&&apt-get install net-tools

ENV DEBIAN_FRONTEND=noninteractive
ENV AppRunMode prod

RUN apt-get install -y vim curl tzdata gawk
RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

RUN mkdir -p /suzaku
RUN mkdir -p /var/log/suzaku

COPY --from=build /suzaku/build /suzaku/build
COPY --from=build /suzaku/configs /suzaku/configs

WORKDIR /suzaku/build/bin
RUN chmod +x *.sh
ENTRYPOINT ["./run_all.sh"]

EXPOSE 10000 17778