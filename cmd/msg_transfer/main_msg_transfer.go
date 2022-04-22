package main

import (
	"suzaku/internal/msg_transfer/logic"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	logic.Initialize()
	logic.Run()

	wg.Wait()
}
