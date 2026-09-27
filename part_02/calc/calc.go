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
	num   float64
}

func strtok(s string) ([]token, error) {
	var tokens []token

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

type parser struct {
	tokens []token
	pos    int
}

func (p *parser) peek() token {
	return p.tokens[p.pos]
}

func (p *parser) next() token {
	t := p.tokens[p.pos]
	p.pos++

	return t
}

func (p *parser) parseExpr() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		t := p.peek()
		if t.type_ != tokPlus && t.type_ != tokMinus {
			break
		}

		p.next()

		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}

		if t.type_ == tokPlus {
			left += right
		} else {
			left -= right
		}
	}

	return left, nil
}

func (p *parser) parseTerm() (float64, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for {
		t := p.peek()
		if t.type_ != tokMul && t.type_ != tokDiv {
			break
		}

		p.next()

		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		if t.type_ == tokMul {
			left *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("Error! Division by zero")
			}

			left /= right
		}
	}

	return left, nil
}

func (p *parser) parseFactor() (float64, error) {
	t := p.peek()

	switch t.type_ {
	case tokMinus:
		p.next()

		v, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		return -v, nil
	case tokPlus:
		p.next()

		return p.parseFactor()
	case tokNumber:
		p.next()

		return t.num, nil
	case tokLBracket:
		p.next()

		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}

		if p.peek().type_ != tokRBracket {
			return 0, fmt.Errorf("Error! Expected closing bracket")
		}

		p.next()

		return v, nil
	}

	return 0, fmt.Errorf("Error! Unexpected token")
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

	p := &parser{tokens: tokens}

	result, err := p.parseExpr()
	if err != nil {
		return 0, err
	}

	if p.peek().type_ != tokEOF {
		return 0, fmt.Errorf("Error! Unexpected token after expression")
	}

	return result, nil
}
