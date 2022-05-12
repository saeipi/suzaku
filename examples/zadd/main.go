package main

import (
	"fmt"
	"suzaku/pkg/common/redis"
	"suzaku/pkg/utils"
	"time"
)

func main() {
	var key = "sid2"
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		m := Member{fmt.Sprintf("名字%d", i), fmt.Sprintf("手机%d", i),utils.GetCurrentTimestampByMill()}
		redis.ZAddObj(key, m.Timestamp, m)
	}

	jsStr := redis.ZRevRangeByScore(key, utils.GetCurrentTimestampByMill(), 0, 5)

	jsStr = redis.ZRangeByScore(key, 0, utils.GetCurrentTimestampByMill(), 0, 5)

	var members = make([]Member, 0)
	utils.JsonToObj(jsStr, &members)
	fmt.Println("")
}

type Member struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Timestamp int64 `json:"timestamp"`
}

/*
// 【小于】 1652346077845 的近50个数值,降序
ZREVRANGEBYSCORE szk:sid1 1652346077845 0 limit 1 50

//【小于等于】 1652346077845 的近50个数值,降序
ZREVRANGEBYSCORE szk:sid1 1652346077845 0 limit 0 2

// 【大于等于】1652345403591；小于1652346078847 的近2个数值,升序
ZRANGEBYSCORE szk:sid1 1652345403591 1652346078847 limit 0 2

// 【大于】1652345403591；小于1652346078847 的近2个数值,升序
ZRANGEBYSCORE szk:sid1 1652345403591 1652346078847 limit 1 2
*/
