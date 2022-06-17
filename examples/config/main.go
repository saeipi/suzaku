package main

import (
	"fmt"
	"suzaku/pkg/common/config"
)

func main() {
	var cfg = config.Config
	fmt.Println(cfg.Monlog.CommitTimeout)
}
