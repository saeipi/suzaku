package main

// apache-jmeter-5.5/bin
// sh jmeter

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"net/http"
	"strconv"
	"time"
)

var (
	client *goredislib.Client
	engine *gin.Engine
	prot   = 9168
	key    = "user:uid:1"
)

func main() {
	initRedisClient()
	initGin()
	run()
}

func initRedisClient() {
	client = goredislib.NewClient(&goredislib.Options{
		Addr: "127.0.0.1:6379",
	})
}

func initGin() {
	engine = gin.Default()
	engine.GET("/spike", spike)
}

func run() {
	addr := ":" + strconv.Itoa(prot)
	engine.Run(addr)
}

func success(ctx *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data[0],
	})
}

func spike(c *gin.Context) {
	userJson, _ := client.Get(context.TODO(), key).Result()
	if userJson != "" {
		success(c, userJson)
		return
	}

	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// 创建一个redisync的实例，用于获得相互排斥的结果。
	// lock.
	rs := redsync.New(pool)

	// 通过对所有实例使用相同的名称来获得一个新的mutex，希望
	// same lock.
	mutexname := "my-global-mutex"
	mutex := rs.NewMutex(mutexname)

	// 为我们给定的mutex获取锁。在此成功之后，在我们解锁之前，其他人不能获得相同的锁（相同的mutex名称）。
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.
	fmt.Println("Do your work that requires the lock.")
	//time.Sleep(40 * time.Second)
	userJson = getUserInfo()

	// 释放锁，以便其他进程或线程可以获得锁。
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
	success(c, userJson)
}

func getUserInfo() string {
	fmt.Println("读取mysql数据库")
	time.Sleep(5)
	userJson := "{uid:1,username:\"saeipi\"}"
	client.Set(context.TODO(), key, userJson, 0)
	return userJson
}
