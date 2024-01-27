package logger

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var (
	once sync.Once
	log  *zerolog.Logger
)

func GetLoggerWithContext(ctx context.Context) *zerolog.Logger {
	once.Do(func() {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger := zerolog.New(output).With().Timestamp().Ctx(ctx).Logger()
		log = &logger
	})
	return log
}

func GetLogger() *zerolog.Logger {
	once.Do(func() {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger := zerolog.New(output).With().Timestamp().Logger()
		log = &logger
	})
	return log
}
