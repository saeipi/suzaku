### 创建镜像
```
docker commit -a "saeipi.com" -m "centos:7.9.2009 net-tools wget vim" 82a4f1d7187c clinux:7.9.2009

参数说明；
-a :提交的镜像作者；
-c :使用Dockerfile指令来创建镜像；
-m :提交时的说明文字；
-p :在commit时，将容器暂停
```
### 创建tag
```
docker tag clinux:7.9.2009 saeipi/clinux:7.9.2009
```
### 镜像推送
```
docker push saeipi/clinux:7.9.2009
```