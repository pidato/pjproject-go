package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

var (
	logWrite = func(entry pjsua2.LogEntry) {
		msg := entry.GetMsg()
		strings.Replace(msg, "\r", "", -1)

		if msg[len(msg)-1] == '\n' {
			msg = msg[37 : len(msg)-1]
		}

		Infof("[PJSIP] %v", msg)
	}
)

type LogWriter struct {
	name string
}

func (l *LogWriter) Write(entry pjsua2.LogEntry) {
	logWrite(entry)
}

type logger struct{}

type Logger interface {
	Infof(format string, v ...interface{})

	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})

	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

func SetLogger(l zerolog.Logger) {
	log = l
}

func SetLogLevel(level zerolog.Level) {
	zerolog.TimeFieldFormat = timeFormat
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	log = zerolog.New(output).Level(level).With().Timestamp().Logger()
}

func Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	log.Panic().Msgf(format, v...)
}
