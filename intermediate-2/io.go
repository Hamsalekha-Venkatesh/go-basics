package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(reader io.Reader) {
	buffer := make([]byte, 1024)
	n, error := reader.Read(buffer)
	if error != nil {
		// return the bytes thats present - otherwise it wil be empty.
		fmt.Println("bytes read :  ", string(buffer[:n]))
	}
}

func writeToWriter(writer *bytes.Buffer, data string) {
	_, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("bytes written :  ", data)
	}
}

func closeResource(closer io.Closer) {
	err := closer.Close()
	if err == nil {
		log.Fatalln(errors.New("resource closed"))
	}
}

func main() {
	readFromReader(strings.NewReader("Hello world from string reader"))

	var writer bytes.Buffer
	writeToWriter(&writer, "Hello world from string writer")
	fmt.Println(writer.String())

	closeResource(io.NopCloser(&writer))
}

func writeToAFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(errors.New("open file error"))
		return
	}

	defer closeResource(file)

	// type conversions type(value); E.g. string(int val)

}
