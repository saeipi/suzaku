package main

import (
	"suzaku/internal/server/api_server"
	"suzaku/pkg/common/log"
)

func main() {
	log.NewLogger("suzaku", "/logs/api.log")
	//go minioc.NewMinioc()
	api_server.NewApiServer().Run()
}
