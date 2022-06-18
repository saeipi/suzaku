#!/usr/bin/env bash

rm -f -r ../build/bin
rm -f -r ../build/configs

mkdir -p ../build/bin

cp -Rp run/*.* ../build/bin
cp -Rp ../configs ../build/configs

PWD=`pwd`
SUZAKUAPP=${PWD}"/../cmd"
RPCSERVICE=${SUZAKUAPP}"/rpc"
INSTALL=${PWD}"/../build"

echo "build api ..."
cd ${SUZAKUAPP}/api
go build -o ${INSTALL}/bin/api

echo "build msg_gateway ..."
cd ${SUZAKUAPP}/msg_gateway
go build -o ${INSTALL}/bin/msg_gateway

echo "build msg_transfer ..."
cd ${SUZAKUAPP}/msg_transfer
go build -o ${INSTALL}/bin/msg_transfer

echo "build push ..."
cd ${SUZAKUAPP}/push
go build -o ${INSTALL}/bin/push

echo "build rpc/auth ..."
cd ${RPCSERVICE}/auth
go build -o ${INSTALL}/bin/rpc_auth

echo "build rpc/cache ..."
cd ${RPCSERVICE}/cache
go build -o ${INSTALL}/bin/rpc_cache

echo "build rpc/chat ..."
cd ${RPCSERVICE}/chat
go build -o ${INSTALL}/bin/rpc_chat

echo "build rpc/friend ..."
cd ${RPCSERVICE}/friend
go build -o ${INSTALL}/bin/rpc_friend

echo "build rpc/group ..."
cd ${RPCSERVICE}/group
go build -o ${INSTALL}/bin/rpc_group

echo "build rpc/user ..."
cd ${RPCSERVICE}/user
go build -o ${INSTALL}/bin/rpc_user


