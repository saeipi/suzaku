package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var w = &sync.WaitGroup{}
	routineCh := make(chan int, 10)
	for i := 0; i < 100; i++ {
		w.Add(1)
		routineCh <- 1
		go RunTask(w, routineCh, i)
	}
	w.Wait()
	fmt.Println("[任务完成]")
}

func RunTask(w *sync.WaitGroup, ch chan int, index int) {
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("第%d个任务", index))
	w.Done()
	<-ch
}
