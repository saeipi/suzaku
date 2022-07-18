nohup redis-server /usr/local/etc/redis/redis.conf &

sleep 5

redis-cli --cluster create \
    172.18.0.11:6379 172.18.0.12:6379 172.18.0.13:6379 172.18.0.14:6379 \
    172.18.0.15:6379 172.18.0.16:6379 172.18.0.17:6379 172.18.0.18:6379 \
    --cluster-yes --cluster-replicas 1

tail -f /dev/null