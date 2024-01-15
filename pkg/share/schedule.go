package share

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron"
)

func taskWrapper(logger *log.Helper, task *Task) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("定时任务:[%s]发生了错误:Err:%s", task.Name, err)
		}
	}()

	logger.Debugf("开始:执行任务: %s, time now:%s", task.Name, time.Now().Format("2006-01-02 15:04:05"))

	// 执行任务
	task.TaskFunc()

	logger.Debugf("结束:执行任务: %s, time now:%s", task.Name, time.Now().Format("2006-01-02 15:04:05"))
}

type Task struct {
	Name      string
	TaskFunc  func()
	Scheduled string
}

func NewTask(name string, taskFunc func(), scheduled string) *Task {
	return &Task{Name: name, TaskFunc: taskFunc, Scheduled: scheduled}
}

func AddTask(logger *log.Helper, schedule *cron.Cron, task *Task) {
	if err := schedule.AddFunc(task.Scheduled, func() {
		// 执行定时任务
		taskWrapper(logger, task)
	}); err != nil {
		logger.Errorf("添加定时任务:[%s]失败:Err:%s", task.Name, err)
	}
}
