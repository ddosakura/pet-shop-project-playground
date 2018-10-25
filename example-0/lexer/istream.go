package lexer

import (
	"github.com/ddosakura/pet-shop-project-playground/example-0/error"
)

var ch = make(chan rune)
var buffer rune

// Put - reader put
func Put(r rune) {
	ch <- r
	if int(r) == 0 {
		close(ch)
	}
}

// Rollback - lexer rollback
func Rollback(r rune) {
	buffer = r
	err.PosSub()
}

// Get - lexer get
func Get() (ret rune) {
	err.PosAdd()
	if buffer == 0 {
		return <-ch
	}
	ret = buffer
	buffer = 0
	return
}
