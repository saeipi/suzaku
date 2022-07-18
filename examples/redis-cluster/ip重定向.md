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

### mac IP地址重定向
```
sudo vim /etc/pf.conf

rdr pass on lo0 inet proto tcp from any to 172.18.0.11 port 7001 -> 127.0.0.1 port 7001
rdr pass on lo0 inet proto tcp from any to 172.18.0.12 port 7002 -> 127.0.0.1 port 7002
rdr pass on lo0 inet proto tcp from any to 172.18.0.13 port 7003 -> 127.0.0.1 port 7003
rdr pass on lo0 inet proto tcp from any to 172.18.0.14 port 7004 -> 127.0.0.1 port 7004
rdr pass on lo0 inet proto tcp from any to 172.18.0.15 port 7005 -> 127.0.0.1 port 7005
rdr pass on lo0 inet proto tcp from any to 172.18.0.16 port 7006 -> 127.0.0.1 port 7006
rdr pass on lo0 inet proto tcp from any to 172.18.0.17 port 7007 -> 127.0.0.1 port 7007
rdr pass on lo0 inet proto tcp from any to 172.18.0.18 port 7008 -> 127.0.0.1 port 7008
```

```
#
# Default PF configuration file.
#
# This file contains the main ruleset, which gets automatically loaded
# at startup.  PF will not be automatically enabled, however.  Instead,
# each component which utilizes PF is responsible for enabling and disabling
# PF via -E and -X as documented in pfctl(8).  That will ensure that PF
# is disabled only when the last enable reference is released.
#
# Care must be taken to ensure that the main ruleset does not get flushed,
# as the nested anchors rely on the anchor point defined here. In addition,
# to the anchors loaded by this file, some system services would dynamically
# insert anchors into the main ruleset. These anchors will be added only when
# the system service is used and would removed on termination of the service.
#
# See pf.conf(5) for syntax.
#

#
# com.apple anchor point
#
rdr pass on lo0 inet proto tcp from any to 172.18.0.11 port 7001 -> 127.0.0.1 port 7001
rdr pass on lo0 inet proto tcp from any to 172.18.0.12 port 7002 -> 127.0.0.1 port 7002
rdr pass on lo0 inet proto tcp from any to 172.18.0.13 port 7003 -> 127.0.0.1 port 7003
rdr pass on lo0 inet proto tcp from any to 172.18.0.14 port 7004 -> 127.0.0.1 port 7004
rdr pass on lo0 inet proto tcp from any to 172.18.0.15 port 7005 -> 127.0.0.1 port 7005
rdr pass on lo0 inet proto tcp from any to 172.18.0.16 port 7006 -> 127.0.0.1 port 7006
rdr pass on lo0 inet proto tcp from any to 172.18.0.17 port 7007 -> 127.0.0.1 port 7007
rdr pass on lo0 inet proto tcp from any to 172.18.0.18 port 7008 -> 127.0.0.1 port 7008
scrub-anchor "com.apple/*"
nat-anchor "com.apple/*"
rdr-anchor "com.apple/*"
rdr-anchor "debookee"
dummynet-anchor "com.apple/*"
anchor "com.apple/*"
anchor "debookee"
load anchor "com.apple" from "/etc/pf.anchors/com.apple"
```

### vim显示行号
```
:set number
```
