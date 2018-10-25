package symbol

const (
	_ int = iota + 255
	// NUM - number
	NUM
	// DIV - div
	DIV
	// MOD - mod
	MOD
	// ID - ID
	ID
	// DONE - DONE
	DONE
)

// Table - the table of symbol
var Table map[string]int
