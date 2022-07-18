# docker build -t teamgram:0.86.2 .

FROM golang:1.18 as build

MAINTAINER saeipi "saeipi@163.com"

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /teamgram
WORKDIR /teamgram
COPY . .
RUN chmod -R 777 *.sh
RUN /bin/sh -c ./build.sh


FROM centos:7.9.2009

RUN yum -y install vim && yum -y install net-tools

RUN mkdir -p /teamgram
COPY --from=build /teamgram/teamgramd /teamgram/teamgramd

WORKDIR /teamgram/teamgramd/bin
RUN chmod -R 777 *.sh
ENTRYPOINT ["./runall2.sh"]