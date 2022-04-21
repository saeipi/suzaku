package main

import (
	"suzaku/internal/server/server_api"
	"suzaku/pkg/common/minioc"
	"suzaku/pkg/common/log"
)

func main() {
	log.NewLogger("suzaku", "./logs/api.log")
	go minioc.NewMinioc()
	server_api.ApiServer.Run()
}
