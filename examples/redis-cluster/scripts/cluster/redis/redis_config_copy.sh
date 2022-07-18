for i in `seq 7001 7008`;
do
  mkdir -p ./../../../configs/cluster/redis/${i}
  cp -Rp ./../../../configs/cluster/redis.conf  ./../../../configs/cluster/redis/${i}
  # 换行
  echo -e >> ./../../../configs/cluster/redis/${i}/redis.conf
  echo -e "port "${i} >> ./../../../configs/cluster/redis/${i}/redis.conf
done

chmod -R 7777 ./../../../configs