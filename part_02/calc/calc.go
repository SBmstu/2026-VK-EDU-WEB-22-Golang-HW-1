package calc

import (
	"fmt"
	"strconv"
	"strings"

	"calc/calc/parser"
	"calc/types"
)


func strtok(s string) ([]types.Token, error) {
	var tokens []types.Token

	i := 0
	for i < len(s) {
		c := s[i]

		if (c >= '0' && c <= '9') || c == '.' {
			start := i

			for i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == '.') {
				i++
			}

			num, err := strconv.ParseFloat(s[start:i], 64)
			if err != nil {
				return nil, fmt.Errorf("Error! Invalid number: %s", s[start:i])
			}

			tokens = append(tokens, types.Token{TokType: types.TokNumber, Num: num})

			continue
		}

		switch c {
		case ' ', '\t', '\n', '\r':
		case '+':
			tokens = append(tokens, types.Token{TokType: types.TokPlus})
		case '-':
			tokens = append(tokens, types.Token{TokType: types.TokMinus})
		case '*':
			tokens = append(tokens, types.Token{TokType: types.TokMul})
		case '/':
			tokens = append(tokens, types.Token{TokType: types.TokDiv})
		case '(':
			tokens = append(tokens, types.Token{TokType: types.TokLBracket})
		case ')':
			tokens = append(tokens, types.Token{TokType: types.TokRBracket})
		default:
			return nil, fmt.Errorf("Error! Unexpected character: %c", c)
		}

		i++
	}

	tokens = append(tokens, types.Token{TokType: types.TokEOF})

	return tokens, nil
}

func Eval(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return 0, fmt.Errorf("Error! Empty expression")
	}

	tokens, err := strtok(expr)
	if err != nil {
		return 0, err
	}

	return parser.CreateParser(tokens).Parse()
}
