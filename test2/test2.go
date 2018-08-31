package main

import (
	py "Pyutils/pyutils/rotationLog"
)

func main() {
	py.CreateDirIfNotExist("logs/rotatation")
	logger := py.CreateLogger("logs/rotatation/app.log")
	logger.Info("this is logger")
	logger.Writer.Rotate()
	logger.Info("this is second logger in new file")
	logger.Writer.Rotate("2018-08-31")
	logger.Info("this is third logger in new file")
}
