package timed

const (
	TimerTaskSyncMessage = "SyncMessage"
)

func (t *timer) InitTimerTask() {
	t.AddTaskSyncMessage(TimerTaskSyncMessage)
}
