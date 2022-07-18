### 本次开机生效：
```
# IPv4 的转发
$ sudo sysctl -w net.inet.ip.forwarding=1
net.inet.ip.forwarding: 0 -> 1

# IPv6 的转发
$ sudo sysctl -w net.inet6.ip6.forwarding=1
net.inet6.ip6.forwarding: 0 -> 1
```

### 开机启动配置，需以 root 身份添加或修改 /etc/sysctl.conf 文件，加入以下两行：
```
net.inet.ip.forwarding=1
net.inet6.ip6.forwarding=1
```

### 查看当前端口转发功能状态：
```
$ sudo sysctl -a | grep forward
net.inet.ip.forwarding: 0
net.inet6.ip6.forwarding: 0
```

### 开启端口转发之后，即可配置端口转发规则。你可以跟着手册来：
```
$ man pfctl
$ man pf.conf
```

### 或者跟着下文手动新建文件。如/etc/pf.anchors/http文件内容如下：
```
rdr pass on lo0 inet proto tcp from any to any port 80 -> 127.0.0.1 port 8080
rdr pass on lo0 inet proto tcp from any to any port 443 -> 127.0.0.1 port 4443
rdr pass on en0 inet proto tcp from any to any port 80 -> 127.0.0.1 port 8080
rdr pass on en0 inet proto tcp from any to any port 443 -> 127.0.0.1 port 4443
```
```
echo "rdr pass proto tcp from any to any port {7001,7002,7003,7004,7005,7006,7007,7008} -> 127.0.0.1" | sudo pfctl -ef -

echo "rdr pass proto tcp from any to {172.18.0.0/24} -> 127.0.0.1" | sudo pfctl -ef -

echo "rdr pass on lo0 inet proto tcp from any to 172.18.0.11 port 7001 -> 10.0.115.108 port 7001" | sudo pfctl -ef -

sudo pfctl -sn
```

### 检查其正确性：
```
$ sudo pfctl -vnf /etc/pf.anchors/http
```

修改PF的主配置文件/etc/pf.conf开启我们添加的锚点http。

pf.conf对指令的顺序有严格要求，相同的指令需要放在一起，否则会报错 Rules must be in order: options, normalization, queueing, translation, filtering.
```
# 在 rdr-anchor "com.apple/*" 下添加
rdr-anchor "http-forwarding"

# 在 load anchor "com.apple" from "/etc/pf.anchors/com.apple" 下添加
load anchor "http-forwarding" from "/etc/pf.anchors/http"
```
### 最后导入并允许运行：
```
$ sudo pfctl -ef /etc/pf.conf
```

### 使用-e命令启用pf服务。使用-E命令强制重启PF服务：
```
$ sudo pfctl -E
```

### 使用-d命令关闭PF：
```
$ sudo pfctl -d
```

从Mavericks起PF服务不再默认开机自启。如需开机启动PF服务，请往下看。

新版Mac OS 10.11 EI Captian加入了系统完整性保护机制，需重启到安全模式执行下述命令关闭文件系统保护。

```
csrutil enable --without fs
```