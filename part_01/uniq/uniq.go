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

func Uniq(lines []string, options Options) ([]string, error) {
	if (options.Count && options.Repeated) || (options.Count && options.Unique) || (options.Repeated && options.Unique) {
		return nil, fmt.Errorf("options -c, -d, -u are mutually exclusive")
	}

	if len(lines) == 0 {
		return []string{}, nil
	}

	var result []string
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

func lineKey(line string, options Options) string {
	runes := []rune(line)
	pos := 0
	if options.SkipFields > 0 {
		fieldsSkipped := 0
		i := 0
		for i < len(runes) && fieldsSkipped < options.SkipFields {
			for i < len(runes) && runes[i] == ' ' {
				i++
			}

			if i >= len(runes) {
				break
			}

			for i < len(runes) && runes[i] != ' ' {
				i++
			}

			fieldsSkipped++

			for i < len(runes) && runes[i] == ' ' {
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
