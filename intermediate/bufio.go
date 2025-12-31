package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("hello world \n  sdfsdfsdf sdfsdfsdf sdfadfsdfsdfsdf sdgvsdgvsdgsdg sdfsdgfsdgfs!dfsdfsdfsdfsd"))
	//data := make([]byte, 20)
	data, err := reader.ReadString('!')
	if err != nil {
		return
	}
	fmt.Printf("%s\n", data)

	// ################## WRITER ###########################
	writer := bufio.NewWriter(os.Stdout)
	_, err = writer.WriteString("hello world!\n")
	if err != nil {
		return
	}
	_, err = writer.WriteString("My name is Hamsalekha Venkatesh \n how can I help you tdsdasdf")
	if err != nil {
		return
	}

	err = writer.Flush()
	if err != nil {
		return
	}

}
