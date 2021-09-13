package muon

import (
	"io"
	"log"
)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

func NewDefaultLogger(out io.Writer) *Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	logger := &Logger{
		Info:    log.New(out, "[info] ", flags),
		Warning: log.New(out, "[warning] ", flags),
		Error:   log.New(out, "[error] ", flags),
		Fatal:   log.New(out, "[fatal] ", flags),
	}
	return logger
}
