package loggerNew

import (
	"log"
	"os"
)

type Logger struct {
	File        *os.File
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		File:        file,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}, nil

}

func (l *Logger) Info(msg string) {
	l.infoLogger.Output(2, "🛈 "+msg)
}

func (l *Logger) Error(msg string) {
	l.errorLogger.Output(2, "😱 "+msg)
}

type LoggerType string

const (
	InfoLoggerType  LoggerType = "Info"
	ErrorLoggerType LoggerType = "Error"
)
