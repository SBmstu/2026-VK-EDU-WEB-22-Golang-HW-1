package types

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyExpression  = errors.New("empty expression")
	ErrInvalidNumber    = errors.New("invalid number")
	ErrUnexpectedChar   = errors.New("unexpected character")
	ErrUnexpectedToken  = errors.New("unexpected token")
	ErrExpectedRBracket = errors.New("expected closing bracket")
	ErrDivisionByZero   = errors.New("division by zero")
	ErrTrailingTokens   = errors.New("unexpected token after expression")
)

func CreateError(err error, val any) error {
	return fmt.Errorf("%w: %v", err, val)
}
