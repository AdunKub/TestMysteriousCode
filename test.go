package main

import (
	base64 "encoding/base64"
	"fmt"
	"strings"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	splitWord := strings.Split(Sort(string(sd)), ":")
	fmt.Println(splitWord)
}

func Sort(param string) string {
	res := strings.Split(param, "")
	newStr := ""
	for i := len(res) - 1; i >= 0; i-- {
		newStr += res[i]
	}
	return newStr
}
