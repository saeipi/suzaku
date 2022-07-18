nohup redis-server /usr/local/etc/redis/redis.conf &

sleep 5

redis-cli --cluster create \
    172.18.0.11:7001 172.18.0.12:7002 172.18.0.13:7003 172.18.0.14:7004 \
    172.18.0.15:7005 172.18.0.16:7006 172.18.0.17:7007 172.18.0.18:7008 \
    --cluster-yes --cluster-replicas 1

tail -f /dev/null