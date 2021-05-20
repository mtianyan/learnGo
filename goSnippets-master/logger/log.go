package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var DefaultLogger = &Logger{
	writer: os.Stdout,
	action: func(wr io.Writer) {
		if _, file, _, ok := runtime.Caller(2); ok {
			fmt.Fprintf(wr, "\n%v 测试结果如下：\n", file)
		}
	},
}

type LogAction func(writer io.Writer)

type Logger struct {
	writer io.Writer
	action LogAction
}

func (l *Logger) Log() {
	l.action(l.writer)
}

func (l *Logger) SetWriter(wr io.Writer) *Logger {
	l.writer = wr
	return l
}

func (l *Logger) SetAction(fn LogAction) *Logger {
	l.action = fn
	return l
}
