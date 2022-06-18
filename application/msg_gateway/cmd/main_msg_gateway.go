package main

import (
	"suzaku/application/msg_gateway/internal/server"
	"suzaku/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
