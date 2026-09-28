package calc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"calc/calc/parser"
	"calc/types"
)


func strtok(s string) ([]types.Token, error) {
	var tokens []types.Token

	runes := []rune(s)

	i := 0
	for i < len(runes) {
		c := runes[i]

		if unicode.IsSpace(c) {
			i++
			continue
		}

		if c >= '0' && c <= '9' || c == '.' {
			start := i

			for i < len(runes) && (runes[i] >= '0' && runes[i] <= '9' || runes[i] == '.') {
				i++
			}

			num, err := strconv.ParseFloat(string(runes[start:i]), 64)
			if err != nil {
				return nil, fmt.Errorf("Error! Invalid number: %s", string(runes[start:i]))
			}

			tokens = append(tokens, types.Token{TokType: types.TokNumber, Num: num})

			continue
		}

		switch c {
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
