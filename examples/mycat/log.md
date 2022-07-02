### linux 查看日志常用命令
#### tail:
```
-n 是显示行号；相当于nl命令；例子如下：

tail -100f test.log 实时监控100行日志

tail -n 10 test.log 查询日志尾部最后10行的日志;

tail -n +10 test.log 查询10行之后的所有日志;
```

#### head:
```
跟tail是相反的，tail是看后多少行日志；例子如下：

head -n 10 test.log 查询日志文件中的头10行日志;

head -n -10 test.log 查询日志文件除了最后10行的其他所有日志;
```

#### cat：
```
tac是倒序查看，是cat单词反写；例子如下：
cat -n test.log |grep "debug" 查询关键字的日志

清空日志文件
第一种：cat /dev/null > filename
第二种：: > filename
第三种：> filename
第四种：echo "" > filename
第五种：echo > filename
```

#### 应用场景一：按行号查看---过滤出关键字附近的日志
```
1）cat -n test.log |grep "debug" 得到关键日志的行号

2）cat -n test.log |tail -n +92|head -n 20 选择关键字所在的中间一行. 然后查看这个关键字前10行和后10行的日志:

tail -n +92表示查询92行之后的日志

head -n 20 则表示在前面的查询结果里再查前20条记录
```

#### 应用场景二：根据日期查询日志
```
sed -n '/2014-12-17 16:17:20/,/2014-12-17 16:17:36/p' test.log

特别说明:上面的两个日期必须是日志中打印出来的日志,否则无效；

先 grep '2014-12-17 16:17:20' test.log 来确定日志中是否有该 时间点
```

#### 应用场景三：日志内容特别多，打印在屏幕上不方便查看
```
(1)使用more和less命令,

如： cat -n test.log |grep "debug" |more 这样就分页打印了,通过点击空格键翻页

(2)使用 >xxx.txt 将其保存到文件中,到时可以拉下这个文件分析

如：cat -n test.log |grep "debug" >debug.txt
```