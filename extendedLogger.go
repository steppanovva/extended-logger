package exLog

import (
	"log"
	"os"
)

type LogLevel int

const (
	Error LogLevel = iota
	Warning
	Info
)

func (lvl LogLevel) isValid() bool {
	switch lvl {
	case Info, Warning, Error:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func (log *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.isValid() {
		return
	}
	log.logLevel = logLvl
}

func (log *LogExtended) PrintInfo(msg string) {
	log.println(Info, "INFO: ", msg)
}

func (log *LogExtended) PrintWarning(msg string) {
	log.println(Warning, "WARNING: ", msg)
}

func (log *LogExtended) PrintError(msg string) {
	log.println(Error, "ERROR: ", msg)
}

func (log *LogExtended) println(lvl LogLevel, prefix, msg string) {
	if log.logLevel < lvl {
		return
	}
	log.Logger.Println(prefix + msg)
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: Error,
	}
}
