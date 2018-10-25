package main

import (
	"bufio"
	"os"

	err "github.com/ddosakura/pet-shop-project-playground/example-0/error"
	"github.com/ddosakura/pet-shop-project-playground/example-0/lexer"
	"github.com/ddosakura/pet-shop-project-playground/example-0/parser"
	"github.com/ddosakura/pet-shop-project-playground/example-0/symbol"
	"github.com/ddosakura/pet-shop-project-playground/logger"
)

func init() {
	logger.Info("call symbolInit()")
	symbol.Init()
}

func main() {
	go func(fileName string) {
		logger.Info("open " + fileName)
		if f, e := os.Open(fileName); e == nil {
			defer f.Close()
			for reader := bufio.NewReader(f); true; {
				r, _, e2 := reader.ReadRune()
				if e2 != nil {
					lexer.Put(0)
					return
				}
				lexer.Put(r)
			}
		}
		err.Finish(-1, fileName+" open failed!")
	}(os.Args[1])
	logger.Info("call parse()")
	parser.Start()
	err.Finish(0)
}
