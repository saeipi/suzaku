#!/usr/bin/env bash

echo "run api ..."
nohup ./api_gw > /var/log/suzaku/api.log 2>&1 &
sleep 1

echo "run msg_gateway ..."
nohup ./msg_gateway > /var/log/suzaku/msg_gateway.log 2>&1 &
sleep 1

echo "run msg_transfer ..."
nohup ./msg_transfer > /var/log/suzaku/msg_transfer.log 2>&1 &
sleep 1

echo "run push ..."
nohup ./push > /var/log/suzaku/push.log 2>&1 &
sleep 1

echo "run rpc_auth ..."
nohup ./rpc_auth > /var/log/suzaku/rpc_auth.log 2>&1 &
sleep 1

echo "run rpc_cache ..."
nohup ./rpc_cache > /var/log/suzaku/rpc_cache.log 2>&1 &
sleep 1

echo "run rpc_chat ..."
nohup ./rpc_chat > /var/log/suzaku/rpc_chat.log 2>&1 &
sleep 1

echo "run rpc_friend ..."
nohup ./rpc_friend > /var/log/suzaku/rpc_friend.log 2>&1 &
sleep 1

echo "run rpc_group ..."
nohup ./rpc_group > /var/log/suzaku/rpc_group.log 2>&1 &
sleep 1

echo "run rpc_user ..."
nohup ./rpc_user > /var/log/suzaku/rpc_user.log 2>&1 &
sleep 1

sleep 10
# fixme prevents the suzaku service exit after execution in the docker container
tail -f /dev/null
# ping 8.8.8.8