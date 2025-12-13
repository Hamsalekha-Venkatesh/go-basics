package main

// needs to be in separate lines
import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello World from GO !", 1223)
	rest, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	defer rest.Body.Close()
	fmt.Println("response = ", rest)

}
