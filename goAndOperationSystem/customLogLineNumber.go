package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE2 = "/tmp/mGo.log"

func main() {
	f, err := os.OpenFile(LOGFILE2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
