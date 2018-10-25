package err

import (
	"os"
	"strconv"

	"github.com/ddosakura/pet-shop-project-playground/logger"
)

var errNum = 0

var lineID = 1
var posID = 1

// LineAdd - go to next line
func LineAdd() {
	lineID++
	posID = 1
}

// PosAdd - go to next pos
func PosAdd() {
	posID++
}

// PosSub - go to last pos
func PosSub() {
	posID--
}

// Warning - log error
func Warning(msg string) {
	logger.Error("(line " + strconv.Itoa(lineID) + ", pos " + strconv.Itoa(posID) + "): " + msg)
	errNum++
}

// Error - log error and exit
func Error(msg string) {
	Warning(msg)
	Finish(1)
}

// Finish - finish message
func Finish(code int, msgs ...string) {
	if code == 0 {
		for msg := range msgs {
			logger.Info(msgs[msg])
		}
	} else {
		for msg := range msgs {
			logger.Error(msgs[msg])
			errNum++
		}
	}
	logger.Info("finish with " + strconv.Itoa(errNum) + " error(s)")
	os.Exit(code)
}
