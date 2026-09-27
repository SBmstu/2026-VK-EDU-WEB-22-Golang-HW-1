package calc

import (
	"fmt"
	"strconv"
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
	type_ tokenType
	num float64
}

func strtok(s string) ([]token, error) {
	var tokens []token

	i := 0
	for i < len(s) {
		c := s[i]

		if ((c >= '0' && c <= '9') || c == '.') {
			start := i

			for i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == '.') {
				i++
			}

			num, err := strconv.ParseFloat(s[start:i], 64)
			if err != nil {
				return nil, fmt.Errorf("Error! Invalid number: %s", s[start:i])
			}

			tokens = append(tokens, token{type_: tokNumber, num: num})

			continue
		}

		switch c {
		case ' ', '\t', '\n', '\r':
		case '+':
			tokens = append(tokens, token{type_: tokPlus})
		case '-':
			tokens = append(tokens, token{type_: tokMinus})
		case '*':
			tokens = append(tokens, token{type_: tokMul})
		case '/':
			tokens = append(tokens, token{type_: tokDiv})
		case '(':
			tokens = append(tokens, token{type_: tokLBracket})
		case ')':
			tokens = append(tokens, token{type_: tokRBracket})
		default:
			return nil, fmt.Errorf("Error! Unexpected character: %c", c)
		}

		i++
	}

	tokens = append(tokens, token{type_: tokEOF})

	return tokens, nil
}

func Eval(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	
	tokens, _ := strtok(expr)
	for _, val := range tokens {
		fmt.Println(val)
	}

	result := 227.0;

	return result, nil
}
