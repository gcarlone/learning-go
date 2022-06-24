package main

import (
	"fmt"
)

type LogLevel uint8

const (
	DebugLevel LogLevel = iota + 1
	WarningLevel
	ErrorLevel
)

func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return fmt.Sprintf("unknown log level: %d", l)
}

func main() {
	fmt.Println(WarningLevel)

	lvl := LogLevel(12)
	fmt.Println(lvl)
}
