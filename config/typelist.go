package config
import "fmt"
//GpmConfig gpm配置信息
type GpmConfig struct{
	//version 依赖配置版本
	Version string `json:"version"`
	//packages 依赖包配置
	Packages []*GpmPackage `json:"packages"`
	//parentConfig 父级依赖配置
	parentConfig *GpmConfig
}

//(this GpmConfig)String 字符串转化方法
func (config *GpmConfig)String()string{
	return fmt.Sprintf("version:%s \n packages:%v", config.Version, config.Packages)
}
//GpmPackage gpm依赖包信息
type GpmPackage struct{
	Name string `json:"name"`
	Version string `json:"version"`
}
func (pack *GpmPackage)String()string{
	return fmt.Sprintf("name:%s \n version:%v", pack.Name, pack.Version)
}