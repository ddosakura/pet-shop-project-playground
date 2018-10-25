package emitter

import (
	"fmt"
	"strconv"

	err "github.com/ddosakura/pet-shop-project-playground/example-0/error"
	"github.com/ddosakura/pet-shop-project-playground/example-0/symbol"
)

// Emit - output
func Emit(t int, val interface{}) {
	switch t {
	case int('+'), int('-'), int('*'), int('/'):
		fmt.Println(string(t))
		break
	case symbol.DIV:
		fmt.Println("DIV")
		break
	case symbol.MOD:
		fmt.Println("MOD")
		break
	case symbol.NUM, symbol.ID:
		fmt.Println(val)
		break
	default:
		err.Warning(strconv.Itoa(t))
	}
}
