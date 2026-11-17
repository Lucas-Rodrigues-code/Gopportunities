package config

import (
	"io"
	"log"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer io.Writer := io.Writer(os.Stdout)
	logger *log.logger := log.New(out:writer,prefix:p, flag:log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer,"DEBUG", logger.flags()),
		info:    log.New(writer,"DEBUG", logger.flags()),
		warning: log.New(writer,"DEBUG", logger.flags()),
		err:     log.New(writer,"DEBUG", logger.flags()),
		writer:  writer,
	}
}
