package main

import (
	"suzaku/examples/chat/ws_server"
)

func main() {
	svr := ws_server.Init(2303)
	svr.Run()
}
