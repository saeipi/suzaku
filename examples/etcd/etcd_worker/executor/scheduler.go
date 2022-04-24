package executor

import (
	"fmt"
	"suzaku/examples/etcd/common"
	"suzaku/examples/etcd/etcd_worker/log"
	"time"
)

// 任务调度
type Scheduler struct {
	jobEventChan      chan *common.JobEvent              // etcd任务事件队列
	jobPlanTable      map[string]*common.JobSchedulePlan // 任务调度计划表
	jobExecutingTable map[string]*common.JobExecute      // 任务执行表
	jobResultChan     chan *common.JobExecuteResult      // 任务结果队列
}

var (
	SG_SCHEDULER *Scheduler
)

// 初始化调度器
func InitScheduler() (err error) {
	SG_SCHEDULER = &Scheduler{
		jobEventChan:      make(chan *common.JobEvent, 1000),
		jobPlanTable:      make(map[string]*common.JobSchedulePlan),
		jobExecutingTable: make(map[string]*common.JobExecute),
		jobResultChan:     make(chan *common.JobExecuteResult, 1000),
	}
	// 启动调度协程
	go SG_SCHEDULER.scheduleLoop()
	return
}

// 处理任务事件
func (s *Scheduler) handleJobEvent(jobEvent *common.JobEvent) {
	var (
		jobSchedulePlan *common.JobSchedulePlan
		jobExecuteInfo  *common.JobExecute
		jobExecuting    bool
		jobExisted      bool
		err             error
	)

	switch jobEvent.EventType {
	case common.JOB_EVENT_SAVE: // 保存任务事件
		if jobSchedulePlan, err = common.BuildJobSchedulePlan(jobEvent.Job); err != nil {
			return
		}
		s.jobPlanTable[jobEvent.Job.Name] = jobSchedulePlan
	case common.JOB_EVENT_DELETE: // 删除任务事件
		if jobSchedulePlan, jobExisted = s.jobPlanTable[jobEvent.Job.Name]; jobExisted {
			delete(s.jobPlanTable, jobEvent.Job.Name)
		}
	case common.JOB_EVENT_KILL: // 强杀任务事件
		// 取消掉Command执行, 判断任务是否在执行中
		if jobExecuteInfo, jobExecuting = s.jobExecutingTable[jobEvent.Job.Name]; jobExecuting {
			jobExecuteInfo.CancelFunc() // 触发command杀死shell子进程, 任务得到退出
		}
	}
}

// 02 重新计算任务调度状态
func (s *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		jobPlan  *common.JobSchedulePlan
		now      time.Time
		nearTime *time.Time
	)
	// 如果任务表为空话，睡眠1S
	if len(s.jobPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}
	// 当前时间
	now = time.Now()
	// 遍历所有任务
	for _, jobPlan = range s.jobPlanTable {
		// func (t Time) Before(u Time) bool，判断时间 t 是否在时间 u 的前面
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			s.TryStartJob(jobPlan)
			jobPlan.NextTime = jobPlan.Expr.Next(now) // 更新下次执行时间
		}
		// 统计最近一个要过期的任务时间 func (t Time) After(u Time) bool，判断时间 t 是否在时间 u 的后面
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}
	// 下次调度间隔（最近要执行的任务调度时间 - 当前时间）
	scheduleAfter = (*nearTime).Sub(now)
	return
}

// 尝试执行任务
func (s *Scheduler) TryStartJob(jobPlan *common.JobSchedulePlan) {
	// 调度 和 执行 是2件事情
	var (
		jobExecute   *common.JobExecute
		jobExecuting bool
	)
	// 执行的任务可能运行很久, 1分钟会调度60次，但是只能执行1次, 防止并发！
	// 如果任务正在执行，跳过本次调度
	if jobExecute, jobExecuting = s.jobExecutingTable[jobPlan.Job.Name]; jobExecuting {
		return
	}
	// 构建执行状态信息
	jobExecute = common.BuildJobExecute(jobPlan)
	// 保存执行状态
	s.jobExecutingTable[jobPlan.Job.Name] = jobExecute
	// 执行任务
	fmt.Println("执行任务:", jobExecute.Job.Name, jobExecute.PlanTime, jobExecute.RealTime)
	SG_EXECUTOR.ExecuteJob(jobExecute)
}

// 处理任务结果
func (s *Scheduler) handleJobResult(result *common.JobExecuteResult) {
	var (
		jobLog *log.JobLog
	)
	// 删除执行状态
	delete(s.jobExecutingTable, result.Execute.Job.Name)

	// 生成执行日志
	if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		jobLog = &log.JobLog{
			JobName:      result.Execute.Job.Name,
			Command:      result.Execute.Job.Command,
			Output:       string(result.Output),
			PlanTime:     result.Execute.PlanTime.UnixNano() / 1000 / 1000,
			ScheduleTime: result.Execute.RealTime.UnixNano() / 1000 / 1000,
			StartTime:    result.StartTime.UnixNano() / 1000 / 1000,
			EndTime:      result.EndTime.UnixNano() / 1000 / 1000,
		}
		if result.Err != nil {
			jobLog.Err = result.Err.Error()
		} else {
			jobLog.Err = ""
		}
		log.SaveJobLog(jobLog)
	}
}

// 01 调度协程
func (s *Scheduler) scheduleLoop() {
	var (
		jobEvent      *common.JobEvent
		scheduleAfter time.Duration
		scheduleTimer *time.Timer
		jobResult     *common.JobExecuteResult
	)
	// 初始化一次(1秒)
	scheduleAfter = s.TrySchedule()
	// 调度的延迟定时器
	scheduleTimer = time.NewTimer(scheduleAfter)

	// 定时任务common.Job
	for {
		select {
		case jobEvent = <-s.jobEventChan: // 监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			s.handleJobEvent(jobEvent)
		case <-scheduleTimer.C: // 最近的任务到期了
		case jobResult = <-s.jobResultChan: // 监听任务执行结果
			s.handleJobResult(jobResult)
		default:
		}
		// 调度一次任务
		scheduleAfter = s.TrySchedule()
		// 重置调度间隔
		scheduleTimer.Reset(scheduleAfter)
	}
}

// 推送任务变化事件
func (s *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	s.jobEventChan <- jobEvent
}

// 回传任务执行结果
func (s *Scheduler) PushJobResult(jobResult *common.JobExecuteResult) {
	s.jobResultChan <- jobResult
}
