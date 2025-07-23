package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime

	return &Logger{
		debug:   log.New(writer, colorCyan, flags),
		info:    log.New(writer, colorGreen, flags),
		warning: log.New(writer, colorYellow, flags),
		err:     log.New(writer, colorRed, flags),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println("DEBUG:", fmt.Sprint(v...), colorReset)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println("INFO:", fmt.Sprint(v...), colorReset)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println("WARNING:", fmt.Sprint(v...), colorReset)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println("ERROR:", fmt.Sprint(v...), colorReset)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf("DEBUG: "+format+colorReset, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf("INFO: "+format+colorReset, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf("WARNING: "+format+colorReset, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf("ERROR: "+format+colorReset, v...)
}