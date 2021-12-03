package uts

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

func Uts2() {
	if len(os.Args) < 2 {
		logrus.Errorf("missing commands")
		return
	}
	switch os.Args[1] {
	case "run":
		run2()
	case "child":
		child()
	default:
		logrus.Errorf("wrong command")
		return
	}
}

func run2() {
	logrus.Info("Setting up...")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	check(cmd.Run())
}

func child() {
	logrus.Infof("Running %v", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cg()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//check(syscall.Sethostname([]byte("newhost")))
	check(cmd.Run())
}
