package main

import (
	"github.com/minio/minio-go/v7"
)

func main() {

}

type client struct {
	client         *minio.Client
	bucketName     string
	targetFilePath string
	useSSL         bool
	location       string
}

//func NewClient() *client {
//	var endpoint = "127.0.0.1:9000"   //minio地址
//	var accessKeyID = "17098899839"   //账号
//	var secretAccessKey = "360001969" //密码
//	var useSSL = false                //使用http或https
//
//
//	// Initialize minio client object.
//	target, err := minio.New(endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
//		Secure: useSSL,
//	})
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return &client{
//		client:         target,
//		bucketName:     "files",                //目标bucket
//		location:       "cn-north-1",           //放在哪个地方，这里为中国
//		targetFilePath: "docs/documents/简介.md", //资源文件夹的路径
//	}
//}
//
////checkBucket 检测目标bucket是否存在，不存在就创建一个
//func (c *client) checkBucket() {
//	isExists, err := c.client.BucketExists(c.bucketName)
//	if err != nil {
//		log.Println("check bucket exist error ")
//		return
//	}
//
//	if !isExists {
//		err2 := c.client.MakeBucket(c.bucketName, c.location)
//		if err2 != nil {
//			log.Println("MakeBucket error ")
//			fmt.Println(err2)
//			return
//		}
//		log.Printf("Successfully created %s\n", c.bucketName)
//	}
//}
//
////UpLoadFile 将整个目录都上传,遍历所有文件夹及子文件夹上传
//func (c *client) UpLoadFile(path string) {
//	rd, _ := ioutil.ReadDir(path)
//	for _, fi := range rd {
//		if fi.IsDir() {
//			c.UpLoadFile(path + "/" + fi.Name())
//		} else {
//			if strings.Index(fi.Name(), ".dat") > 0 {
//				fullPath := path + "/" + fi.Name()
//				rawPathLength := len(c.targetFilePath)
//				objectName := fullPath[rawPathLength:]
//				log.Printf("fullPath=%s  ,objectName=%s\n", fullPath, objectName)
//				n, err := c.client.FPutObject(c.bucketName, objectName, fullPath, minio.PutObjectOptions{
//					ContentType: "",
//				})
//				if err != nil {
//					fmt.Println(err)
//					return
//				}
//				fmt.Println("Successfully uploaded bytes: ", n)
//			}
//		}
//	}
//}
