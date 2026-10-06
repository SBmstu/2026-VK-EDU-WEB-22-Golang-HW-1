package utils

import (
	"fmt"
	"strings"
	"uniq/types"
)


func ValidateOptions(opts types.Options) error {
	var set []string

	if opts.Count {
		set = append(set, types.FlagCount)
	}
	if opts.Repeated {
		set = append(set, types.FlagRepeated)
	}
	if opts.Unique {
		set = append(set, types.FlagUnique)
	}

	if len(set) > 1 {
		return fmt.Errorf("Error! Options -%s are mutually exclusive", strings.Join(set, ", -"))
	}

	return nil
}
