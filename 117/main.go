package main

import (
	"log"
	"os"
	"path"
	"syscall"
)

func main() {
	// 打开文件锁
	LockFile := path.Join("lock")
	lock, e := os.Create(LockFile)
	if e != nil {
		log.Printf("main: 创建文件锁失败: %s", e)
		os.Exit(-1)
	}
	defer os.Remove(LockFile)
	defer lock.Close()

	// 尝试独占文件锁
	e = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if e != nil {
		log.Printf("main: 独占文件锁失败: %s", e)
		os.Exit(-1)
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)

	// your code ...
}
