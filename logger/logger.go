package logger

import (
	"log"
)

// Info - log info message
func Info(msg string) {
	log.Printf("%c[%d;;%dm%s%c[0m %s", 0x1B, 0, 34, "[ info  ]", 0x1B, msg)
}

// Error - log error message
func Error(msg string) {
	log.Printf("%c[%d;;%dm%s%c[0m %s", 0x1B, 1, 31, "[ error ]", 0x1B, msg)
}
