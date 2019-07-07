package install

import (
	"gpm/config"
	"os"
	"sync"
)

//Install 安装包
func Install(config *config.GpmConfig) {
	packages := config.Packages
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(packages))
	for _, item := range packages {
		installPackage(item.Name, item.Version)
	}
}

func installPackage(name string, version string) {
	env := os.Environ()
	goroot := os.Getenv("GOROOT")
	goExcuteFile := goroot + "/bin/go"

	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	println("开始安装包:" + name + ".....")
	pid, err := os.StartProcess(goExcuteFile, []string{"go", "get", "-u", name}, procAttr)
	if err != nil {
		return
	}
	pid.Wait()

	println(name + "安装成功..")
}
