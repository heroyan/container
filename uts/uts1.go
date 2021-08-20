package uts

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Uts1() {
	if len(os.Args) < 2 {
		logrus.Errorf("missing commands")
		return
	}
	switch os.Args[1] {
	case "run":
		run1()
	default:
		logrus.Errorf("wrong command")
		return
	}
}
func run1() {
	logrus.Infof("Running %v", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	check(cmd.Run())
}
func check(err error) {
	if err != nil {
		logrus.Errorln(err)
	}
}
