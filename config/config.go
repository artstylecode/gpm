package config
import
(
	"encoding/json"
	"os"
)
const(
	//DefaultName 默认配置文件名称
	DefaultName="gpm.json"
)

//LoadConfig 加载项目依赖配置
func LoadConfig(path string)(*GpmConfig, error){

	if path==""{
		path=DefaultName
	}
	//fmt.Println(path)
	//加载配置文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config GpmConfig
	//解析配置
	err = json.NewDecoder(file).Decode(&config)
	if err != nil{
		return nil, err
	}
	//var reso DefaultResolver
	//config.parentConfig = resolveParentPacakges(reso, config)
	return &config, nil
}

//resolvePacakges 加载依赖
func resolveParentPacakges(resolver Resolver, config *GpmPackage)*GpmConfig{
	return nil
}