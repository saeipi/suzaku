package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"suzaku/examples/etcd/etcd_master/cfg"
)

type LogMgr struct {
	db *mongo.Database
}

var (
	SG_LOGMGR *LogMgr
)

func InitLogMgr(cfg *cfg.Mongodb) (err error) {
	SG_LOGMGR = &LogMgr{}
	return
}
