package mylog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const (
	debugLevel = 0
	releaseLevel = 1
	errorLevel = 2
	fatalLevel = 3
)

const (
	printDebugLevel = "[debug] "
	printReleaseLevel = "[release] "
	printErrorLevel = "[error] "
	printFatalLevel = "[fatal] "
)

type Logger struct {
	level int
	baseLogger *log.Logger
	baseFile *os.File
}

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger {
		level: debugLevel,
		baseLogger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func New(strLevel string, pathName string, flag int) (*Logger, error) {
	var level int
	switch strings.ToLower(strLevel) {
	case "debug":
		level = debugLevel
	case "release":
		level = releaseLevel
	case "warn":
		level = releaseLevel
	case "error":
		level = errorLevel
	case "fatal":
		level = fatalLevel
	default:
		return nil, errors.New("unknow level: " + strLevel)
	}

	var baseLogger *log.Logger
	var baseFile *os.File
	if pathName != "" {
		now := time.Now()
		fileName := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		file, err := os.Create(path.Join(pathName, fileName))
		if err != nil {
			return nil, err
		}
		baseLogger = log.New(io.MultiWriter(file, os.Stdout), "", flag)
		baseFile = file
	} else {
		baseLogger = log.New(os.Stdout, "", flag)
	}
	logger := new(Logger)
	logger.level = level
	logger.baseLogger = baseLogger
	logger.baseFile = baseFile
	defaultLogger = logger
	return logger, nil
}

func (logger *Logger) Close() {
	if logger.baseFile != nil {
		logger.baseFile.Close()
	}
	logger.baseLogger = nil
	logger.baseFile = nil
}

func (logger *Logger) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	if level < logger.level {
		return
	}
	if logger.baseLogger == nil {
		panic("logger closed")
	}
	format = printLevel + format
	logger.baseLogger.Output(2, fmt.Sprintf(format, a...))
	if level == fatalLevel {
		os.Exit(1)
	}
}

// Debug log
func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

// Release log
func (logger *Logger) Release(format string, a ...interface{}) {
	logger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}

// Error log
func (logger *Logger) Error(format string, a ...interface{}) {
	logger.doPrintf(errorLevel, printErrorLevel, format, a...)
}

// Fatal panic
func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.doPrintf(fatalLevel, printFatalLevel, format, a...)
}

func Export(logger *Logger) {
	if logger != nil {
		defaultLogger = logger
	}
}

// Debug print
func Debug(format string, a ...interface{}) {
	defaultLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

// Release print
func Release(format string, a ...interface{}) {
	defaultLogger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}

// Error print
func Error(format string, a ...interface{}) {
	defaultLogger.doPrintf(errorLevel, printErrorLevel, format, a...)
}

// Fatal print
func Fatal(format string, a ...interface{}) {
	defaultLogger.doPrintf(fatalLevel, printFatalLevel, format, a...)
}

func Close() {
	defaultLogger.Close()
}










