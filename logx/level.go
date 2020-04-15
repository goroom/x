package logx

import (
	"fmt"
	"strings"
)

type Level int

func (l Level) String() string {
	return levelStringMap[l]
}

func (l Level) ConsoleColorString() string {
	return fmt.Sprintf("\033[;%dm%s\033[0m", l.ConsoleColorNum(), l.String())
}

func ParseLevel(str string) Level {
	switch strings.ToLower(str) {
	case "all":
		return ALL
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn", "warning":
		return WARN
	case "error", "err":
		return ERROR
	case "off":
		return OFF
	default:
		return ALL
	}
}

func (l Level) ConsoleColorNum() int {
	switch l {
	case DEBUG:
		return 34
	case INFO:
		return 32
	case WARN:
		return 33
	case ERROR:
		return 35
	default:
		return 37
	}
}

var levelStringMap = []string{
	"ALL", "DEBG", "INFO", "WARN", "EROR", "OFF",
}

const (
	ALL Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	OFF
)
