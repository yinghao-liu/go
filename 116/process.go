package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func startProcessIfNotExist(processName string, param []string) (err error) {
	pids, _ := pgrep(processName)
	if len(pids) > 0 {
		err = errors.New("process exist")
		return
	}
	return startProcess(processName, param)
}

func startProcess(processName string, param []string) (err error) {
	cmd := exec.Command(processName, param...)
	fmt.Printf("%s starting...\n", processName)
	err = cmd.Start()
	if nil != err {
		fmt.Printf("%s start failed\n", processName)
		fmt.Printf("error is %s\n", err.Error())
		return
	}
	fmt.Printf("%s started\n", processName)

	return
}

// kill process by name
func killProcessByName(processName string) (err error) {
	pids, err := pgrep(processName)
	if nil != err {
		return
	}
	for _, j := range pids {
		pid, err := strconv.Atoi(j)
		if nil != err {
			fmt.Printf("error is %s\n", err.Error())
			return err
		}

		err = killProcessByPid(pid)
		if nil != err {
			fmt.Printf("error is %s\n", err.Error())
			return err
		}
	}
	return
}

// kill process by pid
func killProcessByPid(pid int) (err error) {
	proc, err := os.FindProcess(pid)
	if nil != err {
		fmt.Printf("error is %s\n", err.Error())
		return
	}
	err = proc.Kill()
	if nil != err {
		fmt.Printf("error is %s\n", err.Error())
		return
	}
	return
}

// pgrep only pid
func pgrep(processName string) (pids []string, err error) {
	name := path.Base(processName)
	cmd := exec.Command("pgrep", name)
	data, err := cmd.Output()
	if nil != err {
		fmt.Printf("error is %s\n", err.Error())
		return
	}
	pidstr := string(data)
	pidstr = strings.Trim(pidstr, "\n")
	pids = strings.Split(pidstr, "\n")
	fmt.Printf("data is [%#v]\n", pids)
	return
}
