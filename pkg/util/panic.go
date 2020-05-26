package util

import (
	"fmt"
	"log"
	"runtime"
)

func RecoverPanic() {
	err := recover()
	if err != nil {
		log.Print("panic: ", err, " stack: ", getStackInfo())
	}
}

// 获取Panic堆栈信息
func getStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return fmt.Sprintf("%s", buf[:n])
}
