package uniq

import (
	"fmt"
	"strings"

	"uniq/types"
	"uniq/utils"
)


func Uniq(lines []string, options types.Options) ([]string, error) {
	if err := utils.ValidateOptions(options); err != nil {
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

func lineKey(line string, options types.Options) string {
	runes := []rune(line)
	pos := 0
	if options.SkipFields > 0 {
		fieldsSkipped := 0
		i := 0
		for i < len(runes) && fieldsSkipped < options.SkipFields {
			for i < len(runes) && !utils.IsSpace(runes[i]) {
				i++
			}

			if i >= len(runes) {
				break
			}

			for i < len(runes) &&  !utils.IsSpace(runes[i]) {
				i++
			}

			fieldsSkipped++

			for i < len(runes) && !utils.IsSpace(runes[i]) {
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
