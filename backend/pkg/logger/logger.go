package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).Level(l).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.logger.Debug().Msgf(l.getMessage(message), args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn().Msgf(message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.logger.Error().Stack().Msgf(l.getMessage(message), args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.logger.Fatal().Msgf(l.getMessage(message), args...)

	os.Exit(1)
}

func (l *Logger) getMessage(message interface{}) string {
	switch msg := message.(type) {
	case error:
		return msg.Error()
	case string:
		return msg
	default:
		return fmt.Sprintf("%s message has unknown type %v", message, msg)
	}
}
