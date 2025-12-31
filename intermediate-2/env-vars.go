package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//os.Setenv("USER", "root")
	//os.Setenv("HOME", "/home/user")

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Printf("USER: %s \n  HOME_DIT: %s\n", user, home)

	for _, e := range os.Environ() {
		kvpair := strings.Split(e, "=")
		fmt.Println(kvpair)
	}
}
