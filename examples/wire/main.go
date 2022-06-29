package main

// go get github.com/google/wire/cmd/wire

import (
	"suzaku/examples/wire/event"
)

// 使用wire后
func main() {
	e := event.InitializeEvent("hello_world")
	e.Start()
}
