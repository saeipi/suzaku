for i in `seq 7001 7007`;
do
  mkdir -p ./../../../configs/cluster/redis/${i}
  cp -Rp ./../../../configs/cluster/redis/7000/*  ./../../../configs/cluster/redis/${i}
done