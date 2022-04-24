package executor

import (
	"math/rand"
	"os/exec"
	"suzaku/micro/etcd/common"
	"time"
)

// 任务执行器
type Executor struct {
}

var (
	SG_EXECUTOR *Executor
)

//  初始化执行器
func InitExecutor() (err error) {
	SG_EXECUTOR = &Executor{}
	return
}

// 执行一个任务
func (executor *Executor) ExecuteJob(info *common.JobExecute) {
	var (
		cmd     *exec.Cmd
		err     error
		output  []byte
		result  *common.JobExecuteResult
		jobLock *JobLock
	)
	// 任务结果
	result = &common.JobExecuteResult{
		Execute: info,
		Output:  make([]byte, 0),
	}
	// 初始化分布式锁
	jobLock = SG_JOBMGR.CreateJobLock(info.Job.Name)
	// 记录任务开始时间
	result.StartTime = time.Now()
	// 上锁
	// 随机睡眠(0~1s)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	err = jobLock.TryLock()
	defer jobLock.Unlock()
	if err != nil { // 上锁失败
		result.Err = err
		result.EndTime = time.Now()
	} else {
		// 上锁成功后，重置任务启动时间
		result.StartTime = time.Now()

		// 执行shell命令
		cmd = exec.CommandContext(info.CancelCtx, "/bin/bash", "-c", info.Job.Command)

		// 执行并捕获输出
		output, err = cmd.CombinedOutput()

		// 记录任务结束时间
		result.EndTime = time.Now()
		result.Output = output
		result.Err = err
	}
	// 任务执行完成后，把执行的结果返回给Scheduler，Scheduler会从executingTable中删除掉执行记录
	SG_SCHEDULER.PushJobResult(result)
}
