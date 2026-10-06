package types

type TokenType int

type Token struct {
	TokType TokenType
	Num     float64
}

const (
	TokNumber TokenType = iota
	TokPlus
	TokMinus
	TokMul
	TokDiv
	TokLBracket
	TokRBracket
	TokEOF
)
