package uniq

import (
	"fmt"
	"strings"
)

type Options struct {
	Count bool
	Repeated bool
	Unique bool
	IgnoreCase bool
	SkipFields int
	SkipChars int
}

func ValidateOptions(opts Options) error {
	if (opts.Count && opts.Repeated) || (opts.Count && opts.Unique) || (opts.Repeated && opts.Unique) {
		return fmt.Errorf("options -c, -d, -u are mutually exclusive")
	}

	return nil
}

func Uniq(lines []string, options Options) ([]string, error) {
	if err := ValidateOptions(options); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return []string{}, nil
	}

	result := make([]string, 0, len(lines))
	i := 0
	for i < len(lines) {
		key := lineKey(lines[i], options)
		j := i + 1
		for j < len(lines) && lineKey(lines[j], options) == key {
			j++
		}

		count := j - i
		line := lines[i]

		switch {
		case options.Count:
			result = append(result, fmt.Sprintf("%d %s", count, line))
		case options.Repeated:
			if count > 1 {
				result = append(result, line)
			}
		case options.Unique:
			if count == 1 {
				result = append(result, line)
			}
		default:
			result = append(result, line)
		}

		i = j
	}

	return result, nil
}

func isSpace(r rune) bool {
	if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
		return true
	}

	return false
}

func lineKey(line string, options Options) string {
	runes := []rune(line)
	pos := 0
	if options.SkipFields > 0 {
		fieldsSkipped := 0
		i := 0
		for i < len(runes) && fieldsSkipped < options.SkipFields {
			for i < len(runes) && isSpace(runes[i]) {
				i++
			}

			if i >= len(runes) {
				break
			}

			for i < len(runes) && !isSpace(runes[i]) {
				i++
			}

			fieldsSkipped++

			for i < len(runes) && isSpace(runes[i]) {
				i++
			}
		}

		pos = i
	}

	if options.SkipChars > 0 {
		pos += options.SkipChars
		if pos > len(runes) {
			pos = len(runes)
		}
	}

	if pos >= len(runes) {
		return ""
	}

	key := string(runes[pos:])

	if options.IgnoreCase {
		key = strings.ToLower(key)
	}

	return key
}
