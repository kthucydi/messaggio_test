package logging

// Package logging tune logrus logger for use with environment variable
// "LOG_LEVEL"  = {0..6} - logging level: 0 - panic .. 6 - trace, default - 4
// "LOG_FILE_PATH"  = Path with fileName for writing log, default logs/all.log

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

var (
	levels = map[int]logrus.Level{
		0: logrus.PanicLevel,
		1: logrus.FatalLevel,
		2: logrus.ErrorLevel,
		3: logrus.WarnLevel,
		4: logrus.InfoLevel,
		5: logrus.DebugLevel,
		6: logrus.TraceLevel,
	}
	LogrusEntry *logrus.Entry
	Log         = GetLogger()
)

// GetLogger create and return tuning logger
func GetLogger() Logger {
	Init()
	return Logger{LogrusEntry}
}

// Init create tuning logger
func Init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetOutput(setLogFileByEnv())
	setLogLevel(log)
	LogrusEntry = logrus.NewEntry(log)
}

// setLogLevel set log level by environment, if it empty set loglevel = info
func setLogLevel(log *logrus.Logger) {
	if logLevelString := os.Getenv("LOG_LEVEL"); logLevelString != "" {
		if logLevel, err := strconv.Atoi(logLevelString); err == nil {
			if logLevel < 7 && logLevel > -1 {
				log.SetLevel(levels[logLevel])
				return
			}
		}
	}
	log.SetLevel(levels[4])
	// log.SetLevel(logrus.InfoLevel) // also so
}

// setLogFile set log path and file by environment, if it empty set "logs/all.log"
func setLogFileByEnv() (allLogsFile *os.File) {
	if logFilePath := os.Getenv("LOG_FILE_PATH"); logFilePath != "" {
		allLogsFile = setLogFile(filepath.Split(logFilePath))
	} else {
		allLogsFile = setLogFile("logs/", "all.log")
	}
	return allLogsFile
}

// setLogFile with args
func setLogFile(logPath, logFile string) (allLogsFile *os.File) {
	err := os.MkdirAll(logPath, 0777)
	if err != nil {
		logrus.Error(err)
	}
	allLogsFile, err = os.OpenFile(logPath+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		logrus.Error(err, "log write only to Stdout")
		allLogsFile = os.Stdout
	}

	return allLogsFile
}

// func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
// 	return Logger{l.WithField(k, v)}
// }
