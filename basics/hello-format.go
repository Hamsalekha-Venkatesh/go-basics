package main

// needs to be in separate lines
import (
	"fmt"
	net "net/http"
)

func main() {
	fmt.Println("hello World from GO !", 1223)
	rest, err := net.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	defer rest.Body.Close()
	fmt.Println("response = ", rest)
	fmt.Println(0 == 0)

}
