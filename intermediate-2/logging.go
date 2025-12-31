package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("AppLogger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // R|W perm
	if err != nil {
		return
	}
	InfoLogger.SetOutput(file)
	ErrorLogger.SetOutput(file)
	WarnLogger.SetOutput(file)

	InfoLogger.Println("Hello World with info log")
	ErrorLogger.Println("Hello World with info")
	ErrorLogger.Println("Hello World with errors")

	defer file.Close()

}

var (
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
	WarnLogger  = log.New(os.Stdout, "WARN: ", log.Lshortfile)
)
