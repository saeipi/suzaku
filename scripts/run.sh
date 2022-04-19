#!/usr/bin/env bash

source ./main_dir.cfg

length=${#all_main}
for (( i=0; i <= $length; i++ )); do
  entry=${all_main[$i]}
  echo "main" $entry
  if [ -z "$entry" ]; then
    continue
  fi
  nohup go run ./+$entry &
done
echo "run success..."