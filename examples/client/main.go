package main

import (
	"suzaku/examples/client/client"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	uid1 := "123"
	uid2 := "789"

	client1 := client.NewClient(uid1)
	client2 := client.NewClient(uid2)
	if client2 != nil {
	}
	client1.SendUser(uid2)

	wg.Wait()
}
