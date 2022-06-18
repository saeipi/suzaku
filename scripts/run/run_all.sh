#!/usr/bin/env bash

echo "run api ..."
nohup ./api > ../logs/api.log 2>&1 &
sleep 1

echo "run msg_gateway ..."
nohup ./msg_gateway > ../logs/msg_gateway.log 2>&1 &
sleep 1

echo "run msg_transfer ..."
nohup ./msg_transfer > ../logs/msg_transfer.log 2>&1 &
sleep 1

echo "run push ..."
nohup ./push > ../logs/push.log 2>&1 &
sleep 1

echo "run rpc_auth ..."
nohup ./rpc_auth > ../logs/rpc_auth.log 2>&1 &
sleep 1

echo "run rpc_cache ..."
nohup ./rpc_cache > ../logs/rpc_cache.log 2>&1 &
sleep 1

echo "run rpc_chat ..."
nohup ./rpc_chat > ../logs/rpc_chat.log 2>&1 &
sleep 1

echo "run rpc_friend ..."
nohup ./rpc_friend > ../logs/rpc_friend.log 2>&1 &
sleep 1

echo "run rpc_group ..."
nohup ./rpc_group > ../logs/rpc_group.log 2>&1 &
sleep 1

echo "run rpc_user ..."
nohup ./rpc_user > ../logs/rpc_user.log 2>&1 &
sleep 1