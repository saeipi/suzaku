package log

import (
	"go.mongodb.org/mongo-driver/mongo"
	"suzaku/pkg/common/log/monlog"
	"suzaku/examples/etcd/etcd_worker/cfg"
)

// 任务执行日志
type JobLog struct {
	JobName      string `json:"job_name" bson:"job_name"`           // 任务名字
	Command      string `json:"command" bson:"command"`             // 脚本命令
	Err          string `json:"err" bson:"err"`                     // 错误原因
	Output       string `json:"output" bson:"output"`               // 脚本输出
	PlanTime     int64  `json:"plan_time" bson:"plan_time"`         // 计划开始时间
	ScheduleTime int64  `json:"schedule_time" bson:"schedule_time"` // 实际调度时间
	StartTime    int64  `json:"start_time" bson:"start_time"`       // 任务执行开始时间
	EndTime      int64  `json:"end_time" bson:"end_time"`           // 任务执行结束时间
}

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

func SaveJobLog(log *JobLog) {
	monlog.Insert("jobs", log)
}
