package utils

import (
	"fmt"
	"uniq/types"
)


func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
		return true
	}

	return false
}

func ValidateOptions(opts types.Options) error {
	if (opts.Count && opts.Repeated) || (opts.Count && opts.Unique) || (opts.Repeated && opts.Unique) {
		return fmt.Errorf("options -c, -d, -u are mutually exclusive")
	}

	return nil
}
