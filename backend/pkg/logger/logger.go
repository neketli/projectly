package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Fields defines structured key-value pairs for logging.
type Fields map[string]interface{}

// Interface defines the logging interface.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
	WithFields(fields Fields) Interface
}

// Logger provides structured logging capabilities.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New creates a new Logger with specified log level.
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

// Debug logs a debug message.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.logger.Debug().Msgf(l.getMessage(message), args...)
}

// Info logs an info message.
func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info().Msgf(message, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn().Msgf(message, args...)
}

// Error logs an error message.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.logger.Error().Stack().Msgf(l.getMessage(message), args...)
}

// Fatal logs a fatal message and exits.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.logger.Fatal().Msgf(l.getMessage(message), args...)

	os.Exit(1)
}

// WithFields returns a new Logger with structured fields attached to every log entry.
func (l *Logger) WithFields(fields Fields) Interface {
	ctx := l.logger.With()
	for k, v := range fields {
		switch val := v.(type) {
		case string:
			ctx = ctx.Str(k, val)
		case int:
			ctx = ctx.Int(k, val)
		case error:
			ctx = ctx.Err(val)
		default:
			ctx = ctx.Interface(k, val)
		}
	}
	child := ctx.Logger()
	return &Logger{logger: &child}
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
