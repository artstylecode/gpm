package install

import
(
	"gpm/config"
	"os"
)
//Install 安装包
func Install(config *config.GpmConfig){
	packages := config.Packages
	for _, item := range packages {
		installPackage(item.Name, item.Version)
	}
}

func installPackage(name string , version string){
	env := os.Environ()
	println(env)
	procAttr := &os.ProcAttr{
		Env:env,
		Files:[] *os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	_,err := os.StartProcess("/bin/ls", []string{"ls","-l"}, procAttr)
	if err != nil {
		return
	}

}