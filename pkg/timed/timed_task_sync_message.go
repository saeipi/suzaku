package timed

func (t *timer) AddTaskSyncMessage(key string) {
	var args = TaskArgs{
		Key:  key,
		Name: "同步部门",
		Spec: "CRON_TZ=Asia/Shanghai 30 08 ? * *",
	}
	var job = TaskJob{Params: nil, Task: args, Func: SyncMessage}
	t.AddJob(job.Task, job)
}

func SyncMessage(args interface{}) (err error) {
	return
}
