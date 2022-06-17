#!/usr/bin/env bash

echo "run api ..."
nohup ./api &
sleep 1

echo "run msg_gateway ..."
nohup ./msg_gateway &
sleep 1

echo "run msg_transfer ..."
nohup ./msg_transfer &
sleep 1

echo "run push ..."
nohup ./push &
sleep 1

echo "run rpc_auth ..."
nohup ./rpc_auth &
sleep 1

echo "run rpc_cache ..."
nohup ./rpc_cache &
sleep 1

echo "run rpc_chat ..."
nohup ./rpc_chat &
sleep 1

echo "run rpc_friend ..."
nohup ./rpc_friend &
sleep 1

echo "run rpc_group ..."
nohup ./rpc_group &
sleep 1

echo "run rpc_user ..."
nohup ./rpc_user &
sleep 1