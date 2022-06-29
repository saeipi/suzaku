package main

import (
	"suzaku/application/user/internal/server"
	"suzaku/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
