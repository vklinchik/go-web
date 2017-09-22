package xlog

import (
	"os"
	"github.com/sirupsen/logrus"
)

type DebugLevel int

const(
	LogLevelDebug DebugLevel	= 1
	LogLevelInfo DebugLevel 	= 2
	LogLevelWarn DebugLevel 	= 3
	LogLevelError DebugLevel 	= 4 
)
var Logger = &logrus.Logger{}

func init() {
	Logger.Out = os.Stdout //os.Stderr
	Logger.Formatter = new(logrus.TextFormatter) // log.SetFormatter(&log.JSONFormatter{})
	Logger.Level = logrus.InfoLevel //default is set to info
}

// ConfigureLogger configures the level of the logger
func ConfigureLogger(level DebugLevel) {
	switch level {
	case LogLevelDebug:
		Logger.Level = logrus.DebugLevel
	case LogLevelInfo:
		Logger.Level = logrus.InfoLevel
	case LogLevelWarn:
		Logger.Level = logrus.InfoLevel //TODO: ???
	case LogLevelError:
		Logger.Level = logrus.InfoLevel //TODO: ???
	}
}

func GetLogger() *logrus.Logger {
	return Logger
}

func Debug(args ...interface{}) {
	Logger.Debug(args)
}

func Info(args ...interface{}) {
	Logger.Info(args)
}

func Warning(args ...interface{}) {
	Logger.Warning(args)
}
