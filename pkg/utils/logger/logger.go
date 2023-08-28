package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	Logger = Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
