package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println(pid)
}
