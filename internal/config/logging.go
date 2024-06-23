package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	fileMode = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	filePerm = 0644 // File permissions: readable and writable by owner, readable by group and others
)

func MustConfigureLogging(logging Logging) *slog.Logger {
	const operation = "config.MustConfigureLogging"
	var output *os.File

	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(logging.Level)); err != nil {
		panic(fmt.Errorf("%s: %w", operation, err))
	}

	if strings.ToUpper(logging.Output) == "STDOUT" {
		output = os.Stdout
	} else {
		file, err := os.OpenFile(logging.Output, fileMode, filePerm)
		if err != nil {
			panic(fmt.Errorf("%s: %w", operation, err))
		}
		output = file
	}

	handler := slog.NewTextHandler(output, &slog.HandlerOptions{Level: logLevel})
	return slog.New(handler)
}
