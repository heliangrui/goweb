package main

import (
	"go-web/core/config"
	"go-web/core/mqttConfig"
	"go-web/router"
	"go-web/task"
)

// @title 应用服务
// @version 1.0
// @description 应用服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9837
// @BasePath
func main() {
	//初始化路由与日志
	r := router.Init()
	//初始化db
	config.InitDB()
	mqttConfig.InitMQTT()
	//启动定时任务
	task.InitTask()
	// 启动程序
	router.Start(r)

}
