package logger

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() zeroLog.Logger {
	zeroLog.TimeFieldFormat = time.RFC3339

	if os.getEnv("GIN_MODE") == "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	return log.Logger

}
