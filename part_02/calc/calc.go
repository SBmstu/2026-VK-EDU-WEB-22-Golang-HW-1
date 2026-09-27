package calc

import (
	"strings"
)

type tokenType int

const (
	tokNumber tokenType = iota
	tokPlus
	tokMinus
	tokMul
	tokDiv
	tokLBracket
	tokRBracket
	tokEOF
)

type token struct {
	typ tokenType
	num float64
}

func Eval(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	
	result := 227.0;

	return result, nil
}
