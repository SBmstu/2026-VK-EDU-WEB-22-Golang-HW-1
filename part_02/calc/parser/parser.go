package parser

import (
	"calc/types"
)

type Parser struct {
	tokens []types.Token
	pos    int
}


func (p *Parser) peek() types.Token {
	if (p.pos >= len(p.tokens)) {
		return types.Token{TokType: types.TokEOF}
	}

	return p.tokens[p.pos]
}

func (p *Parser) next() types.Token {
	t := p.tokens[p.pos]
	p.pos++

	return t
}

func (p *Parser) parseExpr() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		t := p.peek()
		if t.TokType != types.TokPlus && t.TokType != types.TokMinus {
			break
		}

		p.next()

		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}

		if t.TokType == types.TokPlus {
			left += right
		} else {
			left -= right
		}
	}

	return left, nil
}

func (p *Parser) parseTerm() (float64, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for {
		t := p.peek()
		if t.TokType != types.TokMul && t.TokType != types.TokDiv {
			break
		}

		p.next()

		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		if t.TokType == types.TokMul {
			left *= right
		} else {
			if right == 0 {
				return 0, types.ErrDivisionByZero
			}

			left /= right
		}
	}

	return left, nil
}

func (p *Parser) parseFactor() (float64, error) {
	t := p.peek()

	switch t.TokType {
	case types.TokMinus:
		p.next()

		v, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		return -v, nil
	case types.TokPlus:
		p.next()

		return p.parseFactor()
	case types.TokNumber:
		p.next()

		return t.Num, nil
	case types.TokLBracket:
		p.next()

		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}

		if p.peek().TokType != types.TokRBracket {
			return 0, types.ErrExpectedRBracket
		}

		p.next()

		return v, nil
	}

	return 0, types.ErrUnexpectedToken
}


func CreateParser(tokens []types.Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() (float64, error) {
	result, err := p.parseExpr()
	if err != nil {
		return 0, err
	}

	if p.peek().TokType != types.TokEOF {
		return 0, types.ErrTrailingTokens
	}

	return result, nil
}
