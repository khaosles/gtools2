package gcmd

import (
	"bufio"
	"os/exec"

	"github.com/axgle/mahonia"
	"github.com/khaosles/gtools2/core/log"
)

/*
   @File: cmd.go
   @Author: khaosles
   @Time: 2023/6/3 20:16
   @Desc:
*/

func Exec(cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	glog.Debug("[CMD] Exec=> ", cmdName)
	glog.Debug("[CMD] Param=> ", args)
	bytes, err := cmd.Output()
	if err != nil {
		glog.Error("[CMD] Error ", err.Error())
		return "", err
	}
	resp := string(bytes)
	glog.Debug("[CMD] ", "Command finished")
	return resp, nil
}

// Sync 同步执行cmd并打印输出
func Sync(cmdName string, args ...string) error {
	glog.Debug("[CMD] Exec=> ", cmdName)
	glog.Debug("[CMD] Param=> ", args)
	cmd := exec.Command(cmdName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		glog.Error("[CMD] Error:", err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		glog.Error("[CMD] Error:", err)
		return err
	}
	if err := cmd.Start(); err != nil {
		glog.Error("[CMD] Error:", err)
		return err
	}
	encoder := mahonia.NewEncoder("utf8")
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			glog.Debug("[CMD] ", encoder.ConvertString(scanner.Text()))
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			glog.Error("[CMD] Error: ", encoder.ConvertString(scanner.Text()))
		}
	}()
	if err := cmd.Wait(); err != nil {
		glog.Error("[CMD] Error:", err)
		return err
	}

	glog.Debug("[CMD] ", "Command finished")
	return nil
}

// Asyn 异步执行cmd
func Asyn(cmdName string, args ...string) {
	glog.Debug("[CMD] Exec=> ", cmdName)
	glog.Debug("[CMD] Param=> ", args)
	cmd := exec.Command(cmdName, args...)
	// 仅在mac上有这个参数 windows无法使用
	//cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		glog.Error("[CMD] Error:", err)
		return
	}
	glog.Debug("[CMD] Command start!")
}
