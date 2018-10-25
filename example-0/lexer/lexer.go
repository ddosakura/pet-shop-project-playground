package lexer

import (
	"strconv"

	err "github.com/ddosakura/pet-shop-project-playground/example-0/error"
	"github.com/ddosakura/pet-shop-project-playground/example-0/symbol"
)

// Lexan - lexer
func Lexan(flag ...bool) (token int, val interface{}) {
	var r = Get()
	for ; r > 0; r = Get() {
		// logger.Info("get " + strconv.Itoa(int(r)))
		if isSpace(r) {
			// pass
			continue
		} else if int(r) == 10 {
			err.LineAdd()
			if checkFlag(flag, 0) {
				return 10, nil
			}
			continue
		} else if int(r) == int('#') {
			return FastTravel(10)
		} else if isDigit(r) {
			var n = int(r - 48)
			r = Get()
			for ; isDigit(r); r = Get() {
				n = n*10 + int(r-48)
			}
			Rollback(r)
			return symbol.NUM, n
		} else if isAlpha(r) {
			s := ""
			for ; isDigitOrAlpha(r); r = Get() {
				s += string(r)
			}
			Rollback(r)
			token, ok := symbol.Table[s]
			if ok {
				return token, s
			}
			symbol.Table[s] = symbol.ID
			return symbol.ID, s
		} /* else if r == 0 {
			return symbol.DONE, nil
		}*/
		return int(r), nil
	}
	// err.Error("lexer error!")
	return symbol.DONE, nil
}

// FastTravel - find next word
func FastTravel(f int) (token int, val interface{}) {
	// logger.Info("target=" + strconv.Itoa(f))
	getEnter := false
	if f == 10 {
		// logger.Info("fast travel need to get enter")
		getEnter = true
	}
	for t, _ := Lexan(getEnter); t != f; t, _ = Lexan(getEnter) {
		// fmt.Println("fast look", t)
		if t == symbol.DONE {
			err.Error("can't fast travel to target=" + strconv.Itoa(f))
		}
	}
	token, val = Lexan(getEnter)
	// fmt.Println("fast look ret", token, val)
	return
}

func isSpace(r rune) bool {
	i := int(r)
	return i == 32 || i == 9 || i == 13
}

func isDigit(r rune) bool {
	i := int(r)
	return 48 <= i && i <= 57
}

func isAlpha(r rune) bool {
	i := int(r)
	return (97 <= i && i <= 122) || (65 <= i && i <= 90)
}

func isDigitOrAlpha(r rune) bool {
	return isDigit(r) || isAlpha(r)
}

func checkFlag(flag []bool, pos int) (ret bool) {
	defer func() {
		e := recover()
		if e != nil {
			ret = false
		}
	}()
	return flag[pos]
}
