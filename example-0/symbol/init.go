package symbol

import (
	"github.com/ddosakura/pet-shop-project-playground/logger"
)

// Init - init symtable
func Init() {
	logger.Info("symbol initing...")
	Table = map[string]int{
		"div": DIV,
		"mod": MOD,
	}
	// fmt.Println(NUM, DIV, MOD, ID, DONE)
}
