# docker build -t suzuku:1.0.1 .
## 源镜像
FROM golang as build

## 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV AppRunMode prod
## 作者
MAINTAINER saeipi "saeipi@163.com"
## 在docker的根目录下创立相应的应用目录
RUN mkdir -p /suzaku
## 设置工作目录
WORKDIR /suzaku
## 把（宿主机上）目录下的文件全部复制到/suzaku目录下
COPY . .
## 编译项目
WORKDIR /suzaku/scripts
RUN chmod +x *.sh
RUN /bin/sh -c ./build.sh

WORKDIR /suzaku/build/bin
RUN chmod +x *.sh
ENTRYPOINT ["./run_all.sh"]
## 裸露端口
EXPOSE 10000 17778

#============================================================

## FROM ubuntu

## RUN rm -rf /var/lib/apt/lists/*
## RUN apt-get update && apt-get install apt-transport-https && apt-get install procps\
## &&apt-get install net-tools

## ENV AppRunMode prod
## RUN apt-get install -y vim curl tzdata gawk

## RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

## COPY --from=build /suzaku/build /suzaku/build

## WORKDIR /suzaku/build
## RUN chmod +x *.sh
## ENTRYPOINT ["./run_all.sh"]

## 裸露端口
## EXPOSE 10000 17778