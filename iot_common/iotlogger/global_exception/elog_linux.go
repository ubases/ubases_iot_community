//go:build linux
// +build linux

package global_exception

import (
	"fmt"
	"os"
	"syscall"
)

func RedirectStderr(logPath string) (err error) {
	//"./logs/iot_message_service.log.error"
	logFile, err := os.OpenFile(fmt.Sprintf("%s.error", logPath), os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	err = syscall.Dup3(int(logFile.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		return
	}
	return
}
