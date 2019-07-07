package config

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Resolver 查找上级包依赖配置
type Resolver interface {
	resolve(config *GpmPackage) *GpmConfig
}

//DefaultResolver 默认依赖包依赖处理类
type DefaultResolver struct{}

func (resolver DefaultResolver) resolve(config *GpmPackage) *GpmConfig {
	var parentConfig GpmConfig
	url := getConfigURL(config.Name, "master", DefaultName)
	file, err := getConfigContent(url)
	if err != nil {
		return nil
	}
	//解析配置
	err = json.NewDecoder(file).Decode(&parentConfig)
	return &parentConfig
}

func getConfigContent(url string) (*strings.Reader, error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	jsonTable := doc.Find("table.js-file-line-container tr")
	var configContent string
	jsonTable.Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, s *goquery.Selection) {
			configContent += s.Text()
		})
	})

	return strings.NewReader(configContent), nil

}
func getConfigURL(path string, branch string, configname string) string {
	reg, _ := regexp.Compile("[a-zA-Z.]+")
	matchstr := reg.FindAllString(path, 10)

	otherstr := strings.Join(matchstr[3:], "/")
	matchstr = matchstr[0:3]
	matchstr = append(matchstr, "blob")
	matchstr = append(matchstr, branch)
	matchstr = append(matchstr, configname)
	matchstr = append(matchstr, otherstr)
	return fmt.Sprintf("%s/%s", "https:/", strings.Join(matchstr, "/"))
}
