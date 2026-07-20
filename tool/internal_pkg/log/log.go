package log

import (
	"log"
	"os"
)

var Verbose bool

type Logger interface {
	Printf(format string, a ...interface{})
}

type LoggerFunc func(format string, a ...interface{})

func (f LoggerFunc) Printf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

var defaultLogger Logger = log.New(os.Stderr, "", 0)

func DefaultLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func SetDefaultLogger(l Logger) { _ = "STUB: not implemented"; return }

func Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }
