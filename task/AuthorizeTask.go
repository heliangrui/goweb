package task

func init() {
	initAuthorizeData()
	scheduledTaskList = append(scheduledTaskList, scheduled{f: initAuthorizeData, corm: "0/30 * * * ? ", name: "加载dome定时任务"})
}

func initAuthorizeData() {

}
