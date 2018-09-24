package pylist

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// CreateDirIfNotExist use for creating dir if not exist for rotation log
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// RotateWriter struct log rotate writer. can be use as io.Writer interface
type RotateWriter struct {
	lock           sync.Mutex
	filename       string // should be set to the actual filename
	fp             *os.File
	timeFormat     string
	when           string
	interval       uint
	lastRotateTime time.Time
}

// NewRotateWriter Make a new RotateWriter. Return nil if error occurs during setup.
func NewRotateWriter(filename string, timeFormat string) *RotateWriter {

	w := &RotateWriter{lock: sync.Mutex{}, filename: filename, timeFormat: timeFormat}
	err := w.Rotate()
	if err != nil {
		return nil
	}
	return w
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock() // => lieu co can dung Mutex lock o day ko
	defer w.lock.Unlock()
	return w.fp.Write(append([]byte(time.Now().Format(w.timeFormat)), output...))
}

// Rotate => Perform the actual act of rotating and reopening file.
func (w *RotateWriter) Rotate(suffix ...string) (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	// Close existing file if open
	if w.fp != nil {
		err = w.fp.Close()
		w.fp = nil
		if err != nil {
			return
		}
	}
	timeSuffix := time.Now().Format(time.RFC3339)
	if suffix != nil && len(suffix) > 0 {
		timeSuffix = suffix[0]
	}

	// Rename dest file if it already exists
	_, err = os.Stat(w.filename)
	if err == nil {
		err = os.Rename(w.filename, w.filename+"."+timeSuffix)
		if err != nil {
			return
		}
	}

	// Create a file.
	w.fp, err = os.Create(w.filename)
	return err
}

// TimedRotatingFileHandler => this is the api of python in rotating log
// TimedRotatingFileHandler(filename [,when [,interval [,backupCount]]])
func (w *RotateWriter) TimedRotatingFileHandler(when string, interval uint) {
	w.when = when
	w.interval = interval
	w.lastRotateTime = time.Now()
}

func (w *RotateWriter) doRollover() {
	if w.interval == 0 {
		w.interval = 1
	}
	if w.when == "Day" || w.when == "D" || w.when == "d" {
		// get the string of current time
		currentTime := time.Now().Format(time.RFC3339)[:10]
		// get time from writer
		lasttime := w.lastRotateTime.Format(time.RFC3339)[:10]
		// check the difference to make Rotate or not
		if currentTime != lasttime {
			w.Rotate(currentTime)
			w.lastRotateTime = time.Now()
		}
	}

	if w.when == "Hour" || w.when == "H" || w.when == "h" {
		// get the string of current time
		currentTime := time.Now().Format(time.RFC3339)[:13]
		// get time from writer
		lasttime := w.lastRotateTime.Format(time.RFC3339)[:13]
		// check the difference to make Rotate or not
		if currentTime != lasttime {
			w.Rotate(currentTime)
			w.lastRotateTime = time.Now()
		}
	}

	if w.when == "Minute" || w.when == "M" || w.when == "m" {
		// get the string of current time
		currentTime := time.Now().Format(time.RFC3339)[:16]
		// get time from writer
		lasttime := w.lastRotateTime.Format(time.RFC3339)[:16]
		// check the difference to make Rotate or not
		if currentTime != lasttime {
			w.Rotate(currentTime)
			w.lastRotateTime = time.Now()
		}
	}
}

// Writer can be use as io.Writer interface. Use as an example
type Writer struct {
	io.Writer
	timeFormat string
}

func (w *Writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

// Logger object wrap log.Logger object and RotateWriter
type Logger struct {
	*log.Logger
	Writer *RotateWriter
}

// TimedRotating choose "D" for everyday rotating log, "H" for hours and "M" for minutes rotating log,
func (logger *Logger) TimedRotating(when string) {
	logger.Writer.TimedRotatingFileHandler(when, 1)
}

// Info => write to file an info log as [INFO] ...
func (logger *Logger) Info(output ...interface{}) {
	logger.Writer.doRollover()
	logger.Print("[INFO] ", fmt.Sprintln(output...))
}

// Warning => write to file log as [Warning] ...
func (logger *Logger) Warning(output ...interface{}) {
	logger.Writer.doRollover()
	logger.Print("[WARNING] ", fmt.Sprintln(output...))
	// logger.Printf("[WARNING] %+v\n", output...)
}

// Debug => write to file log as [Debug] ...
func (logger *Logger) Debug(output ...interface{}) {
	logger.Writer.doRollover()
	logger.Print("[DEBUG] ", fmt.Sprintln(output...))
	// logger.Printf("[DEBUG] %+v\n", output...)
}

// Error => write to file log as [Error] ...
func (logger *Logger) Error(output ...interface{}) {
	logger.Writer.doRollover()
	logger.Print("[ERROR] ", fmt.Sprintln(output...))
	// logger.Printf("[ERROR] %+v\n", output...)
}

// Fatal => write to file log as [Fatal] ...
func (logger *Logger) Fatal(output ...interface{}) {
	logger.Writer.doRollover()
	logger.Print("[FATAL] ", fmt.Sprintln(output...))
	// logger.Printf("[FATAL] %+v\n", output...)
}

// CreateLogger => Create a logger from a original file. This method will return logger object and accepted input as file name
func CreateLogger(filename string) Logger {
	// logFile, err := os.Create("app.log")
	// if err != nil {
	// 	panic(err)
	// }
	rotateWriter := NewRotateWriter(filename, "2006-01-02 15:04:05.000000 ")
	var logger = log.New(rotateWriter, "", log.Lshortfile)
	return Logger{logger, rotateWriter}
}
