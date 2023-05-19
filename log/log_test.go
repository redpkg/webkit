package log_test

import (
	"errors"
	"testing"

	"github.com/redpkg/webkit/log"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := log.Init(log.Config{
		Level:   "trace",
		Console: false,
	})
	if !assert.NoError(err) {
		return
	}

	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.Debug().Err(errors.New("foobar")).Msg("err message")
	log.Error().Err(outer()).Msg("err message")
	// log.Fatal().Msg("fatal")
	// log.Panic().Msg("panic")
}

func inner() error {
	return errors.New("inner error")
}

func middle() error {
	return errors.Join(errors.New("middle error"), inner())
}

func outer() error {
	return errors.Join(errors.New("outer error"), middle())
}
