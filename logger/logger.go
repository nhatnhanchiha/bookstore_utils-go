package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type bookstoreLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func getLevel() zapcore.Level {
	switch os.Getenv(envLogLevel) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	output := os.Getenv(envLogOutput)
	if output == "" {
		return "stdout"
	}

	return output
}

func init() {
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLevel()),
		OutputPaths: []string{getOutput()},
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		/*		{
				"level": "info"
				"time": "2006-01-02T15:04:05.000Z0700"
				"msg": "This is a log"
			},*/
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err.Error())
	}
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	err := log.log.Sync()
	if err != nil {
		return
	}
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	log.log.Info(msg, tags...)
	_err := log.log.Sync()
	if _err != nil {
		return
	}
}

func GetLogger() *logger {
	return &log
}
