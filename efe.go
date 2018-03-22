package efe

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/kawakami-o3/efp"
)

type Tokens struct {
	efp.Tokens
}

func (tokens *Tokens) take() *efp.Token {
	tokens.Index++
	return &tokens.Items[tokens.Index]
}

func (tokens *Tokens) unget() {
	tokens.Index--
}

func readExpr(tokens *Tokens, prec int) *Tokens {

	//for { tok := tokens.take() }
	return tokens
}

func Eval(formula string) string {
	fmt.Println(formula)

	ps := efp.ExcelParser()
	ps.Parse(formula)

	pp.Println(&Tokens{ps.Tokens})
	return ""
}
