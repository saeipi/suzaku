package main

import (
	"flag"
	"suzaku/examples/go-zero/register/rpc/register/internal/config"
	"suzaku/examples/go-zero/register/rpc/register/internal/server"
	"suzaku/pkg/utils"
	"sync"
)

var configFile = flag.String("f", "etc/register-api.yaml", "the config file")

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	cfg := config.Config{}
	utils.YamlToStruct("/Users/saeipi/Desktop/Git/suzaku-open/suzaku/examples/go-zero/register/rpc/register/etc/register-api.yaml", &cfg)

	s := server.NewAuthRpcServer(&cfg)
	s.Run()

	wg.Wait()
}
