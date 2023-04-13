package cmd

import (
	"go-web/utils/contants"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

const (
	DefaultParam = "undefined"
	SvgCmd       = "svgCompile"
)

// SvgCompileCmd svg 编译
func SvgCompileCmd(classId string, deviceId string, pelName string, exportPath string, rootPath string, projectId string) {
	if classId == "" {
		classId = DefaultParam
	}
	if deviceId == "" {
		deviceId = DefaultParam
	}
	if pelName == "" {
		pelName = DefaultParam
	}
	if exportPath == "" {
		exportPath = DefaultParam
	}
	command := exec.Command(SvgCmd, contants.JsonTypeDevice, projectId, classId, deviceId, pelName, exportPath)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	zap.L().Info("cmd ->" + command.String())
	command.Start()

}

// SvgCompileCmd svg 编译scene
func SvgCompileSceneCmd(flowchartPath string, destPath string, projectId string) {

	command := exec.Command(SvgCmd, "page", projectId, flowchartPath, destPath)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	zap.L().Info("cmd ->" + command.String())
	command.Start()

}
