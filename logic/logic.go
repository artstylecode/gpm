package logic
import
(
	"gpm/config"
	"gpm/install"
	//"fmt"
)
//Install 安装依赖
func Install(path string){
	config,err := config.LoadConfig(path)
	if err != nil {
		return 
	}
	install.Install(config)

	//fmt.Println(conf.String())


}
