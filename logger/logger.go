package logger

import (
	"log"
	"os"
	"sync"
)

type LoggingConfig struct {
	level int
}

const (
	Error int = 1
	Warn      = 2
	Info      = 3
	Debug     = 4
	Trace     = 5
)

var lv *LoggingConfig

func SetLogLevel(level int) {
	lv = &LoggingConfig{}

	switch level {
	case Trace:
		lv.level = Trace
		break
	case Debug:
		lv.level = Debug
		break
	case Info:
		lv.level = Info
		break
	case Warn:
		lv.level = Warn
		break
	case Error:
		lv.level = Error
		break
	default:
		lv.level = Info
		break
	}
}

func getLogLevel() *LoggingConfig {
	return lv
}

type CustomLogger struct {
	logLevel *LoggingConfig
	trace    *log.Logger
	debug    *log.Logger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
}

func (l *CustomLogger) Trace(format string, v ...interface{}) {
	if l.logLevel.level >= Trace {
		l.trace.Printf(format, v...)
	}
}

func (l *CustomLogger) Debug(format string, v ...interface{}) {
	if l.logLevel.level >= Debug {
		l.debug.Printf(format, v...)
	}
}

func (l *CustomLogger) Info(format string, v ...interface{}) {
	if l.logLevel.level >= Info {
		l.info.Printf(format, v...)
	}
}

func (l *CustomLogger) Warn(format string, v ...interface{}) {
	if l.logLevel.level >= Warn {
		l.warn.Printf(format, v...)
	}
}

func (l *CustomLogger) Error(format string, v ...interface{}) {
	if l.logLevel.level >= Error {
		l.error.Printf(format, v...)
	}
}

func getLogFlags() int {
	return log.LstdFlags
}

var lock = &sync.Mutex{}
var logger *CustomLogger

func GetLoggerInstance() *CustomLogger {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()

		if logger == nil {
			return &CustomLogger{
				logLevel: getLogLevel(),
				trace:    log.New(os.Stdout, "TRACE: ", getLogFlags()),
				debug:    log.New(os.Stdout, "DEBUG: ", getLogFlags()),
				info:     log.New(os.Stdout, "INFO: ", getLogFlags()),
				warn:     log.New(os.Stdout, "WARN: ", getLogFlags()),
				error:    log.New(os.Stdout, "ERROR: ", getLogFlags()),
			}
		}
	}
	return logger
}
