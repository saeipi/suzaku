#!/usr/bin/env bash

cp -Rp ../configs ../build/bin

PWD=`pwd`
SUZAKUAPP=${PWD}"/../cmd"
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
