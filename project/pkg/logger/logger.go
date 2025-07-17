package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type stdLogger struct {
	log *log.Logger
}

func New() Logger {
	return &stdLogger{
		log: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *stdLogger) Info(msg string, args ...interface{}) {
	l.log.Printf("INFO: "+msg, args...)
}

func (l *stdLogger) Error(msg string, args ...interface{}) {
	l.log.Printf("ERROR: "+msg, args...)
}
