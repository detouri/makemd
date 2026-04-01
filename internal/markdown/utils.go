package markdown

import "fmt"

func normaliseLevel(level int) int {
	if level < 2 {
		return 2
	}
	if level > 6 {
		return 6
	}
	return level
}

func toComment(s string) string {
	return fmt.Sprintf("<!-- %s -->", s)
}
