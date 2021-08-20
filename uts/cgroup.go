package uts

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func cg() {
	cgPath := "/sys/fs/cgroup/"
	pidsPath := filepath.Join(cgPath, "pids")
	// 在/sys/fs/cgroup/pids下创建container目录
	os.Mkdir(filepath.Join(pidsPath, "container"), 0755)
	// 设置最大进程数目为20
	check(ioutil.WriteFile(filepath.Join(pidsPath, "container/pids.max"), []byte("20"), 0700))
	// 将notify_on_release值设为1，当cgroup不再包含任何任务的时候将执行release_agent的内容
	check(ioutil.WriteFile(filepath.Join(pidsPath, "container/notify_on_release"), []byte("1"), 0700))
	// 加入当前正在执行的进程
	check(ioutil.WriteFile(filepath.Join(pidsPath, "container/pids.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}
