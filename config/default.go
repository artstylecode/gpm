package config

//Resolver 查找上级包依赖配置
type Resolver interface {
	resolve(config *GpmPackage) *GpmConfig
}

//DefaultResolver 默认依赖包依赖处理类
type DefaultResolver struct{}

func (resolver DefaultResolver) resolve(config *GpmPackage) *GpmConfig {
	return nil
}
