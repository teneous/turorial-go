package basic

import (
	"fmt"
	strfunc "strings"
	"testing"
)

const blank = ""
const space = " "

func TestStringFunc(t *testing.T) {
	println(strfunc.ToLower("HELLO world"))
	println(strfunc.ToUpper("hello world"))

	const HTTP = "http"
	var serviceUrl string = "http://sso.syoka.com"
	if strfunc.HasPrefix(serviceUrl, HTTP) {
		println("service is correct")
	}

	println("count t ,num:", strfunc.Count(serviceUrl, "t"))
}

func TestStringSplit(t *testing.T) {
	music := "rock ,r&b ,electric"
	musics := splits(music)
	fmt.Printf("music col:%v", musics)
}

func splits(str string) []string {
	if len(str) != 0 && len(strfunc.ReplaceAll(str, space, blank)) != 0 {
		return strfunc.Split(str, ",")
	}
	return nil
}
