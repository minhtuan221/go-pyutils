package main

import (
	"fmt"
	"sync"
	"time"

	py "github.com/minhtuan221/go-pyutils/pyutils/rotationLog"
)

var wg sync.WaitGroup

func main() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// url := path.Join(dir, "../logs/rotatation")
	// fmt.Println(url)
	py.CreateDirIfNotExist("./logs/rotatation")
	logger := py.CreateLogger("logs/rotatation/app.log")
	logger.TimedRotating("m")
	logger.Info("this is logger")
	logger.Writer.Rotate()
	logger.Info("this is second logger in new file")
	logger.Writer.Rotate("2018-08-31")
	logger.Info("this is third logger in new file")
	for index := 0; index < 10000; index++ {
		time.Sleep(1 * time.Second)
		go func() {
			logger.Info(index, "This is automatic logging every 10 second for testing")
		}()
	}
	fmt.Println("Waiting for all threads to finish")

}
