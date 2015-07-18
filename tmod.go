package main

import (
	"flag"
	"os"
	"syscall"
	"time"
)

func main() {
	var fileFlag = flag.String("file", "", "file to modify date")
	var offsetFlag = flag.Int("offset", -300, "number of seconds")

	flag.Parse()

	if *fileFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	t := time.Now()
	seconds := int(t.UnixNano()/1e9) + *offsetFlag

	timeval := syscall.Timeval{int32(seconds), 0}
	timevals := []syscall.Timeval{timeval, timeval}
	err := syscall.Utimes(*fileFlag, timevals)
	if err != nil {
		panic(err)
	}
}
