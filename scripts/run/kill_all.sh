#!/usr/bin/env bash

source ./../cfg/style_info.cfg
source ./../cfg/services.cfg

uNames=`uname -s`
osName=${uNames: 0: 4}
system=""
if [ "$osName" = "Darw" ]
then
	system="mac"
elif [ "$osName" = "Linu" ]
then
	system="linux"
elif [ "$osName" = "MING" ]
then
	system="windows"
else
	echo "unknown os"
fi

for i in ${service_names[*]}; do
  #Check whether the service exists
  name="ps -aux |grep -w $i |grep -v grep"
  count="${name}| wc -l"
  if [ $(eval ${count}) -gt 0 ]; then
    pid="${name}| awk '{print \$2}'"
    echo -e "${SKY_BLUE_PREFIX}Killing service:$i pid:$(eval $pid)${COLOR_SUFFIX}"
    #kill the service that existed
    kill -9 $(eval $pid)
    echo -e "${SKY_BLUE_PREFIX}service:$i was killed ${COLOR_SUFFIX}"
  fi
done
