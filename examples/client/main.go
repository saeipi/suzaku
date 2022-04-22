package main

import (
	"strconv"
	"suzaku/examples/client/client"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 10000; i = i + 2 {
		newConnection(strconv.Itoa(i), strconv.Itoa(i+1))
	}
	wg.Wait()
}

func newConnection(uid1 string, uid2 string) {
	uid1 = "uid" + uid1
	uid2 = "uid" + uid2
	client1 := client.NewClient(uid1)
	client2 := client.NewClient(uid2)
	if client2 != nil {
	}
	client1.SendUser(uid2)
}
