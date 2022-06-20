package main

import (
	"fmt"
	"time"
	"suzaku/examples/client/client"
	"suzaku/pkg/utils"
	"sync"
)

func main() {
	wscfg := &MainCfg{}
	utils.YamlToStruct("./configs/msg_gateway.yaml",wscfg)

	str,_ := utils.ObjToJson(wscfg)
	fmt.Println(str)

	var wg sync.WaitGroup
	wg.Add(1)

	manager := client.NewManager()
	manager.Run()

	wg.Wait()
}

type MainCfg struct {
	WsServer WsServer `yaml:"ws_server"`
}
type WsServer struct {
	WriteWait             int `yaml:"write_wait"`
	WriteWaitTime         time.Duration `yaml:"write_wait_time"`
	PongWait              int `yaml:"pong_wait"`
	PongWaitTime          time.Duration `yaml:"pong_wait_time"`
	PingPeriod            int `yaml:"ping_period"`
	PingPeriodTime        time.Duration `yaml:"ping_period_time"`
	MaxMessageSize        int `yaml:"max_message_size"`
	ReadBufferSize        int `yaml:"read_buffer_size"`
	WriteBufferSize       int `yaml:"write_buffer_size"`
	HeaderLength          int `yaml:"header_length"`
	ChanClientSendMessage int `yaml:"chan_client_send_message"`
	ChanServerReadMessage int `yaml:"chan_server_read_message"`
	ChanServerRegister    int `yaml:"chan_server_register"`
	ChanServerUnregister  int `yaml:"chan_server_unregister"`
	MaxConnections        int `yaml:"max_connections"`
	MinimumTimeInterval   int `yaml:"minimum_time_interval"`
}