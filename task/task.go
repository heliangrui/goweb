package task

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var scheduledTaskList = make([]scheduled, 0)

type scheduled struct {
	corm string
	name string
	f    func()
}

var rn *cron.Cron

func init() {
	rn = cron.New()
}

// InitTask 定时任务启动方法 分钟级corn表达式
func InitTask() {
	if len(scheduledTaskList) > 0 {
		for _, task := range scheduledTaskList {
			_, err := rn.AddFunc(task.corm, task.f)
			if err != nil {
				zap.L().Error("定时任务启动失败！:"+err.Error(), zap.String("name:", task.name))
			}
		}
		rn.Start()
	}

}
