package main

import (
	"fmt"
	"net/url"
)

func main() {
	var rawUrl = "https://raw.githubusercontent.com/path?username=hamsav&pass=hamsa1234#bio"
	parsedUrl, err := url.Parse(rawUrl)
	if err == nil {
		fmt.Println(parsedUrl.RawQuery)
	} else {
		fmt.Println("error is ", err.Error())
	}
}
