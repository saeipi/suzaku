package minioc

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/log"
	"suzaku/pkg/utils"
)

var minioc *Minioc

type Minioc struct {
	client *minio.Client
}

func NewMinioc() {
	var (
		operationID string
		minioUrl    *url.URL
		opt         minio.MakeBucketOptions
		client      *minio.Client
		exists      bool
		err         error
	)
	minioc = &Minioc{}

	operationID = utils.OperationIDGenerator()
	log.NewInfo(operationID, utils.GetSelfFuncName(), "minio config: ", config.Config.Credential.Minio)
	minioUrl, err = url.Parse(config.Config.Credential.Minio.Endpoint)
	if err != nil {
		log.NewError(operationID, utils.GetSelfFuncName(), "parse failed, please check config/config.yaml", err.Error())
		return
	}
	log.NewInfo(operationID, utils.GetSelfFuncName(), "Parse ok ", config.Config.Credential.Minio)
	client, err = minio.New(minioUrl.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Config.Credential.Minio.AccessKey, config.Config.Credential.Minio.SecretKey, ""),
		Secure: false,
	})
	log.NewInfo(operationID, utils.GetSelfFuncName(), "new ok ", config.Config.Credential.Minio)
	if err != nil {
		log.NewError(operationID, utils.GetSelfFuncName(), "init minio client failed", err.Error())
		return
	}
	opt = minio.MakeBucketOptions{
		Region:        config.Config.Credential.Minio.Location,
		ObjectLocking: false,
	}
	err = client.MakeBucket(context.Background(), config.Config.Credential.Minio.Bucket, opt)
	if err != nil {
		log.NewError(operationID, utils.GetSelfFuncName(), "MakeBucket failed ", err.Error())
		exists, err = client.BucketExists(context.Background(), config.Config.Credential.Minio.Bucket)
		if err == nil && exists {
			log.NewWarn(operationID, utils.GetSelfFuncName(), "We already own ", config.Config.Credential.Minio.Bucket)
		} else {
			if err != nil {
				log.NewError(operationID, utils.GetSelfFuncName(), err.Error())
			}
			log.NewError(operationID, utils.GetSelfFuncName(), "create bucket failed and bucket not exists")
			return
		}
	}
	minioc.client = client
	log.NewInfo(operationID, utils.GetSelfFuncName(), "minio create and set policy success")
	return
}

func PutObject(objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (info minio.UploadInfo, err error) {
	return minioc.client.PutObject(context.Background(), config.Config.Credential.Minio.Bucket, objectName, reader, objectSize, opts)
}
