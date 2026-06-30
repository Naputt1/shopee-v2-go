package goshopee

import (
	"fmt"
	"io"
	"log"
)

type LeveledLoggerInterface interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type LeveledLogger struct {
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewLeveledLogger(debug, info, warn, err io.Writer) *LeveledLogger {
	return &LeveledLogger{
		DebugLogger: log.New(debug, "DEBUG: ", log.Ltime|log.Lshortfile),
		InfoLogger:  log.New(info, "INFO: ", log.Ltime|log.Lshortfile),
		WarnLogger:  log.New(warn, "WARN: ", log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(err, "ERROR: ", log.Ltime|log.Lshortfile),
	}
}

func (l *LeveledLogger) Debugf(format string, v ...interface{}) {
	if l.DebugLogger != nil {
		l.DebugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Infof(format string, v ...interface{}) {
	if l.InfoLogger != nil {
		l.InfoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Warnf(format string, v ...interface{}) {
	if l.WarnLogger != nil {
		l.WarnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Errorf(format string, v ...interface{}) {
	if l.ErrorLogger != nil {
		l.ErrorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}
