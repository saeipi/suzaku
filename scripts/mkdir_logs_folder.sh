folder="../cmd"
logs_dir="../build/logs/"
for sub in ${folder}/*
do
  if [ "${sub##*.}"x = "md"x ];then
    continue
  fi
  log=${sub//"..\/cmd\/"/""}
  mkdir -p $logs_dir$log
done
