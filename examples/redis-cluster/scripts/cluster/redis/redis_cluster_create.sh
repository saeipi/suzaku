for i in `seq 7000 7007`;
do
	exec redis-server ./configs/cluster/redis/${i}/redis.conf &
done
redis-cli --cluster create\
127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003\
127.0.0.1:7004 127.0.0.1:7005 127.0.0.1:7006 127.0.0.1:7007\
--cluster-replicas 1