package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// 连接Es
func connectEs() (*elastic.Client, error) {
	return elastic.NewClient(
		// 设置Elastic服务地址
		elastic.SetURL("http://127.0.0.1:9200"),
		// 是否转换请求地址，默认为true,当等于true时 请求http://ip:port/_nodes/http，将其返回的url作为请求路径
		elastic.SetSniff(false),
		// 心跳检查,间隔时间
		elastic.SetHealthcheckInterval(time.Second*5),
		// 设置错误日志
		elastic.SetErrorLog(log.New(os.Stderr, "ES-ERROR ", log.LstdFlags)),
		// 设置info日志
		elastic.SetInfoLog(log.New(os.Stdout, "ES-INFO ", log.LstdFlags)),
	)
}

// 测试连接
func TestConnectES(t *testing.T) {
	client, err := connectEs()
	if err != nil {
		t.Error(err)
		return
	}
	// 健康检查
	do, _ := client.ClusterHealth().Index().Do(context.TODO())
	fmt.Println("健康检查:", do)
}

// 创建索引(指定mapping)
func TestCreateIndexMapping(t *testing.T) {
	userMapping := `{
    "mappings":{
        "properties":{
            "name":{
                "type":"keyword"
            },
            "age":{
                "type":"byte"
            },
            "birth":{
                "type":"date"
            }
        }
    }
}`
	client, _ := connectEs()
	// 检测索引是否存在
	indexName := "go-test"
	// 创建上下文
	ctx := context.Background()
	exist, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		t.Errorf("检测索引失败:%s", err)
		return
	}
	if exist {
		t.Error("索引已经存在，无需重复创建！")
		return
	}
	res, err := client.CreateIndex(indexName).BodyString(userMapping).Do(ctx)
	if exist {
		t.Errorf("创建索引失败:%s", err)
		return
	}
	fmt.Println("创建成功:", res)
}

type UserInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Birth string `json:"birth"`
}

// 单条添加
func TestAddOne(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	// 创建userInfo
	userInfo := UserInfo{
		Name:  "张三",
		Age:   18,
		Birth: "1991-03-04",
	}
	res, err := client.Index().Index("go-test").Id("1").BodyJson(userInfo).Do(ctx)
	if err != nil {
		t.Errorf("添加失败:%s", err)
	}
	fmt.Println("添加成功", res)
}

// 批量添加
func TestBatchAdd(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	// 创建用户
	userNames := map[string]string{
		"李四": "1992-04-25",
		"张亮": "1994-07-15",
		"小明": "1991-12-03",
	}
	rand.Seed(time.Now().Unix())
	// 创建bulk
	userBulk := client.Bulk().Index("go-test")
	id := 4
	for n, b := range userNames {
		userTmp := UserInfo{Name: n, Age: rand.Intn(50), Birth: b}
		// 批量添加到bulk
		doc := elastic.NewBulkIndexRequest().Id(strconv.Itoa(id)).Doc(userTmp)
		userBulk.Add(doc)
		id++
	}
	// 检查被添加数据是否为空
	if userBulk.NumberOfActions() < 1 {
		t.Error("被添加的数据不能为空！")
		return
	}
	// 保存
	res, err := userBulk.Do(ctx)
	if err != nil {
		t.Errorf("保存失败:%s", err)
		return
	}
	fmt.Println("保存成功: ", res)
}

// 通过Script方式更新
func TestUpdateOneByScript(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()

	// 根据id更新
	res, err := client.Update().Index("go-test").Id("1").
		Script(elastic.NewScript("ctx._source.birth='1999-09-09'")).Do(ctx)
	if err != nil {
		t.Errorf("根据ID更新单条记录失败:%s", err)
		return
	}
	fmt.Println("根据ID更新成功:", res.Result)

	// 根据条件更新, update .. where name = '阿三'
	res2, err := client.UpdateByQuery("go-test").Query(elastic.NewTermQuery("name", "小明")).
		Script(elastic.NewScript("ctx._source.age=22")).ProceedOnVersionConflict().Do(ctx)
	if err != nil {
		t.Errorf("根据条件更新单条记录失败:%s", err)
		return
	}
	fmt.Println("根据条件更新成功:", res2.Updated)
}

// 使用Doc更新多个字段
func TestUpdateOneByDoc(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	res, _ := client.Update().Index("go-test").Id("5").Doc(map[string]interface{}{
		"name": "小白", "age": 30,
	}).Do(ctx)
	fmt.Println("更新结果:", res.Result)
}

// 批量修改
func TestBatchUpdate(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	bulkReq := client.Bulk().Index("go-test")
	for _, id := range []string{"4", "5", "6", "7"} {
		doc := elastic.NewBulkUpdateRequest().Id(id).Doc(map[string]interface{}{"age": 18})
		bulkReq.Add(doc)
	}
	// 被更新的数量不能小于0
	if bulkReq.NumberOfActions() < 0 {
		t.Error("被更新的数量不能为空")
		return
	}
	// 执行操作
	do, err := bulkReq.Do(ctx)
	if err != nil {
		t.Errorf("批量更新失败:%v", err)
		return
	}
	fmt.Println("更新成功:", do.Updated())
}

// 查询单条
func TestSearchOneEs(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	// 查找一条
	getResult, err := client.Get().Index("go-test").Id("1").Do(ctx)
	if err != nil {
		t.Errorf("获取失败: %s", err)
		return
	}
	// 提取查询结果(json格式)
	json, _ := getResult.Source.MarshalJSON()
	fmt.Printf("查询单条结果:%s \n", json)
}

// 查询多条
func TestSearchMoreES(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	searchResult, err := client.Search().Index("go-test").
		Query(elastic.NewMatchQuery("age", 18)).
		From(0).  //从第几条开始取
		Size(10). // 取多少条
		Pretty(true).
		Do(ctx)
	if err != nil {
		t.Errorf("获取失败: %s", err)
		return
	}
	// 定义用户结构体
	var userList []UserInfo
	for _, val := range searchResult.Each(reflect.TypeOf(UserInfo{})) {
		tmp := val.(UserInfo)
		userList = append(userList, tmp)
	}
	fmt.Printf("查询结果:%v\n", userList)
}

//  根据ID删除
func TestDelById(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	// 根据ID删除
	do, err := client.Delete().Index("go-test").Id("1").Do(ctx)
	if err != nil {
		t.Errorf("删除失败:%s", err)
		return
	}
	fmt.Println("删除成功: ", do.Result)
}

// 根据条件删除
func TestDelByWhere(t *testing.T) {
	client, _ := connectEs()
	ctx := context.Background()
	// 根据条件删除
	do, err := client.DeleteByQuery("go-test").Query(elastic.NewTermQuery("age", 18)).
		ProceedOnVersionConflict().Do(ctx)
	if err != nil {
		t.Errorf("删除失败:%s", err)
		return
	}
	fmt.Println("删除成功: ", do.Deleted)
}
