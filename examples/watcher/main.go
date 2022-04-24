package main

import (
	"fmt"
	"suzaku/examples/watcher/etcd"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	watcher, err := etcd.NewWatcher("suzaku:///OnlineMessageRelay", []string{"127.0.0.1:2379"}, 5000)
	if err != nil {
		fmt.Println(err)
	}
	if watcher == nil {

	}

	wg.Wait()
}
