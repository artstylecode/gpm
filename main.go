package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//logic.Install("")
	getJson()
}

const (
	PATH    = "/blob/"
	GITPATH = "github.com/artstylecode/gpm/search/gpm.json"
)

func getJson() {
	reg, _ := regexp.Compile("[a-zA-Z.]+")
	matchstr := reg.FindAllString(GITPATH, 10)
	otherstr := matchstr[3:]
	matchstr = matchstr[0:3]
	matchstr = append(matchstr, PATH)
	fmt.Printf("%v\n", otherstr)
	fmt.Printf("%s", strings.Join(matchstr, "/"))
}
