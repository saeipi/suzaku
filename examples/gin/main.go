package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", hello)
	engine.Run(":9166")
}

func hello(c *gin.Context) {
	fmt.Println("访问hello Api")
	c.SecureJSON(0, "访问成功")
}
