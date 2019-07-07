package config

import (
	"encoding/json"
	"os"
)

const (
	//DefaultName 默认配置文件名称
	DefaultName = "gpm.json"
)

//LoadConfig 加载项目依赖配置
func LoadConfig(path string) (*GpmConfig, error) {

	if path == "" {
		path = DefaultName
	}
	var err error

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//}
	if file != nil {
		var config GpmConfig
		//解析配置
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			return nil, err
		}
		var reso DefaultResolver

		resolveParentPacakges(reso, &config)

		//
		return &config, nil
	}
	return nil, nil
}

//resolvePacakges 加载依赖
func resolveParentPacakges(resolver Resolver, config *GpmConfig) *GpmConfig {
	if config != nil && len(config.Packages) > 0 {
		for _, item := range config.Packages {
			config.parentConfig = resolver.resolve(item)
		}
		return resolveParentPacakges(resolver, config.parentConfig)
	}
	return nil
}
