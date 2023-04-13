package task

import "fmt"

func init() {
	scheduledTaskList = append(scheduledTaskList, scheduled{f: clearDeviceImportErrorFile, corm: "0 1 * * ?", name: "定期删除dome定时任务"})
}

// 每天凌晨删除空文件
func clearDeviceImportErrorFile() {
	fmt.Printf("dome")
}
