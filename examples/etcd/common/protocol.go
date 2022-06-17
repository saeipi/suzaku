package common

import (
	"context"
	"encoding/json"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)

// 定时任务
type Job struct {
	Name     string `json:"name"`      // 任务名
	Command  string `json:"command"`   // shell命令
	CronExpr string `json:"cron_expr"` // cron表达式
}

// 任务执行状态
type JobExecute struct {
	Job        *Job               // 任务信息
	PlanTime   time.Time          // 理论上的调度时间
	RealTime   time.Time          // 实际的调度时间
	CancelCtx  context.Context    // 任务command的context
	CancelFunc context.CancelFunc // 用于取消command执行的cancel函数
}

// 任务执行结果
type JobExecuteResult struct {
	Execute   *JobExecute // 执行状态
	Output    []byte      // 脚本输出
	Err       error       // 脚本错误原因
	StartTime time.Time   // 启动时间
	EndTime   time.Time   // 结束时间
}

// 变化事件
type JobEvent struct {
	EventType int //  SAVE, DELETE
	Job       *Job
}

// 任务调度计划
type JobSchedulePlan struct {
	Job      *Job                 // 要调度的任务信息
	Expr     *cronexpr.Expression // 解析好的cronexpr表达式
	NextTime time.Time            // 下次调度时间
}

// 提取worker的IP
func ExtractWorkerIP(regKey string) string {
	return strings.TrimPrefix(regKey, JOB_WORKER_DIR)
}

// 反序列化Job
func DeserializeJob(value []byte) (ret *Job, err error) {
	var (
		job *Job
	)

	job = &Job{}
	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 任务变化事件有2种：1）更新任务 2）删除任务
func BuildJobEvent(eventType int, job *Job) (jobEvent *JobEvent) {
	return &JobEvent{
		EventType: eventType,
		Job:       job,
	}
}

// 构造任务执行计划
func BuildJobSchedulePlan(job *Job) (jobSchedulePlan *JobSchedulePlan, err error) {
	var (
		expr *cronexpr.Expression
	)
	// 解析JOB的cron表达式
	if expr, err = cronexpr.Parse(job.CronExpr); err != nil {
		return
	}
	// 生成任务调度计划对象
	jobSchedulePlan = &JobSchedulePlan{
		Job:      job,
		Expr:     expr,
		NextTime: expr.Next(time.Now()),
	}
	return
}

// 构造执行状态信息
func BuildJobExecute(jobSchedulePlan *JobSchedulePlan) (jobExecuteInfo *JobExecute) {
	jobExecuteInfo = &JobExecute{
		Job:      jobSchedulePlan.Job,
		PlanTime: jobSchedulePlan.NextTime, // 计算调度时间
		RealTime: time.Now(),               // 真实调度时间
	}
	jobExecuteInfo.CancelCtx, jobExecuteInfo.CancelFunc = context.WithCancel(context.TODO())
	return
}

// 从etcd的key中提取任务名
// /cron/jobs/job10抹掉/cron/jobs/
func ExtractJobName(jobKey string) string {
	return strings.TrimPrefix(jobKey, JOB_SAVE_DIR)
}

// 从 /cron/killer/job10提取job10
func ExtractKillerName(killerKey string) string {
	return strings.TrimPrefix(killerKey, JOB_KILLER_DIR)
}
