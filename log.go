package pj

import (
	"fmt"
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

		var source string
		if msg[len(msg)-1] == '\n' {
			if len(msg) >= 37 {
				source = strings.TrimSpace(msg[0:37])

				idx := strings.Index(source, " ")
				if idx > -1 {
					source = strings.TrimSpace(source[idx:])
				}

				idx = strings.Index(source, " ")
				if idx > -1 {
					source = strings.TrimSpace(source[:idx])
				}

				//file = file[idx+1:]
				msg = strings.TrimSpace(msg[37 : len(msg)-1])
			}
			//
		} else {
			source = ""
		}

		level := entry.GetLevel()

		var event *zerolog.Event
		switch level {
		case 8:
			event = log.Trace()
		case 7:
			event = log.Trace()
		case 6:
			event = log.Debug()
		case 5:
			event = log.Debug()
		case 4:
			event = log.Info()
		case 3:
			event = log.Warn()
		case 2:
			event = log.Error()
		case 1:
			event = log.Error()
		default:
			return
		}

		//threadId := entry.GetThreadId()
		//threadName := entry.GetThreadName()
		event.
			Str("source", source).
			//Str("thread", threadName).
			//Int64("tid", threadId).
			Msg(fmt.Sprintf("%s %s", logPrefix, msg))
	}
)

func pjToZeroLogLevel(level int) zerolog.Level {
	switch level {
	case 7:
		return zerolog.TraceLevel
	case 6:
		return zerolog.DebugLevel
	case 5:
		return zerolog.InfoLevel
	case 4:
		return zerolog.InfoLevel
	case 3:
		return zerolog.WarnLevel
	case 2:
		return zerolog.ErrorLevel
	case 1:
		return zerolog.PanicLevel
	default:
		return zerolog.NoLevel
	}
}

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

const logPrefix = "[PJ]"

func Infof(format string, v ...interface{}) {
	log.Info().Msgf(fmt.Sprintf("%s %s", logPrefix, format), v...)
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Msgf(fmt.Sprintf("%s %s", logPrefix, format), v...)
}

func Warnf(format string, v ...interface{}) {
	log.Warn().Msgf(fmt.Sprintf("%s %s", logPrefix, format), v...)
}

func Errorf(format string, v ...interface{}) {
	log.Error().Msgf(fmt.Sprintf("%s %s", logPrefix, format), v...)
}

func Panicf(format string, v ...interface{}) {
	log.Panic().Msgf(fmt.Sprintf("%s %s", logPrefix, format), v...)
}
