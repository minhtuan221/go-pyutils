package main

import (
	py "Pyutils/pyutils/rotationLog"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	py.CreateDirIfNotExist("logs/rotatation")
	logger := py.CreateLogger("logs/rotatation/app.log")
	logger.Writer.TimedRotatingFileHandler("m", 1)
	logger.Info("this is logger")
	logger.Writer.Rotate()
	logger.Info("this is second logger in new file")
	logger.Writer.Rotate("2018-08-31")
	logger.Info("this is third logger in new file")
	for index := 0; index < 100; index++ {
		time.Sleep(10 * time.Second)
		logger.Info(index, "This is automatic logging every 10 second for testing")
	}
	fmt.Println("Waiting for all threads to finish")

}
