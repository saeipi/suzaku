package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"suzaku/pkg/common/log/monlog"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/mongodb"
)

type LogMgr struct {
	db *mongo.Database
}

var (
	SG_LOGMGR *LogMgr
)

func InitLogMgr(cfg config.MongoConfig) (err error) {
	SG_LOGMGR = &LogMgr{}
	SG_LOGMGR.db, _ = mongodb.InitMongo(cfg)
	monlog.Shared().SetDB(SG_LOGMGR.db)
	return
}
