package logger

import (
	"fmt"
	"log"
	"os"
)

// TODO: Add 3rd-party logger library.

type Logger struct {
	logger *log.Logger
}

func NewLogger(prefix string) *Logger {
	l := log.New(os.Stdout, fmt.Sprintf("%s ", prefix), log.Ldate|log.Ltime)
	return &Logger{logger: l}
}

func (l *Logger) Log(text string) {
	l.logger.Println(text)
}

func (l *Logger) Error(err error) {
	l.logger.Fatalln(err)
}
