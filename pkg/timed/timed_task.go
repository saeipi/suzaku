package timed

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
)

type Timer interface {
	InitTimerTask()
	AddFunc(args TaskArgs, function func()) (id cron.EntryID, err error)
	AddJob(args TaskArgs, job interface{ Run() }) (id cron.EntryID, err error)
	FindTask(key string) (*Task, bool)
	StartTask(key string)
	StopTask(key string)
	Remove(key string, id int)
	Clear(key string)
	Close()
}

//默认Job接口
type TaskJob struct {
	Params interface{}
	Task   TaskArgs
	Func   func(args interface{}) (err error)
}

func (job TaskJob) Run() {
	defer func() {
		if err := recover(); &err != nil {
			fmt.Println(err)
		}
	}()
	job.Func(job.Params)
}

type TaskArgs struct {
	Key  string
	Name string
	Spec string
}

// 小写开头字段无须外部赋值
type Task struct {
	TaskArgs
	isJob bool
	cron  *cron.Cron
}

// timer 定时任务管理
type timer struct {
	taskList map[string]*Task
	sync.Mutex
}

func NewTimer() Timer {
	return &timer{taskList: make(map[string]*Task)}
}

// 通过函数的方法添加任务
func (t *timer) AddFunc(args TaskArgs, function func()) (id cron.EntryID, err error) {
	if args.Key == "" || args.Spec == "" || function == nil {
		err = errors.New("参数错误")
		return
	}
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[args.Key]; ok == false {
		var task = Task{}
		task.TaskArgs = args
		task.cron = cron.New()
		task.isJob = false
		t.taskList[task.Key] = &task
	}
	id, err = t.taskList[args.Key].cron.AddFunc(args.Spec, function)
	t.taskList[args.Key].cron.Start()
	return
}

// 通过接口的方法添加任务
func (t *timer) AddJob(args TaskArgs, job interface{ Run() }) (id cron.EntryID, err error) {
	if args.Key == "" || args.Spec == "" || job == nil {
		err = errors.New("参数错误")
		return
	}
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[args.Key]; ok == false {
		var task = Task{}
		task.TaskArgs = args
		task.cron = cron.New()
		task.isJob = true
		t.taskList[task.Key] = &task
	}
	id, err = t.taskList[args.Key].cron.AddJob(args.Spec, job)
	t.taskList[args.Key].cron.Start()
	return
}

// FindTaskn 获取对应key的task 可能会为空
func (t *timer) FindTask(key string) (*Task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[key]
	return v, ok
}

// StartTask 开始任务
func (t *timer) StartTask(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[key]; ok {
		v.cron.Start()
	}
}

// StopTask 停止任务
func (t *timer) StopTask(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[key]; ok {
		v.cron.Stop()
	}
}

// Remove 删除指定任务
func (t *timer) Remove(key string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[key]; ok {
		v.cron.Remove(cron.EntryID(id))
	}
}

// Clear 清除任务
func (t *timer) Clear(key string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[key]; ok {
		v.cron.Stop()
		delete(t.taskList, key)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.cron.Stop()
	}
}
