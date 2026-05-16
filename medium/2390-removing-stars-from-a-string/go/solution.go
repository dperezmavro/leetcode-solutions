package removingstarsfromastring

import (
	"bytes"
	"strings"
)

func removeStarsStr(s string) string {
	stack := make([]string, len(s))
	for _, r := range s {
		if r != '*' {
			stack = append(stack, string(r))
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return strings.Join(stack, "")
}

func removeStarsRune(s string) string {
	stack := make([]rune, len(s))

	for _, r := range s {
		if r != '*' {
			stack = append(stack, r)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return ""
	}

	b := bytes.Buffer{}
	for _, r := range stack {
		if r != 0 {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func removeStars(s string) string {
	validIndexes := make([]int, len(s))
	lastValidPtr := 0

	for i, r := range s {
		if r != '*' {
			validIndexes[lastValidPtr] = i
			lastValidPtr++
		} else {
			lastValidPtr--
		}
	}

	if lastValidPtr == 0 {
		return ""
	}

	b := bytes.Buffer{}

	for _, v := range validIndexes[:lastValidPtr] {
		b.WriteByte(s[v])
	}
	return b.String()
}
