package log

import (
	"io"
	"os"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
)

type Config struct {
	Level   string `mapstructure:"level"`
	Console bool   `mapstructure:"console"`
}

func (c Config) level() zerolog.Level {
	switch strings.ToLower(c.Level) {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "off":
		return zerolog.Disabled
	default:
		return zerolog.WarnLevel
	}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// zerolog instance
var logger zerolog.Logger

// Init log
func Init(conf Config) error {
	zerolog.InterfaceMarshalFunc = json.Marshal
	zerolog.TimeFieldFormat = time.RFC3339

	var w io.Writer
	if conf.Console {
		w = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	} else {
		w = os.Stdout
	}

	logger = zerolog.New(w).
		Level(conf.level()).
		With().
		Timestamp().
		Logger()

	return nil
}

// Trace starts a new message with trace level.
func Trace() *zerolog.Event {
	return logger.Trace()
}

// Debug starts a new message with debug level.
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Info starts a new message with info level.
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn starts a new message with warn level.
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error starts a new message with error level.
func Error() *zerolog.Event {
	return logger.Error()
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method, which terminates the program immediately.
func Fatal() *zerolog.Event {
	return logger.Fatal()
}

// Panic starts a new message with panic level. The panic() function
// is called by the Msg method, which stops the ordinary flow of a goroutine.
func Panic() *zerolog.Event {
	return logger.Panic()
}
