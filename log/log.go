package log

import (
	"github.com/fatih/color"
	"os"
)

// Level define log level
type Level int

const (
	// DebugLevel debug level
	DebugLevel Level = iota
	// InfoLevel info level
	InfoLevel
	// WarnLevel warn level
	WarnLevel
	// ErrorLevel error level
	ErrorLevel
)

// SetLevel set log level
func SetLevel(level string) {
	switch level {
	case "debug":
		logLevel = DebugLevel
	case "info":
		logLevel = InfoLevel
	case "warn":
		logLevel = WarnLevel
	case "error":
		logLevel = ErrorLevel
	default:
		logLevel = InfoLevel
	}
}

func Debug(args ...interface{}) {
	if logLevel > DebugLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgCyan)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
}

func Debugf(format string, args ...interface{}) {
	if logLevel > DebugLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgCyan)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
}

func Info(args ...interface{}) {
	if logLevel > InfoLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgGreen)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
}

func Infof(format string, args ...interface{}) {
	if logLevel > InfoLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgGreen)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
}

func Warn(args ...interface{}) {
	if logLevel > WarnLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgYellow)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
}

func Warnf(format string, args ...interface{}) {
	if logLevel > WarnLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgYellow)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
}

func Error(args ...interface{}) {
	if logLevel > ErrorLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgRed)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
}

func Errorf(format string, args ...interface{}) {
	if logLevel > ErrorLevel {
		return
	}

	colorLog.Add(color.Reset)
	colorLog.Add(color.FgRed)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
}

func Fatal(args ...interface{}) {
	colorLog.Add(color.Reset)
	colorLog.Add(color.FgRed)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
	os.Exit(-1)
}

func Fatalf(format string, args ...interface{}) {
	colorLog.Add(color.Reset)
	colorLog.Add(color.FgRed)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
	os.Exit(-1)
}

func Hint(args ...interface{}) {
	colorLog.Add(color.Reset)
	colorLog.Add(color.FgHiBlue)
	_, err := colorLog.Println(args...)
	if err != nil {
		panic(err)
	}
}

func Hintf(format string, args ...interface{}) {
	colorLog.Add(color.Reset)
	colorLog.Add(color.FgHiBlue)
	_, err := colorLog.Printf(format, args...)
	if err != nil {
		panic(err)
	}
}

func DrawParagraph(title string, level Level, writeBody ...func()) {
	if level < logLevel {
		return
	}

	const line = "========================%s========================\n"
	colorLog.Add(color.Reset)

	// head
	colorLog.Add(color.FgHiBlue, color.Underline)
	_, err := colorLog.Printf(line, title)
	if err != nil {
		panic(err)
	}

	// body
	for _, writeBody := range writeBody {
		writeBody()
	}

	// tail
	colorLog.Add(color.Reset)

	// head
	colorLog.Add(color.FgHiBlue, color.Underline)
	needFill := len(title)
	fillStr := make([]byte, 0, needFill)
	for i := 0; i < needFill; i++ {
		fillStr = append(fillStr, '=')
	}
	_, err = colorLog.Printf(line, string(fillStr))
	if err != nil {
		panic(err)
	}
}
