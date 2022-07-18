### linux IP地址重定向
```
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7001 -j DNAT --to-destination 127.0.0.1:7001
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7002 -j DNAT --to-destination 127.0.0.1:7002
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7003 -j DNAT --to-destination 127.0.0.1:7003
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7004 -j DNAT --to-destination 127.0.0.1:7004
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7005 -j DNAT --to-destination 127.0.0.1:7005
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7006 -j DNAT --to-destination 127.0.0.1:7006
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7007 -j DNAT --to-destination 127.0.0.1:7007
iptables -t nat -I PREROUTING -d 172.18.0.11 -p tcp --destination-port 7008 -j DNAT --to-destination 127.0.0.1:7008
```