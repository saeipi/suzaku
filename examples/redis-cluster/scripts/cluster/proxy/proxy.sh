#!/usr/bin/env bash

iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7001 -j DNAT --to-destination 172.18.0.11:7001
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7002 -j DNAT --to-destination 172.18.0.12:7002
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7003 -j DNAT --to-destination 172.18.0.13:7003
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7004 -j DNAT --to-destination 172.18.0.14:7004
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7005 -j DNAT --to-destination 172.18.0.15:7005
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7006 -j DNAT --to-destination 172.18.0.16:7006
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7007 -j DNAT --to-destination 172.18.0.17:7007
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 7008 -j DNAT --to-destination 172.18.0.18:7008

tail -f /dev/null