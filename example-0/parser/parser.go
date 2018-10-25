package parser

import (
	"github.com/ddosakura/pet-shop-project-playground/example-0/emitter"
	err "github.com/ddosakura/pet-shop-project-playground/example-0/error"
	"github.com/ddosakura/pet-shop-project-playground/example-0/lexer"
	"github.com/ddosakura/pet-shop-project-playground/example-0/symbol"
)

var lookahead int
var val interface{}

// Start - the entry of parser
func Start() {
	lookahead, val = lexer.Lexan()
	for lookahead != symbol.DONE {
		func() {
			defer func() {
				e := recover()
				if e != nil {
					if e.(int) > 0 {
						// fmt.Println("pre fast look", lookahead, val)
						lookahead, val = lexer.FastTravel(int(';'))
					} else {
						err.Finish(-1)
					}
				}
			}()
			expr()
			match(int(';'))
		}()
	}
}

func match(t int) {
	// defer logger.Info("match return")
	// if val != nil && lookahead != symbol.NUM {
	// 	logger.Info("call match(" + strconv.Itoa(t) + ") lookahead=" + strconv.Itoa(lookahead) + " val=" + val.(string))
	// } else {
	// 	logger.Info("call match(" + strconv.Itoa(t) + ") lookahead=" + strconv.Itoa(lookahead))
	// }
	if lookahead == t {
		lookahead, val = lexer.Lexan()
	} else {
		err.Warning("match syntax error!")
		panic(1)
	}
}

func expr() {
	// defer logger.Info("expr return")
	// logger.Info("call expr()")
	term()
	for {
		switch lookahead {
		case int('+'), int('-'):
			t := lookahead
			match(lookahead)
			term()
			emitter.Emit(t, nil)
			continue
		default:
			return
		}
	}
}

func term() {
	// defer logger.Info("term return")
	// logger.Info("call term()")
	factor()
	for {
		switch lookahead {
		case int('*'), int('/'), symbol.DIV, symbol.MOD:
			t := lookahead
			match(lookahead)
			factor()
			emitter.Emit(t, nil)
			continue
		default:
			return
		}
	}
}

func factor() {
	// defer logger.Info("factor return")
	// logger.Info("call factor() lookahead=" + strconv.Itoa(lookahead))
	switch lookahead {
	case int('('):
		match('(')
		expr()
		match(')')
		break
	case symbol.NUM:
		emitter.Emit(symbol.NUM, val)
		match(symbol.NUM)
		break
	case symbol.ID:
		emitter.Emit(symbol.ID, val)
		match(symbol.ID)
		break
	default:
		err.Warning("factor syntax error!")
		panic(2)
	}
}
