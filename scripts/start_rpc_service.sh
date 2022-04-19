#!/usr/bin/env bash

source ./function.sh

COLOR_SUFFIX="\033[0m"
BLACK_PREFIX="\033[30m"
RED_PREFIX="\033[31m"
GREEN_PREFIX="\033[32m"
YELLOW_PREFIX="\033[33m"
BLUE_PREFIX="\033[34m"
PURPLE_PREFIX="\033[35m"
SKY_BLUE_PREFIX="\033[36m"

msg_name="open_im_msg"

config_path="../configs/config.yaml"

#service filename
service_filename=(
  #api
  api

  #rpc
  auth
  chat
  friend
  user
  ${msg_name}
)

#service config port name
service_port_name=(
  #api port name
  ApiPort
  #rpc port name
  AuthPort
  ChatPort
  FriendPort
  UserPort
)

for ((i = 0; i < ${#service_filename[*]}; i++)); do
  #Check whether the service exists
  service_name="ps -aux |grep -w ${service_filename[$i]} |grep -v grep"
  count="${service_name}| wc -l"

  if [ $(eval ${count}) -gt 0 ]; then
    pid="${service_name}| awk '{print \$2}'"
    echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} service has been started,pid:$(eval $pid)$COLOR_SUFFIX"
    echo -e "${SKY_BLUE_PREFIX}Killing the service ${service_filename[$i]} pid:$(eval $pid)${COLOR_SUFFIX}"
    #kill the service that existed
    kill -9 $(eval $pid)
    sleep 0.5
  fi
  cd ../bin && echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} service is starting${COLOR_SUFFIX}"
  #Get the rpc port in the configuration file
  portList=$(cat $config_path | grep ${service_port_name[$i]} | awk -F '[:]' '{print $NF}')
  list_to_string ${portList}
  #Start related rpc services based on the number of ports
  for j in ${ports_array}; do
    echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} Service is starting,port number:$j $COLOR_SUFFIX"
    #Start the service in the background
    #    ./${service_filename[$i]} -port $j &
    nohup ./${service_filename[$i]} -port $j >>../logs/openIM.log 2>&1 &
    sleep 1
    pid="netstat -ntlp|grep $j |awk '{printf \$7}'|cut -d/ -f1"
    echo -e "${RED_PREFIX}${service_filename[$i]} Service is started,port number:$j pid:$(eval $pid)$COLOR_SUFFIX"
  done
done
