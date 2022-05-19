package main

import (
	"flag"
	"fmt"
	"os"
)

var configFile = flag.String("config", "etc/idgen.yaml", "the config file")

func main() {
	flag.Parse()
	fmt.Println(*configFile)
	var args = os.Args
	fmt.Println(args)
}

/*
go run main.go -config etc/idgen.yaml
 */