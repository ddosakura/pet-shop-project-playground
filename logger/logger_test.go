package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	info("Hello World!")
	error("find a bug")
	info("Hello World!")
}
