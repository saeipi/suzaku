package main

import (
	"suzaku/internal/server/api_server"
)

func main() {
	//log.NewLogger("suzaku", "/logs/api.log")
	//go minioc.NewMinioc()
	api_server.NewApiServer().Run()
}
