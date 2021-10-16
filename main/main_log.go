package main

import (
	"log"
	"mylucky/mylog"
)

func main() {
	name := "Leaf"
	mylog.Debug("My name is %v", name)
	mylog.Release("My name is %v", name)
	mylog.Error("My name is %v", name)

	logger, err := mylog.New("release", ".", log.Lshortfile|log.LstdFlags)
	if err != nil {
		return
	}
	defer mylog.Close()
	logger.Debug("will not print")
	logger.Release("My name is %v", name)

	mylog.Export(logger)

	mylog.Debug("will not print")
	mylog.Release("My name is %v", name)
}
