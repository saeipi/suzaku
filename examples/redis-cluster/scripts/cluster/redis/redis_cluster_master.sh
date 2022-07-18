nohup redis-server /usr/local/etc/redis/redis.conf &

sleep 5

redis-cli --cluster create \
    redis-node-01:7001 redis-node-02:7002 redis-node-03:7003 redis-node-04:7004 \
    redis-node-05:7005 redis-node-06:7006 redis-node-07:7007 redis-node-08:7008 \
    --cluster-yes --cluster-replicas 1

tail -f /dev/null