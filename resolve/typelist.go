package resolve
import "fmt"
//GpmConfig gpm配置信息
type GpmConfig struct{
	//version 依赖配置版本
	version string `json:"version"`
	//packages 依赖包配置
	packages []*GpmPackage `json:"packages"`
	//parentConfig 父级依赖配置
	parentConfig *GpmConfig
}

//(this GpmConfig)String 字符串转化方法
func (config *GpmConfig)String()string{
	return fmt.Sprintf("version:%s \n packages:%v", config.version, config.packages)
}
//GpmPackage gpm依赖包信息
type GpmPackage struct{
	name string `json:"name"`
	version string `json:"version"`
}
func (pack *GpmPackage)String()string{
	return fmt.Sprintf("name:%s \n version:%v", pack.name, pack.version)
}