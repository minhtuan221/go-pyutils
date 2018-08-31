package pylist

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
)

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
	lock       sync.Mutex
	filename   string // should be set to the actual filename
	fp         *os.File
	timeFormat string
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
	w.lock.Lock()
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

// Writer can be use as io.Writer interface
type Writer struct {
	io.Writer
	timeFormat string
}

func (w *Writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

// Logger object wrap log.Logger object
type Logger struct {
	*log.Logger
	Writer *RotateWriter
}

func (logger *Logger) Info(output ...interface{}) {
	logger.Printf("[INFO] %+v\n", output)
}

func (logger *Logger) Warning(output ...interface{}) {
	logger.Printf("[WARNING] %+v\n", output)
}

func (logger *Logger) Debug(output ...interface{}) {
	logger.Printf("[DEBUG] %+v\n", output)
}

func (logger *Logger) Error(output ...interface{}) {
	logger.Printf("[ERROR] %+v\n", output)
}

func (logger *Logger) Fatal(output ...interface{}) {
	logger.Printf("[FATAL] %+v\n", output)
}

// Create a logger from a original file
func CreateLogger(filename string) Logger {
	// logFile, err := os.Create("app.log")
	// if err != nil {
	// 	panic(err)
	// }
	rotateWriter := NewRotateWriter(filename, "2006-01-02 15:04:05.000000 ")
	var logger = log.New(rotateWriter, "", log.Lshortfile)
	return Logger{logger, rotateWriter}
}
