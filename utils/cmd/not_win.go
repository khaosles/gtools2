//go:build !windows

package gcmd

import (
	"os/exec"
	"syscall"

	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: not_win.go
   @Author: khaosles
   @Time: 2023/6/26 18:12
   @Desc:
*/

// Asyn 异步执行cmd
func Asyn(cmdName string, args ...string) {
	glog.Debug("[CMD] Exec=> ", cmdName)
	glog.Debug("[CMD] Param=> ", args)
	cmd := exec.Command(cmdName, args...)
	// 仅在mac上有这个参数 windows无法使用
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		glog.Error("[CMD] Error:", err)
		return
	}
	glog.Debug("[CMD] Command start!")
}
