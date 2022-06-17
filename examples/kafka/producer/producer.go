package main

// 生产者

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

var Topic = "ws2ms_chat"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//配置发布者
	config := sarama.NewConfig()
	// 成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 设置超时时间 这个超时时间一旦过期，新的订阅者在这个超时时间后才创建的，就不能订阅到消息了
	config.Producer.Timeout = 5 * time.Second

	// 连接发布者，并创建发布者实例
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	// 程序退出时释放资源
	defer client.Close()
	// 例子1发单个消息
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	// 设置一个逻辑上的分区名
	msg.Topic = Topic
	// 这个是发布的内容
	content := "this is a test log"
	send01(client, msg, content)

	//例子2发多个消息
	for _, word := range []string{"Welcome11", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		send01(client, msg, word)
	}
	wg.Wait()
}

//发消息
func send01(client sarama.SyncProducer, msg *sarama.ProducerMessage, content string) {
	msg.Value = sarama.StringEncoder(content)
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
