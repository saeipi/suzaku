package repository_mongo

import (
	"context"
	"suzaku/pkg/common/config"
	"time"
)

func NewContext() (ctx context.Context, cancelFunc context.CancelFunc) {
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.Timeout)*time.Second)
	return
}
