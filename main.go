package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//logic.Install("")
	url := getConfigUrl()
	_, _ = getConfigContent(url)
	// reader := strings.NewReader(string(str))
	// var config config.GpmConfig
	// json.NewDecoder(reader).Decode(&config)

	// fmt.Println(config.String())
}

const (
	PATH    = "blob/master"
	GITPATH = "github.com/artstylecode/gpm/gpm.json"
)

func getConfigContent(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	buff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile("test.txt", buff, 0666)
	//fmt.Println(string(buff))
	return nil, err

}
func getConfigUrl() string {
	reg, _ := regexp.Compile("[a-zA-Z.]+")
	matchstr := reg.FindAllString(GITPATH, 10)

	otherstr := strings.Join(matchstr[3:], "/")
	matchstr = matchstr[0:3]
	matchstr = append(matchstr, PATH)
	matchstr = append(matchstr, otherstr)
	return fmt.Sprintf("%s/%s", "https:/", strings.Join(matchstr, "/"))
}
