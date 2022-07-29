```
Durablity:持久化选项，Durable（持久化保存），Transient（即时保存）, 
持久化保存会在RabbitMQ宕机或者重启后，未消费的消息仍然存在，即时保存在RabbitMQ宕机或者重启后交换机会不存在。需要重新定义该Exchange。
```

```
队列类型
Classic Quorum Stream
```