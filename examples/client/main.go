package main

import (
	"suzaku/examples/client/client"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	manager := client.NewManager()
	manager.Run()

	wg.Wait()
}
